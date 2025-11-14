package server

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/taheri24/helitask/pkg/config"
	"github.com/taheri24/helitask/pkg/logger"
	"github.com/taheri24/helitask/pkg/utils"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, defaultApp *gin.Engine, cfg *config.Config, logger logger.Logger) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listenPort := cfg.Server.Port
			if utils.IsNumber(listenPort) {
				listenPort = ":" + listenPort
			}
			if err := defaultApp.Run(listenPort); err != nil {
				slog.Error("Failed to start server", slog.Any("bindngErr", err))
				return err
			}
			return nil
		},
	})
}
