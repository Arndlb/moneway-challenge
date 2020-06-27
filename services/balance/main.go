package main

import (
	"github.com/Arndlb/moneway-challenge/services/balance/db"
	"github.com/Arndlb/moneway-challenge/services/balance/proto"
	"log"
)

func main() {
	// Initialize Database and get gocql session
	session, err := db.InitDb()
	if err != nil {
		log.Fatalln(err)
	}
	// CLose database session
	defer session.Close()

	// Initialize grpc server on port 8080
	err = proto.InitServer(session)
	if err != nil {
		log.Fatalln(err)
	}

}
