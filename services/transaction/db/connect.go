package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"log"
	"time"
)

// Initialize database and return session
func InitDb() (gocqlx.Session, error) {
	session, err := connect()
	if err != nil {
		return session, err
	}
	// Create transaction table
	err = CreateTable(session, "transaction")
	if err != nil {
		return session, err
	}
	return session, nil
}

// Connect to database and keyspace bank
func connect() (gocqlx.Session, error) {
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "bank"
	cluster.Timeout = 1 * time.Minute
	cluster.Consistency = gocql.Quorum
	// Wrap gocql session in to gocqlx session
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}
	return session, nil
}
