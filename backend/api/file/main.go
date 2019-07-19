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
	rg.POST("/file", addFile)
	return router
}

type getFileQuery struct {
	FileId string `uri:"fileId"`
}

/**
 * @apiIgnore
 * @api {get} /file/:fileId GetFile
 * @apiVersion 1.0.0
 * @apiGroup File
 * @apiPermission none
 * @apiName GetFile
 * @apiDescription Get file
 *
 * @apiParam {--} Param see [File Service](#api-Service-file_File_Query)
 * @apiSuccess (Success 200) {Response} response see [File Service](#api-Service-file_File_Query)
 * @apiUse FileServiceDown
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

type addFileCreate struct {
	Stream []byte `form:"stream"`
}

/**
 * @apiIgnore
 * @api {post} /file AddFile
 * @apiVersion 1.0.0
 * @apiGroup File
 * @apiPermission none
 * @apiName AddFile
 * @apiDescription Add file
 *
 * @apiParam {--} Param see [File Service](#api-Service-file_File_Create)
 * @apiSuccess (Success 200) {Response} response see [File Service](#api-Service-file_File_Create)
 * @apiUse FileServiceDown
 */
func addFile(c *gin.Context) {
	var p addFileCreate

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		rsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
			Stream: p.Stream,
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
