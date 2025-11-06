package sqlite

import (
	"context"
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

func (d *DeckRepo) Create(ctx context.Context, deck *model.Deck) error {
	// create deck first to obtain ID
	if err := d.db.WithContext(ctx).Create(deck).Error; err != nil {
		return fmt.Errorf("failed to create deck: %w", err)
	}

	// create cards
	cards := deck.Cards()
	for _, c := range cards {
		c.DeckID = deck.ID
		if err := d.db.WithContext(ctx).Create(c).Error; err != nil {
			return fmt.Errorf("failed to create card: %w", err)
		}
	}

	return nil
}

// Update the deck and synchronize its cards: upsert provided cards and delete removed ones
func (d *DeckRepo) Update(ctx context.Context, deck *model.Deck) error {
	db := d.db.WithContext(ctx)

	// persist deck by primary key
	if err := db.Save(deck).Error; err != nil {
		return fmt.Errorf("failed to update deck: %w", err)
	}

	// fetch existing cards
	var existing []*model.Card
	if err := db.Where("deck_id = ?", deck.ID).Find(&existing).Error; err != nil {
		return fmt.Errorf("failed to load existing cards: %w", err)
	}

	// build map for desired
	desired := deck.Cards()
	desiredByID := make(map[uint]*model.Card, len(desired))
	for _, c := range desired {
		if c.ID != 0 {
			desiredByID[c.ID] = c
		}
	}

	// delete removed cards (present in existing but not in desired)
	for _, ex := range existing {
		if _, ok := desiredByID[ex.ID]; !ok {
			if err := db.Delete(&model.Card{}, ex.ID).Error; err != nil {
				return fmt.Errorf("failed to delete removed card: %w", err)
			}
		}
	}

	// upsert desired cards
	for _, c := range desired {
		c.DeckID = deck.ID
		if c.ID == 0 {
			if err := db.Create(c).Error; err != nil {
				return fmt.Errorf("failed to create card: %w", err)
			}
		} else {
			// update only relevant fields explicitly to avoid zero-value skipping
			updates := map[string]interface{}{
				"front": c.Front,
				"back":  c.Back,
			}
			if err := db.Model(&model.Card{}).Where("id = ?", c.ID).Updates(updates).Error; err != nil {
				return fmt.Errorf("failed to update card: %w", err)
			}
		}
	}

	return nil
}

// find deck by ID (including cards)
func (d *DeckRepo) FindByID(ctx context.Context, id uint) (*model.Deck, error) {
	var deck model.Deck
	if err := d.db.WithContext(ctx).First(&deck, "id=?", id).Error; err != nil {
		return nil, err
	}

	var cards []*model.Card
	if err := d.db.WithContext(ctx).Where("deck_id = ?", id).Find(&cards).Error; err != nil {
		return nil, err
	}

	deck.LoadCards(cards)

	return &deck, nil
}

// find all decks (not include cards)
func (d *DeckRepo) FindAll(ctx context.Context) ([]*model.Deck, error) {
	var decks []*model.Deck
	err := d.db.WithContext(ctx).Find(&decks).Error
	if err != nil {
		return nil, fmt.Errorf("decks not found: %v", err)
	}
	return decks, nil
}

// delete a deck (also cascade delete its card)
func (d *DeckRepo) Delete(ctx context.Context, id uint) error {
	db := d.db.WithContext(ctx)

	if err := db.Where("deck_id = ?", id).Delete(&model.Card{}).Error; err != nil {
		return fmt.Errorf("failed to delete cards for deck %d: %w", id, err)
	}

	if err := db.Delete(&model.Deck{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete deck with id %d: %w", id, err)
	}

	return nil
}

func (d *DeckRepo) SaveCard(ctx context.Context, card *model.Card) error {
	var deck model.Deck
	if err := d.db.WithContext(ctx).First(&deck, "id = ?", card.DeckID).Error; err != nil {
		return fmt.Errorf("card must belong to existing deck: %w", err)
	}

	if err := d.db.WithContext(ctx).Save(card).Error; err != nil {
		return fmt.Errorf("failed to save card: %w", err)
	}

	return nil
}
