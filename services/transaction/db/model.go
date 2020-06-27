package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"log"
	"time"
)

// Schema for transaction table
var transactionMetadata = table.Metadata{
	Name:    "transaction",
	Columns: []string{"id", "sender_id", "receiver_id", "created_at", "description", "amount", "currency", "notes"},
	PartKey: []string{"id"},
	SortKey: []string{"created_at"},
}

var transactionTable = table.New(transactionMetadata)

// Transaction represents a transaction
// which can be a transfer, a card payment, etc.
type Transaction struct {
	ID          gocql.UUID `db:"id"`
	SenderID    gocql.UUID `db:"sender_id"`
	ReceiverID  gocql.UUID `db:"receiver_id"`
	CreatedAt   time.Time  `db:"created_at"`
	Description string     `db:"description"`
	Amount      int64      `db:"amount"`
	Currency    string     `db:"currency"`
	Notes       string     `db:"notes"`
}

// Create the table account if not exist
func CreateTable(session gocqlx.Session, name string) error {
	query := "CREATE TABLE IF NOT EXISTS " + name + " (" +
		"id uuid PRIMARY KEY," +
		"sender_id uuid ," +
		"receiver_id uuid ," +
		"created_at timestamp," +
		"description varchar," +
		"notes varchar," +
		"amount int," +
		"currency varchar)"

	err := session.Session.Query(query).Exec()
	if err != nil {
		return err
	}
	return nil
}

// Add a new transaction to database and return info of the transaction
func CreateTransaction(session gocqlx.Session, transaction *Transaction) (*Transaction, error) {
	uuid, _ := gocql.RandomUUID()
	transaction.ID = uuid
	transaction.CreatedAt = time.Now()

	log.Print(transaction)

	q := session.Query(transactionTable.Insert()).BindStruct(transaction)
	if err := q.ExecRelease(); err != nil {
		return transaction, err
	}
	return transaction, nil
}

// Get and return a specific transaction
func GetTransactionsById(session gocqlx.Session, id gocql.UUID) ([]Transaction, error) {
	var transactions []Transaction
	q := session.Query(transactionTable.Select()).BindMap(qb.M{"id": id})
	if err := q.SelectRelease(&transactions); err != nil {
		log.Print(err)
		return transactions, err
	}

	return transactions, nil
}

// Update a specific transaction and return the id of the updated transaction
func UpdateTransaction(session gocqlx.Session, id gocql.UUID, notes string, description string) (gocql.UUID, error) {
	transaction := Transaction{
		ID:          id,
		Notes:       notes,
		Description: description,
	}

	stmt, names := qb.Update("transaction").
		Set("description", "notes").Where(qb.Eq("id")).ToCql()

	q := session.Query(stmt, names).BindStruct(&transaction)
	if err := q.ExecRelease(); err != nil {
		return transaction.ID, err
	}
	return transaction.ID, nil
}
