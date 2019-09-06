//+build !test

package utils

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// StartTestServer start the test server
func StartTestServer(f func() *gin.Engine, m string, p string, b io.Reader,
	setter func(r *http.Request)) *httptest.ResponseRecorder {
	router := f()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(m, "/"+GetStringConfig("api_config", "version")+p, b)
	if setter != nil {
		setter(req)
	}
	router.ServeHTTP(w, req)
	if req.Body != nil {
		req.Body.Close()
	}
	return w
}

// CheckInTest check if running in test
func CheckInTest() bool {
	return strings.HasSuffix(os.Args[0], ".test") || strings.Contains(os.Args[0], "/_test/") || flag.Lookup("test.v") != nil
}

// SingleRoleTest test single role permission
func SingleRoleTest(f func() *gin.Engine, allow bool, role string, method string, path string, p url.Values) bool {
	r := StartTestServer(f, method, path, strings.NewReader(p.Encode()),
		func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Set("Authorization", role)
		})
	if allow && r.Code == 403 {
		return false
	} else if !allow && r.Code != 403 {
		return false
	}
	return true
}

// RoleTest test role permission
func RoleTest(f func() *gin.Engine, role Role, method string, path string, p url.Values) int {
	var test Role
	test.Guest = SingleRoleTest(f, role.Guest, "", method, path, p)
	test.User = SingleRoleTest(f, role.User, "user", method, path, p)
	test.Self = SingleRoleTest(f, role.Self, "self", method, path, p)
	test.Admin = SingleRoleTest(f, role.Admin, "admin", method, path, p)
	if !test.Guest {
		return 1
	}
	if !test.User {
		return 2
	}
	if !test.Self {
		return 3
	}
	if !test.Admin {
		return 4
	}
	return 0
}

// GetTestData request test data
func GetTestData(f func() *gin.Engine, method string, path string, p url.Values, role string) (int, map[string]interface{}) {
	var data map[string]interface{}
	r := StartTestServer(f, method, path, strings.NewReader(p.Encode()),
		func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Set("Authorization", role)
		})
	if r.Code == 200 && r.Body.String() != "{}" && r.Body.String()[0] == '{' {
		LogPanic(json.Unmarshal(r.Body.Bytes(), &data))
	} else {
		return r.Code, nil
	}
	return r.Code, data
}

// If is something like <bool>?<exp1>:<exp2>
func If(b bool, t interface{}, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}
