package ports

import "kumemori/internal/core/domain"

type DeckService interface {
	ReadDecks() ([]*domain.Deck, error)
	ReadDeck(id uint) (*domain.Deck, error)
	SaveDeck(deck domain.Deck) error
}

type DeckRepository interface {
	ReadDecks() ([]*domain.Deck, error)
	ReadDeck(id uint) (*domain.Deck, error)
	SaveDeck(deck domain.Deck) error
}
