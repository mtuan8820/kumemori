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

func (s *SqliteCardRepository) ReadCardsByDeck(id uint) ([]*domain.Card, error) {
	var cards []*domain.Card
	err := s.db.Where("deck_id = ?", id).Find(&cards).Error
	if err != nil {
		return nil, fmt.Errorf("failed to read cards for deck %d: %w", id, err)
	}
	return cards, nil
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
	result := s.db.Create(&card)
	if result.Error != nil {
		return fmt.Errorf("failed to create card: %w", result.Error)
	}
	return nil
}

func (s *SqliteCardRepository) DeleteCard(id uint) error {
	result := s.db.Delete(&domain.Card{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete card: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("card with ID %d not found", id)
	}
	return nil
}

func (s *SqliteCardRepository) UpdateCard(card domain.Card) error {
	result := s.db.Model(&domain.Card{}).
		Where("id = ?", card.ID).
		Updates(card)

	if result.Error != nil {
		return fmt.Errorf("failed to update card: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("card with ID %d not found", card.ID)
	}
	return nil
}
