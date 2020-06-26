package proto

import (
	"context"
	v1 "github.com/moneway-challenge/api/v1"
)

type balanceServiceServer struct {}

func (b balanceServiceServer) GetBalance(ctx context.Context, request *v1.GetBalanceRequest) (*v1.GetBalanceResponse, error) {
	panic("implement me")
}

func (b balanceServiceServer) UpdateBalance(ctx context.Context, request *v1.UpdateBalanceRequest) (*v1.UpdateBalanceResponse, error) {
	panic("implement me")
}

func NewBalanceServiceServer() v1.BalanceServiceServer {
	return &balanceServiceServer{}
}
