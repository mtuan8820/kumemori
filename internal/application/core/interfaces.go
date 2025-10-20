package core

import (
	"context"
	"fmt"
	"kumemori/internal/domain/repo"
	"kumemori/internal/util/errors"
)

type UseCase interface {
	// Execute processes the use case with the given input and returns the result or an error
	Execute(input any) (any, error)
}

type UseCaseHandler struct {
	TxFactory repo.TransactionFactory
}

func NewUseCaseHandler(txFactory repo.TransactionFactory) *UseCaseHandler {
	return &UseCaseHandler{
		TxFactory: txFactory,
	}
}

// ExecuteInTransaction executes the given function within a transaction
func (h *UseCaseHandler) ExecuteInTransaction(
	ctx context.Context,
	fn func(context.Context, repo.Transaction) (any, error),
) (any, error) {
	tx, err := h.TxFactory.NewTransaction(ctx, nil)

	if err != nil {
		return nil, errors.Wrapf(err, errors.ErrorTypeSystem, "failed to create transaction")
	}
	// committed := false
	// defer func() {
	// 	if !committed {
	// 		_ = tx.Rollback()
	// 	}
	// }()

	// Execute function within transaction
	result, err := fn(ctx, tx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			fmt.Println("rollback error:", rerr)
		}
		return nil, err
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, errors.Wrapf(err, errors.ErrorTypeSystem, "failed to commit transaction")
	}
	// committed = true

	return result, nil
}

// Input defines the base interface for all use case inputs
type Input interface {
	Validate() error
}

// BaseInput provides common input validation functionality
type BaseInput struct{}

// Validate performs basic validation on the input
func (b *BaseInput) Validate() error {
	return nil
}

// Output defines the base interface for all use case outputs
type Output interface {
	GetStatus() string
}

// BaseOutput provides common output functionality
type BaseOutput struct {
	Status string `json:"status,omitempty"`
}

// GetStatus returns the status of the output
func (o *BaseOutput) GetStatus() string {
	return o.Status
}

// NewSuccessOutput creates a new success output
func NewSuccessOutput() *BaseOutput {
	return &BaseOutput{Status: "success"}
}

// ValidationError returns a validation error with the given message and details
func ValidationError(message string, details map[string]any) error {
	return errors.NewValidationError(message, nil).WithDetails(details)
}

// NotFoundError returns a not found error with the given message
func NotFoundError(message string) error {
	return errors.New(errors.ErrorTypeNotFound, message)
}

// BusinessError returns a business error with the given message
func BusinessError(message string) error {
	return errors.New(errors.ErrorTypeBusiness, message)
}
