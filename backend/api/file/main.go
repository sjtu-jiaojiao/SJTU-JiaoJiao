package main

import (
	"context"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/file/:fileId", getFile)
	return router
}

type getFileQuery struct {
	FileId string `uri:"fileId"`
}

/**
 * @apiIgnore
 * @api {} /file
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
func getFile(c *gin.Context) {
	var p getFileQuery

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		rsp, err := srv.Query(context.TODO(), &file.FileQueryRequest{
			FileId: p.FileId,
		})
		if utils.LogContinue(err, utils.Warning, "File service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("file", setupRouter())
}
