package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DalyChouikh/gin-articles/internal/models"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func SaveLists() {
	tmpArticleList = models.ArticleList
}

func RestoreLists() {
	models.ArticleList = tmpArticleList
}

func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}
