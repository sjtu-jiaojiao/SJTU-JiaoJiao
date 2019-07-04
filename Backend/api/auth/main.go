package main

import (
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API1",
	})
}

func main() {
	router, rg := utils.CreateAPIGroup()

	rg.GET("/auth", test)

	utils.RunService("auth", router)
}
