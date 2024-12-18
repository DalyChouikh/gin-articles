package api

import (
	"net/http"
	"strconv"

	"github.com/DalyChouikh/gin-articles/internal/models"
	"github.com/DalyChouikh/gin-articles/internal/utils"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	utils.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(articleID); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
