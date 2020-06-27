package main

import (
	"github.com/moneway-challenge/services/transaction/db"
	"github.com/moneway-challenge/services/transaction/proto"
	"log"
)

func main() {
	session, err := db.InitDb()
	if   err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	err = proto.InitClient()
	if   err != nil {
		log.Fatalln(err)
	}
}