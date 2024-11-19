package server

import (
	"github.com/DalyChouikh/gin-articles/internal/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	return router
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", api.ShowIndexPage)
}
