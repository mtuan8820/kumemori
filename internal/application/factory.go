package application

import (
	"context"
	"kumemori/internal/application/deck"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

// Factory provides methods to create application use cases
type Factory struct {
	deckService service.IDeckService
	txFactory   repo.TransactionFactory
}

// NewFactory creates a new application factory
func NewFactory(
	deckService service.IDeckService,
	txFactory repo.TransactionFactory,
) *Factory {
	return &Factory{
		deckService: deckService,
		txFactory:   txFactory,
	}
}

// UpdateDeckUseCase returns a new update deck use case
func (f *Factory) UpdateDeck(ctx context.Context, input any) (any, error) {
	usecase := deck.NewUpdateUseCase(f.deckService, f.txFactory)
	return usecase.Execute(ctx, input)
}
