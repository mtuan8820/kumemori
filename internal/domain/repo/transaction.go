package repo

import (
	"context"
	"time"
)

// TransactionOptions represents options for a transaction
type TransactionOptions struct {
	// ReadOnly indicates if the transaction is read-only
	ReadOnly bool
	// Timeout specifies the transaction timeout duration
	Timeout time.Duration
	// Isolation specifies the isolation level
	Isolation string
}

// Transaction defines an interface for transactions
type Transaction interface {
	// Begin starts the transaction
	Begin() error
	// Commit commits the transaction
	Commit() error
	// Rollback rolls back the transaction
	Rollback() error
	// WithContext returns a new transaction with the given context
	WithContext(ctx context.Context) Transaction
}

// TransactionFactory defines an interface for creating transactions
type TransactionFactory interface {
	NewTransaction(ctx context.Context, opts any) (Transaction, error)
}
