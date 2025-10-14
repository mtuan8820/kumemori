package service

import "kumemori/internal/domain/model"

// IDeckService define interface for deck service
type IDeckService interface {
	CreateDeck(name string, cards []*model.Card) (*model.Deck, error)
	GetDecks() ([]*model.Deck, error)
	Delete(id uint) error
	Save(deck *model.Deck) error
	FindById(id uint) (*model.Deck, error)
	Update(deckID uint, name string, cards []model.Card, actions []string) error

	AddCard(deckID uint, card model.Card) error
	DeleteCard(deckID uint, cardID uint) error
	UpdateCard(deckID uint, cardID uint, front string, back string) error
	FindAllCards(deckID uint) ([]*model.Card, error)
}
