package server

import (
	"context"

	"golang-gin/base"

	"go.uber.org/fx"
)

var Module = fx.Options(
	base.Module,
	fx.Invoke(Server),
)

func Server(
	lifecycle fx.Lifecycle,
	handler base.Router,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				handler.Gin.Run("localhost:3000")
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
