package proto

import (
	"google.golang.org/grpc"
	"log"
)

func InitClient() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()



}