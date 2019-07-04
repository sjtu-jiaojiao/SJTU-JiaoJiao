package utils

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func test1(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API1",
	})
}

func TestCreateAPIGroup(t *testing.T) {
	Convey("Service create test", t, func() {
		_, rg := CreateAPIGroup()
		So(rg.BasePath(), ShouldEqual, "/v1")
	})
}

func TestRunService(t *testing.T) {
	Convey("Service run fail test", t, func() {
		router, _ := CreateAPIGroup()
		So(func() {
			RunService("auth1", router)
		}, ShouldPanic)
	})
	Convey("Service run success test", t, func() {
		router, rg := CreateAPIGroup()
		rg.GET("/auth1", test1)
		So(func() {
			RunService("auth1", router)
		}, ShouldNotPanic)
	})
}
