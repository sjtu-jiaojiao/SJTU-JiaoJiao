//+build !test

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
)

// CreateAPIGroup create an API router group
func CreateAPIGroup() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	rg := router.Group("/" + GetStringConfig("api_config", "version"))
	return router, rg
}

// RunWebService start a web service
func RunWebService(name string, router *gin.Engine) {
	service := web.NewService(
		web.Name(GetStringConfig("api_config", "namespace") + "." +
			GetStringConfig("api_config", "version") + "." + name),
	)

	LogPanic(service.Init())
	service.Handle("/", router)
	LogPanic(service.Run())
}

// InitMicroService init a micro service
func InitMicroService(name string) micro.Service {
	service := micro.NewService(
		micro.Name(GetStringConfig("srv_config", "namespace")+"."+name),
		micro.RegisterTTL(time.Second*time.Duration(consulConf.Get("srv_config", "ttl").Int(60))),
		micro.RegisterInterval(time.Second*time.Duration(consulConf.Get("srv_config", "interval").Int(60))),
	)

	service.Init()
	return service
}

// RunMicroService run a micro service
func RunMicroService(service micro.Service) {
	LogPanic(service.Run())
}
