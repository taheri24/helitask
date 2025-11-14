package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/taheri24/helitask/pkg/domain"
	"github.com/taheri24/helitask/pkg/logger"
	"go.uber.org/fx"
)

func ProvideTodoHandler(repository domain.TodoRepository) TodoHandler {
	return TodoHandler{repository}
}

var Module fx.Option

func init() {
	var (
		todoHandler TodoHandler
	)
	svcProviders := fx.Provide(ProvideTodoHandler)
	Module = fx.Module("apiHttpRoutingV0", svcProviders,
		fx.Populate(&todoHandler),
		fx.Invoke(
			func(appEngine *gin.Engine, logger logger.Logger) {
				helper.defaultLogger = logger
				apiRouter := appEngine.Group("/api/v0")
				{
					g, h := apiRouter.Group("/todo"), todoHandler
					g.POST("/", h.CreateTodoItem)
					g.GET("/:id", h.GetTodoItem)
				}

			},
		))

}
