package main

import (
	"github.com/moneway-challenge/services/transaction/db"
	"log"
)

func main() {
	err := db.InitDb()
	if   err != nil {
		log.Fatalln(err)
	}
}