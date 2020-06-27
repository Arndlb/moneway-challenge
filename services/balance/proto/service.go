package proto

import (
	"context"
	"errors"
	"github.com/gocql/gocql"
	v1 "github.com/moneway-challenge/api/v1"
	"github.com/moneway-challenge/services/balance/db"
	"github.com/scylladb/gocqlx"
	"log"
)

type balanceServiceServer struct {
	session gocqlx.Session
}

func (b *balanceServiceServer) GetBalance(ctx context.Context, request *v1.GetBalanceRequest) (*v1.GetBalanceResponse, error) {
	log.Print("Get")
	accountID, err := gocql.ParseUUID(request.AccountId)
	if err !=  nil {
		return nil, err
	}

	balance, err := db.GetBalance(b.session, accountID)
	if err != nil {
		return nil, err
	}

	response := &v1.GetBalanceResponse{
		Balance: balance,
	}
	return response, nil
}

func (b *balanceServiceServer) UpdateBalanceCredit(ctx context.Context, request *v1.UpdateBalanceCreditRequest) (*v1.UpdateBalanceCreditResponse, error) {
	log.Print("update")
	accountID, err := gocql.ParseUUID(request.AccountId)
	if err !=  nil {
		return nil, err
	}
	amount := request.Amount
	if amount <= 0 {
		return nil, errors.New("error: Cannot add a negative or zero number")
	}
	oldBalance, err := db.GetBalance(b.session, accountID)
	if err != nil {
		return nil, err
	}

	newBalance := oldBalance + amount
	err = db.UpdateBalance(b.session, accountID, newBalance)
	if err != nil {
		return nil, err
	}
	response := &v1.UpdateBalanceCreditResponse{
		Balance: newBalance,
	}
	return response, nil
}

func (b *balanceServiceServer) UpdateBalanceDebit(ctx context.Context, request *v1.UpdateBalanceDebitRequest) (*v1.UpdateBalanceDebitResponse, error) {
	log.Print("update")
	accountID, err := gocql.ParseUUID(request.AccountId)
	if err !=  nil {
		return nil, err
	}
	oldBalance, err := db.GetBalance(b.session, accountID)
	if err != nil {
		return nil, err
	}

	amount := request.Amount
	if amount <= 0  || oldBalance < amount {
		return nil, errors.New("error: invalid debit")
	}

	newBalance := oldBalance - amount
	err = db.UpdateBalance(b.session, accountID, newBalance)
	if err != nil {
		return nil, err
	}
	response := &v1.UpdateBalanceDebitResponse{
		Balance: newBalance,
	}
	return response, nil
}

func NewBalanceServiceServer(session gocqlx.Session) v1.BalanceServiceServer {
	return &balanceServiceServer{
		session: session,
	}
}
