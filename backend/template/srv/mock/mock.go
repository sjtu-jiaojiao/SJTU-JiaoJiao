package mock

import (
	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (a *mockSrv) (ctx context.Context, req *, opts ...client.CallOption) (*, error) {
	var ret 
	
	return &ret, nil
}

func New Service() . Service {
	return new(mockSrv)
}
