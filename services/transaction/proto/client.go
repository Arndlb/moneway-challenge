package proto

import (
	"google.golang.org/grpc"
)

func InitClient() error {
	// Run client to send message to blance transaction, send to port 8080
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return err
	}

	// Close grpc client
	defer cc.Close()

	return nil
}
