package sqlite

import (
	"errors"
	"fmt"
	"kumemori/internal/domain/model"

	"gorm.io/gorm"
)

type DeckRepo struct {
	db *gorm.DB
}

func NewDeckRepo(db *gorm.DB) *DeckRepo {
	return &DeckRepo{db: db}
}

// create a new deck with cards
func (s *DeckRepo) Save(deck *model.Deck) error {
	if err := s.db.Create(&deck).Error; err != nil {
		return fmt.Errorf("failed to create deck: %w", err)
	}
	return nil
}

// find deck by ID (including cards)
func (s *DeckRepo) FindByID(id uint) (*model.Deck, error) {
	var deck model.Deck
	err := s.db.First(&deck, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deck with ID %d not found", id)
		}
		return nil, err
	}
	return &deck, nil
}

// find all decks (not include cards)
func (s *DeckRepo) FindAll() ([]*model.Deck, error) {
	var decks []*model.Deck
	err := s.db.Find(&decks).Error
	if err != nil {
		return nil, fmt.Errorf("decks not found: %v", err)
	}
	return decks, nil
}

// delete a deck (also cascade delete its card)
func (s *DeckRepo) Delete(id uint) error {
	if err := s.db.Delete(&model.Deck{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete deck with id %d: %w", id, err)
	}

	return nil
}

// func (s *DeckRepo) AddCard() error {
// 	return nil
// }

// func (s *DeckRepo) DeleteCard() error {
// 	return nil
// }
