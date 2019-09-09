package main

import (
	"context"
	"errors"
	db "jiaojiao/database"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"
	"time"

	"github.com/h2non/filetype"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/micro/go-micro/client"
)

type srv struct{}

type msgInfo struct {
	Time    time.Time    `bson:"time"`
	Forward bool         `bson:"forward"`
	Type    message.Type `bson:"type"`
	Msg     string       `bson:"msg"`
	Unread  bool         `bson:"unread"`
}

type chatLog struct {
	ID       primitive.ObjectID `bson:"_id"`
	FromUser int32              `bson:"fromUser"`
	ToUser   int32              `bson:"toUser"`
	Badge    int32              `bson:"badge"`
	Infos    []msgInfo          `bson:"infos"`
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
 * @apiParam {bytes} msg file stream bytes if type is picture or video, plain text if type is text
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *message.MessageCreateRequest, rsp *message.MessageCreateResponse) error {
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Type, req.Msg) || req.FromUser == req.ToUser {
		rsp.Status = message.MessageCreateResponse_INVALID_PARAM
		return nil
	}

	if (req.Type == message.Type_PICTURE && !utils.CheckFile(req.Msg, filetype.IsImage)) ||
		(req.Type == message.Type_VIDEO && !utils.CheckFile(req.Msg, filetype.IsVideo)) {
		rsp.Status = message.MessageCreateResponse_INVALID_TYPE
		return nil
	}

	msg := req.Msg
	// upload file
	if req.Type == message.Type_PICTURE || req.Type == message.Type_VIDEO {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		microRsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
			File: req.Msg,
		})

		if utils.LogContinue(err, utils.Error) {
			return err
		}
		if microRsp.Status != file.FileCreateResponse_SUCCESS {
			_, s := utils.LogContinueS("File create return "+microRsp.Status.String(), utils.Error)
			return errors.New(s)
		}

		msg = []byte(microRsp.FileID)
	}

	// find existence
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")
	var res1, res2 chatLog
	err1 := collection.FindOne(ctx, bson.M{
		"fromUser": req.FromUser,
		"toUser":   req.ToUser,
	}).Decode(&res1)
	err2 := collection.FindOne(ctx, bson.M{
		"fromUser": req.ToUser,
		"toUser":   req.FromUser,
	}).Decode(&res2)

	if err1 == nil && err2 == nil {
		utils.Error("chat structure violated")
		return errors.New("chat structure violated")
	}

	// new chat
	if err1 != nil && err2 != nil {
		_, err := collection.InsertOne(ctx, bson.M{
			"fromUser": req.FromUser,
			"toUser":   req.ToUser,
			"badge":    1,
			"infos": bson.A{bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    req.Type,
				"msg":     string(msg),
				"unread":  true,
			}},
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		rsp.Status = message.MessageCreateResponse_SUCCESS
		return nil
	}

	// exist
	from := req.FromUser
	to := req.ToUser
	forward := true
	res := res1
	if err1 != nil && err2 == nil {
		from = req.ToUser
		to = req.FromUser
		forward = false
		res = res2
	}

	_, err := collection.UpdateOne(ctx, bson.M{
		"fromUser": from,
		"toUser":   to,
	}, bson.M{
		"$push": bson.M{
			"infos": bson.M{
				"$each": bson.A{
					bson.M{
						"time":    time.Now(),
						"forward": forward,
						"type":    req.Type,
						"msg":     string(msg),
						"unread":  true,
					},
				},
				"$position": 0,
			},
		},
		"$set": bson.M{
			"badge": res.Badge + 1,
		},
	})
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	rsp.Status = message.MessageCreateResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Message.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Message.Find
 * @apiDescription Find Message
 *
 * @apiParam {int32} fromUser user who want to find
 * @apiParam {int32} toUser user who chat with from user
 * @apiParam {int32} way 1 for read message <br> 2 for query history message <br> Note: only 1 will set unread to false
 * @apiParam {uint32{0-20}} limit=20 limit of return message infos, only for history query
 * @apiParam {uint32} offset=0 offset from the latest message info, only for history query
 * @apiSuccess {int32} fromUser user who launch the chat at first time
 * @apiSuccess {int32} toUser user who accept the chat at first time
 * @apiSuccess {int32} badge count of message still unread
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {list} infos see below MessageInfo
 * @apiSuccess (MessageInfo) {int64} time message create time
 * @apiSuccess (MessageInfo) {bool} forward false for chat from toUser to fromUser <br> true for chat from fromUser to toUser
 * @apiSuccess (MessageInfo) {int32} type 1 for text <br> 2 for picture <br> 3 for video
 * @apiSuccess (MessageInfo) {bytes} text plain message text if type is text <br> fileID if type is picture or video
 * @apiSuccess (MessageInfo) {bool} unread false for having read <br> true for not having read
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *message.MessageFindRequest, rsp *message.MessageFindResponse) error {
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Way) || req.FromUser == req.ToUser {
		rsp.Status = message.MessageFindResponse_INVALID_PARAM
		return nil
	}
	if req.Limit == 0 {
		req.Limit = 20
	}
	if req.Limit > 20 {
		req.Limit = 20
	}

	decodeRes := func(src *chatLog, dest *message.MessageFindResponse) {
		dest.FromUser = src.FromUser
		dest.ToUser = src.ToUser
		dest.Badge = src.Badge
		for _, v := range src.Infos {
			dest.Infos = append(dest.Infos, &message.MessageInfo{
				Time:    v.Time.Unix(),
				Forward: v.Forward,
				Type:    v.Type,
				Msg:     v.Msg,
				Unread:  v.Unread,
			})
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")
	var res1, res2 chatLog
	err1 := collection.FindOne(ctx, bson.M{
		"fromUser": req.FromUser,
		"toUser":   req.ToUser,
	}).Decode(&res1)
	err2 := collection.FindOne(ctx, bson.M{
		"fromUser": req.ToUser,
		"toUser":   req.FromUser,
	}).Decode(&res2)

	if err1 == nil && err2 != nil {
		rsp.FromUser = req.FromUser
		rsp.ToUser = req.ToUser
	} else if err1 != nil && err2 == nil {
		rsp.FromUser = req.ToUser
		rsp.ToUser = req.FromUser
	} else if err1 != nil && err2 != nil {
		rsp.Status = message.MessageFindResponse_SUCCESS
		return nil
	} else if err1 == nil && err2 == nil {
		utils.Error("chat structure violated")
		return errors.New("chat structure violated")
	}

	if req.Way == message.MessageFindRequest_READ_MESSAGE {
		var res chatLog
		cur, err := collection.Aggregate(ctx, bson.A{
			bson.M{
				"$match": bson.M{
					"fromUser": rsp.FromUser,
					"toUser":   rsp.ToUser,
				},
			},
			bson.M{
				"$project": bson.M{
					"fromUser": 1,
					"toUser":   1,
					"badge":    1,
					"infos": bson.M{
						"$filter": bson.M{
							"input": "$infos",
							"as":    "item",
							"cond": bson.M{
								"$and": bson.A{
									bson.M{"$eq": bson.A{"$$item.forward", req.FromUser == rsp.FromUser}},
									bson.M{"$eq": bson.A{"$$item.unread", true}},
								},
							},
						},
					},
				},
			},
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		cur.Next(ctx)
		err = cur.Decode(&res)
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		decodeRes(&res, rsp)

		_, err = collection.UpdateMany(ctx, bson.M{
			"fromUser":      rsp.FromUser,
			"toUser":        rsp.ToUser,
			"infos.forward": req.FromUser == rsp.FromUser,
			"infos.unread":  true,
		}, bson.M{
			"$set": bson.M{
				"badge":                0,
				"infos.$[elem].unread": false,
			},
		}, &options.UpdateOptions{
			ArrayFilters: &options.ArrayFilters{
				Filters: bson.A{bson.M{
					"elem.forward": req.FromUser == rsp.FromUser,
					"elem.unread":  true,
				}},
			},
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
	} else {
		var res chatLog
		err := collection.FindOne(ctx, bson.M{
			"fromUser": rsp.FromUser,
			"toUser":   rsp.ToUser,
		}, &options.FindOneOptions{
			Projection: bson.M{
				"infos": bson.M{
					"$slice": bson.A{req.Offset, req.Limit},
				},
			},
		}).Decode(&res)
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		decodeRes(&res, rsp)
	}
	rsp.Status = message.MessageFindResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Message.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Message.Query
 * @apiDescription Query New Message, do NOT set read
 *
 * @apiParam {int32} userID user who wants to pull new message
 * @apiParam {bool} oldMsg=0 true to get all message, not only the new
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiSuccess {list} news see below NewMessage
 * @apiSuccess (NewMessage) {int32} fromUser user who launch the chat at first time
 * @apiSuccess (NewMessage) {int32} toUser user who accept the chat at first time
 * @apiSuccess (NewMessage) {int32} badge count of message still unread
 * @apiSuccess (NewMessage) {MessageInfo} info newest one msg for each dialog, see below MessageInfo
 * @apiSuccess (MessageInfo) {int64} time message create time
 * @apiSuccess (MessageInfo) {bool} forward false for chat from toUser to fromUser <br> true for chat from fromUser to toUser
 * @apiSuccess (MessageInfo) {int32} type 1 for text <br> 2 for picture <br> 3 for video
 * @apiSuccess (MessageInfo) {string} text plain message text if type is text <br> fileID if type is picture or video
 * @apiSuccess (MessageInfo) {bool} unread always true
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *message.MessageQueryRequest, rsp *message.MessageQueryResponse) error {
	if !utils.RequireParam(req.UserID) {
		rsp.Status = message.MessageQueryResponse_INVALID_PARAM
		return nil
	}

	decodeRes := func(src *chatLog, dest *message.NewMessage) {
		dest.FromUser = src.FromUser
		dest.ToUser = src.ToUser
		dest.Badge = src.Badge
		dest.Info = &message.MessageInfo{
			Time:    src.Infos[0].Time.Unix(),
			Forward: src.Infos[0].Forward,
			Type:    src.Infos[0].Type,
			Msg:     src.Infos[0].Msg,
			Unread:  src.Infos[0].Unread,
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	var cond bson.A
	if req.OldMsg {
		cond = bson.A{
			bson.M{
				"fromUser": req.UserID,
			},
			bson.M{
				"toUser": req.UserID,
			},
		}
	} else {
		cond = bson.A{
			bson.M{
				"fromUser": req.UserID,
				"badge":    bson.M{"$gt": 0},
				"infos": bson.M{
					"$elemMatch": bson.M{
						"forward": false,
						"unread":  true,
					},
				},
			},
			bson.M{
				"toUser": req.UserID,
				"badge":  bson.M{"$gt": 0},
				"infos": bson.M{
					"$elemMatch": bson.M{
						"forward": true,
						"unread":  true,
					},
				},
			},
		}
	}

	cur, err := collection.Find(ctx, bson.M{
		"$or": cond,
	}, &options.FindOptions{Projection: bson.M{
		"infos": bson.M{
			"$slice": 1,
		},
	}})
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	for cur.Next(ctx) {
		var res chatLog
		var newMessage message.NewMessage
		err = cur.Decode(&res)
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		decodeRes(&res, &newMessage)
		rsp.News = append(rsp.News, &newMessage)
	}

	rsp.Status = message.MessageQueryResponse_SUCCESS
	return nil
}

func main() {
	db.InitMongoDB("messagemongo")
	service := utils.InitMicroService("message")
	utils.LogPanic(message.RegisterMessageHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
