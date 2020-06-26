package main

import (
	"github.com/moneway-challenge/services/balance/db"
	"github.com/moneway-challenge/services/balance/proto"
	"log"
)

func main() {
	session, err := db.InitDb()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close();

	err = proto.InitServer(session)
	if err != nil {
		log.Fatalln(err)
	}

}
