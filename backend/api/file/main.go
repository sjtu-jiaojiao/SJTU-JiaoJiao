package main

import (
	"context"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"github.com/h2non/filetype"
	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/file/:fileID", getFile)
	return router
}

/**
 * @apiDefine FileServiceDown
 * @apiError (Error 500) FileServiceDown File service down
 */

/**
 * @api {get} /file/:fileID GetFile
 * @apiVersion 1.0.0
 * @apiGroup File
 * @apiPermission none
 * @apiName GetFile
 * @apiDescription Get file
 *
 * @apiParam {--} Param see [File Service](#api-Service-File_Query)
 * @apiSuccess (Success 200) {File} file file itself
 * @apiUse InvalidParam
 * @apiError (Error 404) FileNotFound file not found
 * @apiUse FileServiceDown
 */
func getFile(c *gin.Context) {
	type param struct {
		FileID string `uri:"fileID" binding:"required"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)

		rsp, err := srv.Query(context.TODO(), &file.FileRequest{
			FileID: p.FileID,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}

		if rsp.Status == file.FileQueryResponse_SUCCESS {
			if utils.CheckFile(rsp.File, filetype.IsImage, filetype.IsAudio, filetype.IsVideo) {
				t, err := filetype.Match(rsp.File)
				if utils.LogContinue(err, utils.Error) {
					c.JSON(500, err)
					return
				}
				c.Data(200, t.MIME.Value, rsp.File)
			} else {
				c.AbortWithStatus(403)
			}
		} else {
			c.AbortWithStatus(404)
		}
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("file", setupRouter())
}
