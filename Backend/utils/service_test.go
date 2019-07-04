package utils

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	. "github.com/smartystreets/goconvey/convey"
)

func testResponse(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Test response",
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
		router, rg := CreateAPIGroup()
		rg.GET("/auth", testResponse)
		service := web.NewService(
			web.Name(GetConfig("srv_config", "namespace") + "." +
				GetConfig("api_config", "version") + ".api_test"),
		)

		service.Init()
		service.Handle("/", router)
		service.Run()
		// So(func() {
		// 	RunService("auth_test", router)
		// }, ShouldPanic)
	})
	// Convey("Service run success test", t, func() {
	// 	router, rg := CreateAPIGroup()
	// 	rg.GET("/auth_test", testResponse)
	// 	So(func() {
	// 		RunService("auth_test", router)
	// 	}, ShouldNotPanic)
	// })
}
