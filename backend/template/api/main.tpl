package main

import (
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/{{SERVICE_NAME}}", )
	return router
}

/**
 * @apiIgnore
 * @api {} /{{SERVICE_NAME}} 
 * @apiVersion 1.0.0
 * @apiGroup 
 * @apiPermission 
 * @apiName 
 * @apiDescription 
 *
 * @apiParam 
 * @apiSuccess 
 * @apiError (Error 500) 
 */
func (c *gin.Context) {
	type param struct {
		string `form:""`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("{{SERVICE_NAME}}", setupRouter())
}
