package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewElasticRoute),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	esRoute ElasticRoute,
) Routes {
	return Routes{
		esRoute,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
