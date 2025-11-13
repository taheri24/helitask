package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module fx.Option

func init() {
	var (
		todoHandler TodoHandler
	)
	decoratedServices := fx.Populate(helper, &todoHandler)
	Module = fx.Module("apiHttpRoutingV0", decoratedServices, fx.Invoke(

		func(appEngine *gin.Engine) {
			apiRouter := appEngine.Group("/api/v0")
			todoRouter := apiRouter.Group("/todo")
			todoRouter.POST("/", todoHandler.CreateTodoItem)
		},
	))

}
