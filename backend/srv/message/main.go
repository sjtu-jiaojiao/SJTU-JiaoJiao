package main

import (
	"context"
	db "jiaojiao/database"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/micro/go-micro/client"
)

type srv struct{}

/**
 * @api {rpc} /rpc Message.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Message.Create
 * @apiDescription Create Message
 *
 * @apiParam {int32} fromUser user who launch the chat at first time
 * @apiParam {int32} toUser user who accept the chat at first time
 * @apiParam {int32} badge count of message still unread
 * @apiParam {int64} time message create time
 * @apiParam {bool} forward false for chat from toUser to fromUser <br> true for chat from fromUser to toUser
 * @apiParam {int32} type 1 for text <br> 2 for picture <br> 3 for video
 * @apiParam {string} text plain message text if type is text <br> fileID if type is picture or video
 * @apiParam {bool} unread false for having read <br> true for not having read
 * @apiSuccess {} status -1 for invalid param <br> 1 for success
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *message.MessageCreateRequest, rsp *message.MessageCreateResponse) error {
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Type) {
		rsp.Status = message.MessageCreateResponse_INVALID_PARAM
		return nil
	}

	if req.Type == message.MessageCreateRequest_PICTURE || req.Type == message.MessageCreateRequest_VIDEO {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		microRsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
			File: req.File,
		})

		if err != nil || microRsp.Status != file.FileCreateResponse_SUCCESS {
			rsp.Status = message.MessageCreateResponse_INVALID_PARAM
			return nil
		}
		req.Text = microRsp.FileID
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("sellinfo")
	res1, err1 := collection.Find(ctx, bson.M{
		"fromUser": req.FromUser,
		"toUser":   req.ToUser,
	})
	res2, err2 := collection.Find(ctx, bson.M{
		"fromUser": req.ToUser,
		"toUser":   req.FromUser,
	})

	if err1 == nil && err2 != nil { //fromUser to toUser
		_, err := collection.UpdateOne(ctx, bson.M{
			"fromUser": req.FromUser,
			"toUser":   req.ToUser,
		}, bson.M{"$push": bson.M{
			"infos": bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    req.Type,
				"text":    req.Text,
				"unread":  true,
			},
		}})
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageCreateResponse_UNKNOWN
			return nil
		}

		_, err = collection.UpdateOne(ctx, bson.M{
			"fromUser": req.FromUser,
			"toUser":   req.ToUser,
		}, bson.M{"$set": bson.M{
			"badge": res1.Current.Lookup("badge").Int32() + 1,
		}})
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageCreateResponse_UNKNOWN
			return nil
		}
		rsp.Status = message.MessageCreateResponse_SUCCESS
		return nil
	} else if err1 != nil && err2 == nil { //toUser to fromUser
		_, err := collection.UpdateOne(ctx, bson.M{
			"fromUser": req.ToUser,
			"toUser":   req.FromUser,
		}, bson.M{"$push": bson.M{
			"infos": bson.M{
				"time":    time.Now(),
				"forward": false,
				"type":    req.Type,
				"text":    req.Text,
				"unread":  true,
			},
		}})
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageCreateResponse_UNKNOWN
			return nil
		}

		_, err = collection.UpdateOne(ctx, bson.M{
			"fromUser": req.ToUser,
			"toUser":   req.FromUser,
		}, bson.M{"$set": bson.M{
			"badge": res2.Current.Lookup("badge").Int32() + 1,
		}})
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageCreateResponse_UNKNOWN
			return nil
		}
		rsp.Status = message.MessageCreateResponse_SUCCESS
		return nil
	} else if err1 != nil && err2 != nil { //new chat
		_, err := collection.InsertOne(ctx, bson.M{
			"fromUser": req.FromUser,
			"toUser":   req.ToUser,
			"badge":    1,
			"infos": bson.A{bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    req.Type,
				"text":    req.Text,
				"unread":  true,
			}},
		})
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageCreateResponse_UNKNOWN
			return nil
		}
		rsp.Status = message.MessageCreateResponse_SUCCESS
		return nil
	} else {
		rsp.Status = message.MessageCreateResponse_INVALID_PARAM
	}
	return nil
}

func main() {
	service := utils.InitMicroService("message")
	utils.LogPanic(message.RegisterMessageHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
