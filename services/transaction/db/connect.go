package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"log"
	"time"
)

func InitDb() (gocqlx.Session, error) {
	session, err := connect()
	if err != nil {
		return session, err
	}

	err = CreateTable(session, "transaction")
	if err != nil {
		return session, err
	}
	return session, nil
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
