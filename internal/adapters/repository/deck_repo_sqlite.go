package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"kumemori/internal/core/domain"
)

type SqliteDeckRepository struct {
	db *gorm.DB
}

func NewDeckSqliteRepository(db *gorm.DB) *SqliteDeckRepository {
	return &SqliteDeckRepository{db: db}
}

func (s *SqliteDeckRepository) ReadDeck(id uint) (*domain.Deck, error) {
	var deck domain.Deck
	err := s.db.First(&deck, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deck with ID %d not found", id)
		}
		return nil, err
	}
	return &deck, nil
}

func (s *SqliteDeckRepository) ReadDecks() ([]*domain.Deck, error) {
	var decks []*domain.Deck
	err := s.db.Find(&decks).Error
	if err != nil {
		return nil, fmt.Errorf("decks not found: %v", err)
	}
	return decks, nil
}

func (s *SqliteDeckRepository) CreateDeck(deck domain.Deck) error {
	return nil
}

func (s *SqliteDeckRepository) DeleteDeck(id uint) error {
	return nil
}

func (s *SqliteDeckRepository) UpdateDeck(deck domain.Deck) error {
	return nil
}
