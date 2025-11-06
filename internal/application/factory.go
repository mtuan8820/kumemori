package application

import (
	"context"
	"kumemori/internal/application/deck"
	"kumemori/internal/domain/model"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

// Factory provides methods to create application use cases
type Factory struct {
	ctx         context.Context
	deckService service.IDeckService
	txFactory   repo.TransactionFactory
}

// NewFactory creates a new application factory
func NewFactory(
	ctx context.Context,
	deckService service.IDeckService,
	txFactory repo.TransactionFactory,
) *Factory {
	return &Factory{
		ctx:         ctx,
		deckService: deckService,
		txFactory:   txFactory,
	}
}

func (f *Factory) CreateDeck(input any) (any, error) {
	usecase := deck.NewCreateUseCase(f.ctx, f.deckService, f.txFactory)
	return usecase.Execute(input)
}

func (f *Factory) NewCreateDeckInput(name string, cards []model.Card) *deck.CreateInput {
	return &deck.CreateInput{
		Name:  name,
		Cards: cards,
	}
}

func (f *Factory) GetAllDecks() (any, error) {
	usecase := deck.NewGetAllUseCase(f.ctx, f.deckService, f.txFactory)
	return usecase.Execute(nil)
}

func (f *Factory) GetCards(id uint) (any, error) {
	usecase := deck.NewGetCardsUseCase(f.ctx, f.deckService, f.txFactory)
	return usecase.Execute(id)
}

func (f *Factory) UpdateDeck(input any) (any, error) {
	usecase := deck.NewUpdateUseCase(f.ctx, f.deckService, f.txFactory)
	return usecase.Execute(input)
}

func (f *Factory) NewUpdateInput(ID uint, Name string, CurrLength int, CardsToUpdate []deck.UpdateCardInput) *deck.UpdateInput {
	return &deck.UpdateInput{
		ID:            ID,
		Name:          Name,
		CurrLength:    CurrLength,
		CardsToUpdate: CardsToUpdate,
	}
}

func (f *Factory) NewUpdateCardInput(ID uint, Front string, Back string, Action string) *deck.UpdateCardInput {
	return &deck.UpdateCardInput{
		ID:     ID,
		Front:  Front,
		Back:   Back,
		Action: Action,
	}
}

func (f *Factory) DeleteDeck(deckId uint) (any, error) {
	usecase := deck.NewDeleteUseCase(f.ctx, f.deckService, f.txFactory)
	return usecase.Execute(deckId)
}
