package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"log"
	"time"
)

func InitDb() error {
	session, err := connect()
	if err != nil {
		return err
	}

	err = CreateTable(session, "transaction")
	if err != nil {
		return err
	}
	id1, _ := gocql.ParseUUID("a7bea45b-e79e-4d4c-8df8-9c21afeace37")
	_ = UpdateTransaction(session, id1, "ncvrjfjvfnv", "dededededede")

	return nil
}

func connect() (gocqlx.Session, error) {
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "bank"
	cluster.Timeout = 1 * time.Minute
	cluster.Consistency = gocql.Quorum
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}
	return session, nil
}
