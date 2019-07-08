package main

import (
	"context"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
)

type srvInfo struct{}

type srvContent struct{}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Query
 * @apiDescription Query sell info
 *
 * @apiParam {int32} sellInfoId sell info id.
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for non-exist
 * @apiUse DBServerDown
 */
//func (a *srv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, rsp *sellinfo.SellInfoQueryResponse) error {
//
//	return nil
//}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Create
 * @apiDescription create sell info
 *
 * @apiParam {int32} sellInfoId sell info id
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for non-exist
 * @apiUse DBServerDown
 */
func (a *srvInfo) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, rsp *sellinfo.SellInfoCreateResponse) error {
	return nil
}

/**
 * @api {rpc} /rpc sellinfo.Content.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.Content.Create
 * @apiDescription create sell info content
 *
 * @apiParam {string} [contentId] content id, left empty for first upload
 * @apiParam {string} [contentToken] content token, left empty for first upload
 * @apiParam {bytes} content binary content
 * @apiParam {int32} type 1 for picture <br> 2 for video
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for invalid token
 * @apiUse DBServerDown
 */
func (a *srvContent) Create(ctx context.Context, req *sellinfo.ContentCreateRequest, rsp *sellinfo.ContentCreateResponse) error {
	return nil
}

func main() {
	service := utils.InitMicroService("sellinfo")
	utils.LogPanic(sellinfo.RegisterSellInfoHandler(service.Server(), new(srvInfo)))
	utils.LogPanic(sellinfo.RegisterContentHandler(service.Server(), new(srvContent)))
	utils.RunMicroService(service)
}
