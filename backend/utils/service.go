//+build !test

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/web"
)

// CreateAPIGroup create an API router group
func CreateAPIGroup() (*gin.Engine, *gin.RouterGroup) {
	if LocalConf.Deploy == "product" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	rg := router.Group("/" + GetStringConfig("api_config", "version"))
	return router, rg
}

// RunWebService start a web service
func RunWebService(name string, router *gin.Engine) {
	if !CheckInTest() {
		service := web.NewService(
			web.Name(GetStringConfig("api_config", "namespace")+"."+
				GetStringConfig("api_config", "version")+"."+name),
			web.RegisterTTL(time.Second*time.Duration(ConsulConf.Get("srv_config", "ttl").Int(60))),
			web.RegisterInterval(time.Second*time.Duration(ConsulConf.Get("srv_config", "interval").Int(60))),
		)

		LogPanic(service.Init())
		service.Handle("/", router)
		Info("Running web service \"%s\"", name)
		LogPanic(service.Run())
	}
}

// InitMicroService init a micro service
func InitMicroService(name string) micro.Service {
	if !CheckInTest() {
		service := micro.NewService(
			micro.Name(GetServiceName(name)),
			micro.RegisterTTL(time.Second*time.Duration(ConsulConf.Get("srv_config", "ttl").Int(60))),
			micro.RegisterInterval(time.Second*time.Duration(ConsulConf.Get("srv_config", "interval").Int(60))),
		)
		service.Init()
		return service
	}
	return micro.NewService()
}

// RunMicroService run a micro service
func RunMicroService(service micro.Service) {
	if !CheckInTest() {
		Info("Running micro service \"%s\"", service.Options().Server.Options().Name)
		LogPanic(service.Run())
	}
}

// GetServiceName get micro service name
func GetServiceName(name string) string {
	return GetStringConfig("srv_config", "namespace") + "." + name
}

// CallMicroService call a micro service
func CallMicroService(name string, f func(name string, c client.Client) interface{}, m func() interface{}) interface{} {
	if CheckInTest() {
		f(GetServiceName(name), client.DefaultClient) // just useless test
		return m()
	}
	Info("Calling micro service \"%s\"", name)
	c := client.NewClient(
		client.RequestTimeout(time.Second*30),
		client.Selector(selector.DefaultSelector),
	)
	return f(GetServiceName(name), c)
}
