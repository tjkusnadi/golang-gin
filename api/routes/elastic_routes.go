package routes

import (
	"golang-gin/api/controllers"
	"golang-gin/base"
)

type ElasticRoute struct {
	controller controllers.ElasticController
	router     base.Router
}

func NewElasticRoute(
	controller controllers.ElasticController,
	router base.Router) ElasticRoute {
	return ElasticRoute{
		controller: controller,
		router:     router,
	}
}

func (esRoute ElasticRoute) Setup() {
	elasticRoute := esRoute.router.Gin.Group("/elastic")
	{
		elasticRoute.GET("/info", esRoute.controller.GetInfo())
	}
}
