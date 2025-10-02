package application

import (
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
func (f *Factory) UpdateDeckUseCase() *deck.UpdateUseCase {
	return deck.NewUpdateUseCase(f.deckService, f.txFactory)
}
