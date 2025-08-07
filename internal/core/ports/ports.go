package ports

import "kumemori/internal/core/domain"

type DeckService interface {
	ReadDecks() ([]*domain.Deck, error)
	ReadDeck(id uint) (*domain.Deck, error)
}

type DeckRepository interface {
	ReadDecks() ([]*domain.Deck, error)
	ReadDeck(id uint) (*domain.Deck, error)
	CreateDeck(deck domain.Deck) error
	DeleteDeck(id uint) error
	UpdateDeck(deck domain.Deck) error
}

type CardRepository interface {
	CreateCard(card domain.Card) error
	ReadCardsByDeck(deckID uint) ([]*domain.Card, error)
	ReadCard(id uint) (*domain.Card, error)
	ReadCards() ([]*domain.Card, error)
	UpdateCard(card domain.Card) error
	DeleteCard(id uint) error
}
