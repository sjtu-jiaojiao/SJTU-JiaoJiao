package main

import (
	"context"
	db "jiaojiao/database"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/micro/go-micro/client"
)

type srv struct{}

type MsgInfo struct {
	Time    time.Time                `bson:"time"`
	Forward bool                     `bson:"forward"`
	Type    message.MessageInfo_Type `bson:"type"`
	Text    string                   `bson:"text"`
	Unread  bool                     `bson:"unread"`
}

type ChatLog struct {
	ID       primitive.ObjectID `bson:"_id"`
	FromUser int32              `bson:"fromUser"`
	ToUser   int32              `bson:"toUser"`
	Badge    int32              `bson:"badge"`
	Infos    []MsgInfo          `bson:"infos"`
}

/**
 * @api {rpc} /rpc Message.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Message.Create
 * @apiDescription Create Message
 *
 * @apiParam {int32} fromUser user who launch the chat at first time
 * @apiParam {int32} toUser user who accept the chat at first time
 * @apiParam {int32} type 1 for text <br> 2 for picture <br> 3 for video
 * @apiParam {string} [text] plain message text if type is text
 * @apiParam {bytes} [file] file stream bytes, valid only if type is picture or video
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *message.MessageCreateRequest, rsp *message.MessageCreateResponse) error {
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Type) {
		rsp.Status = message.MessageCreateResponse_INVALID_PARAM
		return nil
	}

	if req.Type == message.MessageCreateRequest_TEXT && !utils.RequireParam(req.Text) {
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
	collection := db.MongoDatabase.Collection("message")
	var res1, res2 ChatLog
	err1 := collection.FindOne(ctx, bson.M{
		"fromUser": req.FromUser,
		"toUser":   req.ToUser,
	}).Decode(&res1)
	err2 := collection.FindOne(ctx, bson.M{
		"fromUser": req.ToUser,
		"toUser":   req.FromUser,
	}).Decode(&res2)

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
			"badge": res1.Badge + 1,
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
			"badge": res2.Badge + 1,
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

/**
 * @api {rpc} /rpc Message.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Message.Find
 * @apiDescription Find Message
 *
 * @apiParam {int32} fromUser user who launch the chat at first time
 * @apiParam {int32} toUser user who accept the chat at first time
 * @apiParam {int32} way 1 for only pull message <br> 2 for read message <br> 3 for query history message
 * @apiSuccess {int32} fromUser user who launch the chat at first time
 * @apiSuccess {int32} toUser user who accept the chat at first time
 * @apiSuccess {int32} badge count of message still unread
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {list} infos see below MessageInfo
 * @apiSuccess (MessageInfo) {int64} time message create time
 * @apiSuccess (MessageInfo) {bool} forward false for chat from toUser to fromUser <br> true for chat from fromUser to toUser
 * @apiSuccess (MessageInfo) {int32} type 1 for text <br> 2 for picture <br> 3 for video
 * @apiSuccess (MessageInfo) {string} text plain message text if type is text <br> fileID if type is picture or video
 * @apiSuccess (MessageInfo) {bool} unread false for having read <br> true for not having read
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *message.MessageFindRequest, rsp *message.MessageFindResponse) error {
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Way) {
		rsp.Status = message.MessageFindResponse_INVALID_PARAM
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")
	_, err1 := collection.Find(ctx, bson.M{
		"fromUser": req.FromUser,
		"toUser":   req.ToUser,
	})
	_, err2 := collection.Find(ctx, bson.M{
		"fromUser": req.ToUser,
		"toUser":   req.FromUser,
	})

	if err1 == nil && err2 != nil {
		rsp.FromUser = req.FromUser
		rsp.ToUser = req.ToUser
	} else if err1 != nil && err2 == nil {
		rsp.FromUser = req.ToUser
		rsp.ToUser = req.FromUser
	} else if err1 != nil && err2 != nil {
		rsp.Status = message.MessageFindResponse_NOT_FOUND
		return nil
	} else {
		rsp.Status = message.MessageFindResponse_INVALID_PARAM
		return nil
	}

	if req.Way == message.MessageFindRequest_ONLY_PULL {
		err := collection.FindOne(ctx, bson.M{
			"fromUser": rsp.FromUser,
			"toUser":   rsp.ToUser,
			"infos": bson.M{
				"forward": req.FromUser == rsp.FromUser,
				"unread":  true,
			},
		}).Decode(&rsp)
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageFindResponse_UNKNOWN
			return nil
		}
		rsp.Status = message.MessageFindResponse_SUCCESS
		return nil
	} else if req.Way == message.MessageFindRequest_READ_MESSAGE {
		err := collection.FindOneAndUpdate(ctx, bson.M{
			"fromUser": rsp.FromUser,
			"toUser":   rsp.ToUser,
			"infos": bson.M{
				"forward": req.FromUser == rsp.FromUser,
				"unread":  true,
			},
		}, bson.M{
			"$set": bson.M{
				"badge":  0,
				"unread": false,
			},
		}).Decode(&rsp)
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageFindResponse_UNKNOWN
			return nil
		}
		rsp.Status = message.MessageFindResponse_SUCCESS
		return nil
	} else if req.Way == message.MessageFindRequest_HISTORY {
		err := collection.FindOne(ctx, bson.M{
			"fromUser": rsp.FromUser,
			"toUser":   rsp.ToUser,
		}).Decode(&rsp)
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = message.MessageFindResponse_UNKNOWN
			return nil
		}
		rsp.Status = message.MessageFindResponse_SUCCESS
		return nil
	}
	return nil
}

func main() {
	db.InitMongoDB("messagemongo")
	service := utils.InitMicroService("message")
	utils.LogPanic(message.RegisterMessageHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
