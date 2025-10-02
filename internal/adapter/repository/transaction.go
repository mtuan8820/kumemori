package repository

import (
	"context"
	"fmt"
	"kumemori/internal/domain/repo"
	"time"

	"gorm.io/gorm"
)

// DefaultTimeout defines the default context timeout for database operations
const DefaultTimeout = 30 * time.Second

// Transaction represents a database transaction
type Transaction struct {
	ctx     context.Context
	Session *gorm.DB
}

func (tx *Transaction) Begin() error {
	if tx == nil {
		return fmt.Errorf("invalid transaction")
	}

	if tx.Session == nil {
		return fmt.Errorf("invalid session")
	}

	tx.Session = tx.Session.Begin()
	if tx.Session.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Session.Error)
	}
	return nil
}

func (tx *Transaction) Commit() error {
	if tx != nil && tx.Session != nil {
		return tx.Session.Commit().Error
	}
	return nil
}

func (tx *Transaction) Rollback() error {
	if tx != nil && tx.Session != nil {
		return tx.Session.Rollback().Error
	}
	return nil
}

// WithContext returns a new transaction with the given context
func (tx *Transaction) WithContext(ctx context.Context) repo.Transaction {
	if tx == nil {
		return nil
	}
	newTx := *tx
	newTx.ctx = ctx
	if newTx.Session != nil {
		newTx.Session = newTx.Session.WithContext(ctx)
	}
	return &newTx
}

type gormTransactionFactory struct {
	db *gorm.DB
}

func (f *gormTransactionFactory) NewTransaction(ctx context.Context, opts any) (repo.Transaction, error) {
	return &Transaction{ctx: ctx, Session: f.db}, nil
}

func NewGormTransactionFactory(db *gorm.DB) repo.TransactionFactory {
	return &gormTransactionFactory{db: db}
}
