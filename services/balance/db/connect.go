package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"time"
)

func InitDb() (gocqlx.Session, error) {
	session, err := connect()
	if err != nil {
		return session, err
	}

	// Create the account table
	err = CreateTable(session, "accounts")
	if err != nil {
		return session, err
	}
	return session, nil
}

// Connect to the database and return a session, connect to the keyspace bank
func connect() (gocqlx.Session, error) {
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "bank"
	cluster.Timeout = 1 * time.Minute
	cluster.Consistency = gocql.Quorum
	// Wrap the gocql.Sesiion into a gocqlx.Session
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return session, err
	}
	return session, nil
}
