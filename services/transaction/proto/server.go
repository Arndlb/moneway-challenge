package proto

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitServer() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}
