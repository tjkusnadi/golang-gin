package base

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewRouter),
	fx.Provide(NewLogger),
	fx.Provide(NewElasticsearch),
	fx.Provide(NewPostgres),
)
