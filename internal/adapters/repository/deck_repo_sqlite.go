package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"kumemori/internal/core/domain/entity"
)

type SqliteDeckRepository struct {
	db *gorm.DB
}

func NewDeckSqliteRepository(db *gorm.DB) *SqliteDeckRepository {
	return &SqliteDeckRepository{db: db}
}

func (s *SqliteDeckRepository) ReadDeck(id uint) (*entity.Deck, error) {
	var deck entity.Deck
	err := s.db.First(&deck, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deck with ID %d not found", id)
		}
		return nil, err
	}
	return &deck, nil
}

func (s *SqliteDeckRepository) ReadDecks() ([]*entity.Deck, error) {
	var decks []*entity.Deck
	err := s.db.Find(&decks).Error
	if err != nil {
		return nil, fmt.Errorf("decks not found: %v", err)
	}
	return decks, nil
}

func (s *SqliteDeckRepository) CreateDeck(deck entity.Deck) error {
	if err := s.db.Create(&deck).Error; err != nil {
		return fmt.Errorf("failed to create deck: %w", err)
	}
	return nil
}

func (s *SqliteDeckRepository) DeleteDeck(id uint) error {
	if err := s.db.Delete(&entity.Deck{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete deck with id %d: %w", id, err)
	}
	return nil
}

func (s *SqliteDeckRepository) UpdateDeck(deck entity.Deck) error {
	result := s.db.Model(&deck).Select("Name", "Description").Updates(&deck)

	if result.Error != nil {
		return fmt.Errorf("failed to update deck: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("deck with ID %d not found", deck.ID)
	}
	return nil
}
