package main

import (
	"net/http"

	"github.com/DalyChouikh/gin-articles/internal/server"
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func main() {
	router := server.SetupRouter()
	server.RegisterRoutes(router)
	router.Run(":8080")
}
