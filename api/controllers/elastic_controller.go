package controllers

import (
	"golang-gin/api/services"
	"golang-gin/base"

	"github.com/gin-gonic/gin"
)

type ElasticController struct {
	logger  base.Logger
	service services.ElasticService
}

func NewElasticController(
	logger base.Logger,
	service services.ElasticService,
) ElasticController {
	return ElasticController{
		logger:  logger,
		service: service,
	}
}

func (ec ElasticController) GetInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		elasticInfo, _ := ec.service.GetInfo()

		c.JSON(200, gin.H{
			"info": elasticInfo,
		})
	}
}
