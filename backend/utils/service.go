//+build !test

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

// CreateAPIGroup create an API router group
func CreateAPIGroup() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	rg := router.Group("/" + GetConfig("api_config", "version"))
	return router, rg
}

// RunService start a micro service
func RunService(name string, router *gin.Engine) {
	service := web.NewService(
		web.Name(GetConfig("srv_config", "namespace") + "." +
			GetConfig("api_config", "version") + "." + name),
	)

	LogPanic(service.Init())
	service.Handle("/", router)
	LogPanic(service.Run())
}
