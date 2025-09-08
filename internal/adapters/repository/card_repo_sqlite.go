package repository

import (
	"errors"
	"fmt"

	"kumemori/internal/core/domain/entity"

	"gorm.io/gorm"
)

type SqliteCardRepository struct {
	db *gorm.DB
}

func NewCardSqliteRepository(db *gorm.DB) *SqliteCardRepository {
	return &SqliteCardRepository{db: db}
}

func (s *SqliteCardRepository) ReadCardsByDeck(id uint) ([]*entity.Card, error) {
	var cards []*entity.Card
	err := s.db.Where("deck_id = ?", id).Find(&cards).Error
	if err != nil {
		return nil, fmt.Errorf("failed to read cards for deck %d: %w", id, err)
	}
	return cards, nil
}

func (s *SqliteCardRepository) ReadCards() ([]*entity.Card, error) {
	var cards []*entity.Card
	err := s.db.Find(&cards).Error
	if err != nil {
		return nil, fmt.Errorf("cards not found: %v", err)
	}
	return cards, nil
}

func (s *SqliteCardRepository) CreateCard(card entity.Card) error {
	result := s.db.Create(&card)
	if result.Error != nil {
		return fmt.Errorf("failed to create card: %w", result.Error)
	}
	return nil
}

func (s *SqliteCardRepository) DeleteCard(id uint) error {
	result := s.db.Delete(&entity.Card{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete card: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("card with ID %d not found", id)
	}
	return nil
}

func (s *SqliteCardRepository) UpdateCard(card entity.Card) error {
	result := s.db.Model(&entity.Card{}).
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

func (s *SqliteCardRepository) ReadCard(id uint) (*entity.Card, error) {
	var card entity.Card

	if err := s.db.First(&card, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("card with ID %d not found", id)
		}
		return nil, err
	}
	return &card, nil
}
