package repository

import (
	"fmt"

	"gorm.io/gorm"

	"kumemori/internal/core/domain"
)

type SqliteCardRepository struct {
	db *gorm.DB
}

func NewCardSqliteRepository(db *gorm.DB) *SqliteCardRepository {
	return &SqliteCardRepository{db: db}
}

func (s *SqliteCardRepository) ReadCardsByDeck(id uint) error {
	return nil
}

func (s *SqliteCardRepository) ReadCards() ([]*domain.Card, error) {
	var cards []*domain.Card
	err := s.db.Find(&cards).Error
	if err != nil {
		return nil, fmt.Errorf("cards not found: %v", err)
	}
	return cards, nil
}

func (s *SqliteCardRepository) CreateCard(card domain.Card) error {
	return nil
}

func (s *SqliteCardRepository) DeleteCard(id uint) error {
	return nil
}

func (s *SqliteCardRepository) UpdateCard(card domain.Card) error {
	return nil
}
