# Helitask Todo API happy path coverage

[Options]
variable base_url = http://localhost:8080/api/v0

# Create a new todo item
POST {{base_url}}/todo/
[Headers]
Content-Type: application/json
Accept: application/json

{
  "description": "Review Hurl API flows",
  "due_date": "2025-03-01T10:00:00Z"
}

HTTP 201
[Captures]
jsonpath "$.id" todo_id

[Asserts]
jsonpath "$.id" matches "^[0-9a-f-]{36}$"

# Retrieve the todo item that was just created
GET {{base_url}}/todo/{{todo_id}}
[Headers]
Accept: application/json

HTTP 200
[Asserts]
jsonpath "$.id" == "{{todo_id}}"
jsonpath "$.description" == "Review Hurl API flows"
jsonpath "$.due_date" matches "^2025-03-01 10:00:00 \+0000 UTC$"
