package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func test(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)

	service.Init()

	// setup Greeter Server Client
	//cl = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	router := gin.Default()
	router.GET("/greeter", test)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		//log.Fatal(err)
	}
}
