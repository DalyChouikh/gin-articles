package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DalyChouikh/gin-articles/internal/testutils"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := testutils.GetRouter(true)

	r.GET("/", ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testutils.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
