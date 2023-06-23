package server

import (
	"context"

	"golang-gin/api/controllers"
	"golang-gin/api/repositories"
	"golang-gin/api/routes"
	"golang-gin/api/services"
	"golang-gin/base"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	base.Module,
	services.Module,
	repositories.Module,
	fx.Invoke(Server),
)

func Server(
	lifecycle fx.Lifecycle,
	handler base.Router,
	routes routes.Routes,
	logger base.Logger,
	elastic base.Elastic,
	env base.Env,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				routes.Setup()

				logger.Zap.Info("Starting application on port %s", env.ServerPort)
				handler.Gin.Run("localhost:" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
