package main

import (
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /auth GetUser
 * @apiVersion 0.1.0
 * @apiGroup User
 * @apiName Auth
 * @apiDescription This is the Description.
 * @apiSuccess {String} firstname Firstname of the User.
 */
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
