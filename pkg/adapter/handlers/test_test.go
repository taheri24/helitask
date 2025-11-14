package handlers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spyzhov/ajson"
	"github.com/taheri24/helitask/pkg/logger"
	"github.com/taheri24/helitask/pkg/ports/storage"
	"github.com/taheri24/helitask/pkg/ports/storage/sqlite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func setupApp(t *testing.T, datasetFn string) (*gin.Engine, *fxtest.App) {
	db, app := sqlite.NewDb(t, datasetFn).Debug(), gin.New()
	app.Use(handlerNameInHeader)
	fxApp := fxtest.New(t, fx.NopLogger, fx.Provide(logger.Nop), fx.Supply(db, app), storage.Module, Module)
	return app, fxApp
}

func setupHTTP(httpMethod, path, body string) (*http.Request, *httptest.ResponseRecorder) {
	b := strings.NewReader(body)

	return httptest.NewRequest(httpMethod, path, b), httptest.NewRecorder()
}
func extractJsonVal(source []byte, key string) string {
	root, err := ajson.Unmarshal(source)
	if err != nil {
		panic(err)
	}
	keyNode, err := root.GetKey(key)
	if err != nil {
		panic(err)
	}
	return keyNode.MustString()
}

func handlerNameInHeader(c *gin.Context) {
	c.Writer.Header().Set("X-Handler-Name", extractFuncShortName(c.Handler()))

	c.Next()

}

func extractFuncShortName(handler gin.HandlerFunc) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

	// Optional: Clean up the full path (remove package prefixes)
	shortName := funcName
	if idx := strings.LastIndex(funcName, "/"); idx != -1 {
		shortName = funcName[idx+1:]
	}
	return shortName
}

func init() {
	gin.SetMode(gin.TestMode)
}
