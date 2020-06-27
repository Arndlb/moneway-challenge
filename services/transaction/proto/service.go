package proto

import (
	"context"
	v1 "github.com/Arndlb/moneway-challenge/api/v1"
	"github.com/Arndlb/moneway-challenge/services/transaction/db"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
)

type transactionServiceServer struct {
	session gocqlx.Session
}

// Create a transaction with info given by grpc message and return the new id via grpc server
func (t *transactionServiceServer) CreateTransaction(ctx context.Context, request *v1.CreateTransactionRequest) (*v1.CreateTransactionResponse, error) {
	senderID, err := gocql.ParseUUID(request.SenderId)
	if err != nil {
		return nil, err
	}
	receiverID, err := gocql.ParseUUID(request.ReceiverId)
	if err != nil {
		return nil, err
	}

	transaction := db.Transaction{
		SenderID:    senderID,
		ReceiverID:  receiverID,
		Description: request.Description,
		Notes:       request.Notes,
		Amount:      request.Amount,
		Currency:    request.Currency,
	}

	newTransaction, err := db.CreateTransaction(t.session, &transaction)
	if err != nil {
		return nil, err
	}

	response := &v1.CreateTransactionResponse{
		Id: newTransaction.ID.String(),
	}
	return response, nil
}

// Update an existing transaction with info given by grpc message and return id of the transaction via grpc server
func (t *transactionServiceServer) UpdateTransaction(ctx context.Context, request *v1.UpdateTransactionRequest) (*v1.UpdateTransactionResponse, error) {
	ID, err := gocql.ParseUUID(request.Id)
	if err != nil {
		return nil, err
	}

	id, err := db.UpdateTransaction(t.session, ID, request.Notes, request.Description)
	if err != nil {
		return nil, err
	}

	response := &v1.UpdateTransactionResponse{
		Id: id.String(),
	}
	return response, nil

}

// Incomplete function
func (t *transactionServiceServer) GetTransactions(ctx context.Context, request *v1.GetTransactionRequest) (*v1.GetTransactionResponse, error) {
	panic("implement me")
}

// Return NewTransactionServiceServer interface with database session
func NewTransactionServiceServer(session gocqlx.Session) v1.TransactionServiceServer {
	return &transactionServiceServer{
		session: session,
	}
}
