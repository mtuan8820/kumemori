package sqlite

import (
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
func (d *DeckRepo) Save(deck *model.Deck) error {
	if err := d.db.Save(&deck).Error; err != nil {
		return fmt.Errorf("failed to save deck: %w", err)
	}
	return nil
}

// find deck by ID (including cards)
func (d *DeckRepo) FindByID(id uint) (*model.Deck, error) {
	var deck model.Deck
	if err := d.db.Preload("Cards").First(&deck, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &deck, nil
}

// find all decks (not include cards)
func (d *DeckRepo) FindAll() ([]*model.Deck, error) {
	var decks []*model.Deck
	err := d.db.Find(&decks).Error
	if err != nil {
		return nil, fmt.Errorf("decks not found: %v", err)
	}
	return decks, nil
}

// delete a deck (also cascade delete its card)
func (d *DeckRepo) Delete(id uint) error {
	if err := d.db.Delete(&model.Deck{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete deck with id %d: %w", id, err)
	}

	return nil
}

func (d *DeckRepo) SaveCard(card *model.Card) error {
	var deck model.Deck
	if err := d.db.First(&deck, "id = ?", card.DeckID).Error; err != nil {
		return fmt.Errorf("card must belong to existing deck: %w", err)
	}

	if err := d.db.Save(card).Error; err != nil {
		return fmt.Errorf("failed to save card: %w", err)
	}

	return nil
}
