package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"log"
)

var accountMetadata = table.Metadata{
	Name: "accounts",
	Columns: []string{"account_id", "first_name", "last_name", "amount", "currency"},
	PartKey: []string{"account_id"},
	SortKey: []string{"last_name"},
}

var accountTable = table.New(accountMetadata)

type Account struct {
	AccountID gocql.UUID  `db:"account_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Amount    int64  `db:"amount"`
	Currency  string `db:"currency"`
}

// Create the table account if not exist
func CreateTable(session gocqlx.Session, name string) error {
	query := "CREATE TABLE IF NOT EXISTS " + name + " (" +
		"account_id uuid PRIMARY KEY," +
		"first_name varchar," +
		"last_name varchar," +
		"amount int," +
		"currency varchar)"

	err := session.Session.Query(query).Exec()
	if err != nil {
		return err
	}
	// Create an admin account for testing
	CreateAccount(session, "admin", "admin", 0, "euro")
	return nil
}

// Create an account in ScyllaDB with unique ID
func CreateAccount(session gocqlx.Session, lastName string, firstName string, amount int64, currency string) {
	uuid, _ := gocql.RandomUUID()
	account := Account{
		AccountID: uuid,
		FirstName: firstName,
		LastName:  lastName,
		Currency:  currency,
		Amount:    amount,
	}

	q := session.Query(accountTable.Insert()).BindStruct(account)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
}

func GetBalance(session gocqlx.Session, accountID gocql.UUID) (int64, error) {
	var account []Account
	q := session.Query(accountTable.Select()).BindMap(qb.M{"account_id": accountID})
	if err := q.SelectRelease(&account); err != nil {
		return 0, err
	}
	return account[0].Amount, nil
}

func UpdateBalance(session gocqlx.Session, accountID gocql.UUID, amount int64) error {
	account := Account{
		AccountID: accountID,
		Amount:    amount,
	}

	stmt, names := qb.Update("accounts").
		Set("amount").Where(qb.Eq("account_id")).ToCql()

	q := session.Query(stmt, names).BindStruct(&account)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
	return nil
}
