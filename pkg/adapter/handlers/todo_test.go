package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateTodoItem tests the CreateTodoItem handler
func TestCreateTodoItem(t *testing.T) {
	app, fxApp := setupApp(t, "")
	fxApp.RequireStart()
	defer fxApp.RequireStop()
	h := &TodoHandler{}
	t.Run("Setup Request#1", func(t *testing.T) {
		input := `{"description": "Test Todo", "due_date": "2025-12-31T23:59:59Z"}`
		req, w := setupHTTP("POST", "/api/v0/todo/", input)
		app.ServeHTTP(w, req)
		if !assert.Equal(t, http.StatusCreated, w.Code) {
			t.Log(w.Body.String())
		}
		assert.Equal(t, w.Header().Get("X-Handler-Name"), extractFuncShortName(h.CreateTodoItem))

		id := extractJsonVal(w.Body.Bytes(), "id")
		t.Run("Read it ", func(t *testing.T) {
			req, w := setupHTTP("GET", `/api/v0/todo/`+id, "")
			app.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, id, extractJsonVal(w.Body.Bytes(), "id"))
			assert.Equal(t, w.Header().Get("X-Handler-Name"), extractFuncShortName(h.GetTodoItem))

		})
		t.Run("Not Found ", func(t *testing.T) {
			req, w := setupHTTP("GET", `/api/v0/todo/00000000-0000-0000-0000-000000000000`, "")
			app.ServeHTTP(w, req)
			assert.Equal(t, http.StatusNotFound, w.Code)

		})

	})

}

// TestCreateTodoItemValidation tests input validation for CreateTodoItem
func TestCreateTodoItemValidation(t *testing.T) {
	app, fxApp := setupApp(t, "")
	fxApp.RequireStart()
	defer fxApp.RequireStop()

	type TestCase struct {
		name           string
		input          string
		expectedStatus int
		errorContains  string
	}

	testCases := []TestCase{
		{
			name:           "Empty description",
			input:          `{"description": "", "due_date": "2025-12-31T23:59:59Z"}`,
			expectedStatus: http.StatusBadRequest,
			errorContains:  "description is required",
		},
		{
			name:           "Description too long",
			input:          fmt.Sprintf(`{"description": "%s", "due_date": "2025-12-31T23:59:59Z"}`, strings.Repeat("a", MaxDescriptionLength+1)),
			expectedStatus: http.StatusBadRequest,
			errorContains:  "description exceeds maximum length",
		},
		{
			name:           "Missing due_date",
			input:          `{"description": "Test Todo"}`,
			expectedStatus: http.StatusBadRequest,
			errorContains:  "due_date is required",
		},
		{
			name:           "Zero due_date",
			input:          `{"description": "Test Todo", "due_date": "0001-01-01T00:00:00Z"}`,
			expectedStatus: http.StatusBadRequest,
			errorContains:  "due_date is required",
		},
		{
			name:           "Valid input",
			input:          `{"description": "Test Todo", "due_date": "2025-12-31T23:59:59Z"}`,
			expectedStatus: http.StatusCreated,
			errorContains:  "",
		},
		{
			name:           "Description at minimum length",
			input:          `{"description": "A", "due_date": "2025-12-31T23:59:59Z"}`,
			expectedStatus: http.StatusCreated,
			errorContains:  "",
		},
		{
			name:           "Description at maximum length",
			input:          fmt.Sprintf(`{"description": "%s", "due_date": "2025-12-31T23:59:59Z"}`, strings.Repeat("a", MaxDescriptionLength)),
			expectedStatus: http.StatusCreated,
			errorContains:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, w := setupHTTP("POST", "/api/v0/todo/", tc.input)
			app.ServeHTTP(w, req)

			if !assert.Equal(t, tc.expectedStatus, w.Code) {
				t.Log(w.Body.String())
			}

			if tc.errorContains != "" {
				errorMsg := extractJsonVal(w.Body.Bytes(), "error")
				assert.Contains(t, errorMsg, tc.errorContains)
			}
		})
	}
}

// TestGetTodoItemInvalidUUID tests UUID validation for GetTodoItem
func TestGetTodoItemInvalidUUID(t *testing.T) {
	app, fxApp := setupApp(t, "")
	fxApp.RequireStart()
	defer fxApp.RequireStop()

	type TestCase struct {
		name           string
		uuid           string
		expectedStatus int
		errorContains  string
	}

	testCases := []TestCase{
		{
			name:           "Invalid UUID format",
			uuid:           "not-a-uuid",
			expectedStatus: http.StatusBadRequest,
			errorContains:  "Invalid UUID",
		},
		{
			name:           "UUID with invalid characters",
			uuid:           "12345-67890-abcdef-ghijk",
			expectedStatus: http.StatusBadRequest,
			errorContains:  "Invalid UUID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, w := setupHTTP("GET", fmt.Sprintf("/api/v0/todo/%s", tc.uuid), "")
			app.ServeHTTP(w, req)

			if !assert.Equal(t, tc.expectedStatus, w.Code) {
				t.Log(w.Body.String())
			}

			if tc.expectedStatus == http.StatusBadRequest {
				errorMsg := extractJsonVal(w.Body.Bytes(), "error")
				assert.Contains(t, errorMsg, tc.errorContains)
			}
		})
	}
}

// TestGetTodoItem tests the CreateTodoItem handler
func TestGetTodoItem(t *testing.T) {
	app, fxApp := setupApp(t, "sample1.sql")
	fxApp.RequireStart()
	defer fxApp.RequireStop()
	h := &TodoHandler{}
	t.Run("ContentOfSample1.sql", func(t *testing.T) {

		type TestCase struct {
			uuid           string
			httpStatusCode int
			description    string
		}
		testCases := []TestCase{
			{"3f6c1a4e-9966-4f1c-a2a9-1b8df67f8cc3", http.StatusOK, "Buy groceries"},
			{"c2e89319-e563-4a0b-9ef0-349beb3ef672", http.StatusOK, "Finish project report"},
			{"8a2b2a84-0583-4a58-8c11-7e7b4d62c06a", http.StatusOK, "Call the electrician"},
			{"ad8c040c-d2c0-4dd8-9f2f-dc191b020b8d", http.StatusOK, "Schedule dentist appointment"},
			{"02c2a8a8-2b2b-4ce5-9d33-9b18e4b0e15f", http.StatusOK, "Plan weekend trip"},
			{"00000000-0000-0000-0000-000000000000", http.StatusNotFound, ""}, // false possibility
		} // it points to records of sample1.sql
		for _, testCase := range testCases {
			req, w := setupHTTP("GET", fmt.Sprintf(`/api/v0/todo/%s`, testCase.uuid), "")
			app.ServeHTTP(w, req)
			if !assert.Equal(t, testCase.httpStatusCode, w.Code) {
				t.Log(w.Body.String())
			}
			if testCase.httpStatusCode != http.StatusOK {
				continue
			}
			assert.Equal(t, w.Header().Get("X-Handler-Name"), extractFuncShortName(h.GetTodoItem))
			assert.Equal(t, extractJsonVal(w.Body.Bytes(), "description"), testCase.description)

		}
	})

}
