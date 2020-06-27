package proto

import (
	"context"
	"fmt"
	v1 "github.com/moneway-challenge/api/v1"
	"google.golang.org/grpc"
)

func InitClient() error {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		return err
	}
	defer cc.Close()
	client := v1.NewBalanceServiceClient(cc)

	request := &v1.GetBalanceRequest{AccountId: "ad61ea77-f3df-4347-b4b9-4f80815f79c8"}
	resp, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Balance)

	request2 := &v1.UpdateBalanceCreditRequest{AccountId: "ad61ea77-f3df-4347-b4b9-4f80815f79c8", Amount: 100}
	resp2, _ := client.UpdateBalanceCredit(context.Background(), request2)
	fmt.Printf("Receive response => [%v]", resp2.Balance)

	return nil
}