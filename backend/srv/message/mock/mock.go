package mock

import (
	"context"
	"errors"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"

	"github.com/h2non/filetype"
	"github.com/micro/go-micro/client"
)

type mockMessageSrv struct{}

// Find is message find mock
func (a *mockMessageSrv) Find(ctx context.Context, req *message.MessageFindRequest, opts ...client.CallOption) (*message.MessageFindResponse, error) {
	var rsp message.MessageFindResponse
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Way) {
		rsp.Status = message.MessageFindResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.FromUser == 3000 {
		return nil, errors.New("")
	}
	rsp.Status = message.MessageFindResponse_SUCCESS
	return &rsp, nil
}

// Query is message query mock
func (a *mockMessageSrv) Query(ctx context.Context, req *message.MessageQueryRequest, opts ...client.CallOption) (*message.MessageQueryResponse, error) {
	var rsp message.MessageQueryResponse
	if !utils.RequireParam(req.UserID) {
		rsp.Status = message.MessageQueryResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.UserID == 3000 {
		return nil, errors.New("")
	}
	rsp.Status = message.MessageQueryResponse_SUCCESS
	return &rsp, nil
}

// Create is message create mock
func (a *mockMessageSrv) Create(ctx context.Context, req *message.MessageCreateRequest, opts ...client.CallOption) (*message.MessageCreateResponse, error) {
	var rsp message.MessageCreateResponse
	if !utils.RequireParam(req.FromUser, req.ToUser, req.Type, req.Msg) || req.FromUser == req.ToUser {
		rsp.Status = message.MessageCreateResponse_INVALID_PARAM
		return &rsp, nil
	}

	if (req.Type == message.Type_PICTURE && !utils.CheckFile(req.Msg, filetype.IsImage)) ||
		(req.Type == message.Type_VIDEO && !utils.CheckFile(req.Msg, filetype.IsVideo)) {
		rsp.Status = message.MessageCreateResponse_INVALID_TYPE
		return &rsp, nil
	}

	if req.FromUser == 3000 {
		return nil, errors.New("")
	}
	rsp.Status = message.MessageCreateResponse_SUCCESS
	return &rsp, nil
}

// NewMessageService is service mock
func NewMessageService() message.MessageService {
	return new(mockMessageSrv)
}
