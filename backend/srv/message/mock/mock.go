package mock

import (
	"context"
	message "jiaojiao/srv/message/proto"

	"github.com/micro/go-micro/client"
)

type mockMessageSrv struct{}

func (a *mockMessageSrv) Find(ctx context.Context, in *message.MessageFindRequest, opts ...client.CallOption) (*message.MessageFindResponse, error) {
	var rsp message.MessageFindResponse
	// TODO
	return &rsp, nil
}

func (a *mockMessageSrv) Query(ctx context.Context, in *message.MessageQueryRequest, opts ...client.CallOption) (*message.MessageQueryResponse, error) {
	var rsp message.MessageQueryResponse
	// TODO
	return &rsp, nil
}

func (a *mockMessageSrv) Create(ctx context.Context, in *message.MessageCreateRequest, opts ...client.CallOption) (*message.MessageCreateResponse, error) {
	var rsp message.MessageCreateResponse
	// TODO
	return &rsp, nil
}

// is service mock
func NewMessageService() message.MessageService {
	return new(mockMessageSrv)
}
