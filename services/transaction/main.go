package main

import (
	"log"
	"moneway.go/moneway-challenge/services/transaction/db"
)

func main() {
	err := db.connect()
	if   err != nil {
		log.Fatalln(err)
	}
}