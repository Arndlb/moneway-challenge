package proto

import (
	"fmt"
	v1 "github.com/Arndlb/moneway-challenge/api/v1"
	"github.com/scylladb/gocqlx"
	"google.golang.org/grpc"
	"net"
)

func InitServer(session gocqlx.Session) error {
	// Server is listening on :
	address := "0.0.0.0:8081"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Printf("Server is listening on %v ...", address)


	var server = NewTransactionServiceServer(session)
	// Create a new grpc server
	s := grpc.NewServer()
	// Regsiter the grpcr service in order to respond to the client
	v1.RegisterTransactionServiceServer(s, server)

	// Listen on port 8081
	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
