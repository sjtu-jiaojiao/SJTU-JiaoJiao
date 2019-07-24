package mock

import (
	"context"
	transaction "jiaojiao/srv/transaction/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (mockSrv) Create(ctx context.Context, in *transaction.TransactionCreateRequest, opts ...client.CallOption) (*transaction.TransactionCreateResponse, error) {
	return nil, nil
}

func (mockSrv) Update(ctx context.Context, in *transaction.TransactionUpdateRequest, opts ...client.CallOption) (*transaction.TransactionUpdateResponse, error) {
	return nil, nil
}

func (mockSrv) Find(ctx context.Context, in *transaction.TransactionFindRequest, opts ...client.CallOption) (*transaction.TransactionFindResponse, error) {
	return nil, nil
}

// is service mock
func NewTransactionService() transaction.TransactionService {
	return new(mockSrv)
}
