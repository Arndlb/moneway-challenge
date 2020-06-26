package proto

import (
	"fmt"
	v1 "github.com/moneway-challenge/api/v1"
	"github.com/scylladb/gocqlx"
	"google.golang.org/grpc"
	"net"
)


func InitServer(session gocqlx.Session) error {
	address := "0.0.0.0:8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Printf("Server is listening on %v ...", address)

	var server = NewBalanceServiceServer()
	s := grpc.NewServer()
	v1.RegisterBalanceServiceServer(s, server)

	 err = s.Serve(lis)
	 if err != nil {
		 return err
	 }
	return nil
}