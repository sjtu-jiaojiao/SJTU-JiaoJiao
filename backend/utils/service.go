//+build !test

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
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
	if !CheckInTest() {
		service := web.NewService(
			web.Name(GetStringConfig("api_config", "namespace")+"."+
				GetStringConfig("api_config", "version")+"."+name),
			web.RegisterTTL(time.Second*time.Duration(ConsulConf.Get("srv_config", "ttl").Int(60))),
			web.RegisterInterval(time.Second*time.Duration(ConsulConf.Get("srv_config", "interval").Int(60))),
		)

		LogPanic(service.Init())
		service.Handle("/", router)
		LogPanic(service.Run())
	}
}

// RunMicroService run a micro service
func RunMicroService(name string, f func(s server.Server, hdlr interface{},
	opts ...server.HandlerOption) error, hdlr interface{}, opts ...server.HandlerOption) {
	if !CheckInTest() {
		service := micro.NewService(
			micro.Name(GetServiceName(name)),
			micro.RegisterTTL(time.Second*time.Duration(ConsulConf.Get("srv_config", "ttl").Int(60))),
			micro.RegisterInterval(time.Second*time.Duration(ConsulConf.Get("srv_config", "interval").Int(60))),
		)

		service.Init()
		LogPanic(f(service.Server(), hdlr, opts...))
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
	return f(GetServiceName(name), client.DefaultClient)
}
