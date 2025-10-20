package service

import (
	"context"
	"kumemori/internal/domain/model"
)

// IDeckService define interface for deck service
type IDeckService interface {
	CreateDeck(ctx context.Context, name string, cards []*model.Card) (*model.Deck, error)
	GetDecks(ctx context.Context) ([]*model.Deck, error)
	Delete(ctx context.Context, id uint) error
	Save(ctx context.Context, deck *model.Deck) error
	FindById(ctx context.Context, id uint) (*model.Deck, error)
	Update(ctx context.Context, deckID uint, name string, cards []model.Card, actions []string) error

	AddCard(ctx context.Context, deckID uint, card model.Card) error
	DeleteCard(ctx context.Context, deckID uint, cardID uint) error
	UpdateCard(ctx context.Context, deckID uint, cardID uint, front string, back string) error
	FindAllCards(ctx context.Context, deckID uint) ([]*model.Card, error)
}
