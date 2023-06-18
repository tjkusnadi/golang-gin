package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

func NewRouter() Router {
	httpRouter := gin.Default()

	httpRouter.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return Router{
		Gin: httpRouter,
	}
}
