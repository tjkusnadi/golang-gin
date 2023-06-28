package controllers

import (
	"golang-gin/api/services"
	"golang-gin/base"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type ElasticController struct {
	logger  base.Logger
	service services.ElasticService
}

type CreateIndexBody struct {
	IndexName string `json:"indexName" validate:"nonnil,min=1,max=10"`
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

func (ec ElasticController) CreateIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateIndexBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			ec.logger.Zap.Error("error ......", requestBody.IndexName)
			c.JSON(400, gin.H{
				"error": "index name required",
			})
			return
		}

		if err := validator.Validate(requestBody); err != nil {
			c.JSON(400, gin.H{
				"error":   err,
				"message": "error when validate index name",
			})
			return
		}

		ec.logger.Zap.Info("IndexName", requestBody.IndexName)

		ec.service.CreateIndex(requestBody.IndexName)

		c.JSON(200, gin.H{
			"success": true,
		})
	}
}
