//+build !test

package utils

import (
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// StartTestServer start the test server
func StartTestServer(f func() *gin.Engine, m string, p string, b io.Reader,
	setter func(r *http.Request)) *httptest.ResponseRecorder {
	router := f()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(m, "/"+GetStringConfig("api_config", "version")+"/"+p, b)
	if setter != nil {
		setter(req)
	}
	router.ServeHTTP(w, req)
	return w
}

// CheckInTest check if running in test
func CheckInTest() bool {
	return strings.HasSuffix(os.Args[0], ".test") || strings.Contains(os.Args[0], "/_test/") || flag.Lookup("test.v") != nil
}
