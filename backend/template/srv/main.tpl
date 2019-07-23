package main

import (
	"context"
	"jiaojiao/utils"
)

type srv struct{}

/**
 * @apiIgnore
 * @api {rpc} /rpc 
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName 
 * @apiDescription 
 *
 * @apiParam {} 
 * @apiSuccess {} status -1 for invalid param <br> 1 for success
 * @apiUse DBServerDown
 */
func (a *srv) (ctx context.Context, req *, rsp *) error {
	return nil
}

func main() {
	service := utils.InitMicroService("{{SERVICE_NAME}}")
	utils.LogPanic(.Register(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
