package db

import "time"

// Transaction represents a transaction
// which can be a transfer, a card payment, etc.
type Transaction struct {
	ID string
	AccountID string
	CreatedAt time.Time
	Description string
	Amount int64
	Currency string
	Notes string
}