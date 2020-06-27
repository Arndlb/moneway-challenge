package proto

import (
	"fmt"
	v1 "github.com/Arndlb/moneway-challenge/api/v1"
	"github.com/scylladb/gocqlx"
	"google.golang.org/grpc"
	"net"
)

func InitServer(session gocqlx.Session) error {
	// Grpc server listen on :
	address := "0.0.0.0:8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Printf("Server is listening on %v ...", address)


	var server = NewBalanceServiceServer(session)
	// Create a new grpc server
	s := grpc.NewServer()
	// Register the Balance grpc services
	v1.RegisterBalanceServiceServer(s, server)

	// Listen on port 8080
	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
