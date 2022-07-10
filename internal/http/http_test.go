package http

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Request(c *gin.Engine, method string, url string) ([]byte, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, url, nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	return responseData, w
}

func TestSync(t *testing.T) {
	r := GetRoute()
	defer CircleCtxCancel()
	Request(r, "GET", "/schedule")

	start := time.Now().Second()
	Request(r, "POST", "/add?timeDuration=5s&sync")
	if elapsed := time.Now().Second() - start; elapsed != 5 {
		t.Errorf("time operation not valid %v", elapsed)
	}
}
