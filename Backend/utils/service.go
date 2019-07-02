package utils

import (
	"github.com/micro/go-micro/web"
)

func RunService(name string){
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