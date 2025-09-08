package ports

import "kumemori/internal/core/domain/entity"

type DeckRepository interface {
	ReadDecks() ([]*entity.Deck, error)
	ReadDeck(id uint) (*entity.Deck, error)
	CreateDeck(deck entity.Deck) error
	DeleteDeck(id uint) error
	UpdateDeck(id uint, name string) error
}

type CardRepository interface {
	CreateCard(card entity.Card) error
	ReadCardsByDeck(deckID uint) ([]*entity.Card, error)
	ReadCard(id uint) (*entity.Card, error)
	ReadCards() ([]*entity.Card, error)
	UpdateCard(card entity.Card) error
	DeleteCard(id uint) error
}
