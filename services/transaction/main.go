package main

import (
	"github.com/Arndlb/moneway-challenge/services/transaction/db"
	"github.com/Arndlb/moneway-challenge/services/transaction/proto"
	"log"
)

func main() {
	// Initialize database and return a session
	session, err := db.InitDb()
	if err != nil {
		log.Fatalln(err)
	}
	// Close datbase session
	defer session.Close()

	// Init grpc client in order to communicate with balance grpc server
	err = proto.InitClient()
	if err != nil {
		log.Fatalln(err)
	}

	// Init grpc server in order to respond to a client
	err = proto.InitServer(session)
	if err != nil {
		log.Fatalln(err)
	}
}
