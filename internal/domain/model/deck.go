package model

import (
	"errors"
	"time"
)

// Card entity
type Card struct {
	ID           uint
	DeckID       uint
	Front        string
	Back         string
	CreatedAt    time.Time
	Repetitions  int
	Lapses       int
	EaseFactor   float64
	Interval     int
	Due          time.Time
	LastReviewed time.Time
}

// Deck entity & aggregate root
type Deck struct {
	ID            uint
	Name          string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	NewCardLimit  int
	ReviewLimit   int
	LastStudiedAt time.Time

	Cards []*Card
}

// ---- Aggregate behavior methods ----

// AddCard adds a card to the deck
func (d *Deck) AddCard(card Card) error {
	if d.NewCardLimit > 0 && len(d.Cards) >= d.NewCardLimit {
		return errors.New("deck has reached new card limit")
	}
	card.DeckID = d.ID
	card.CreatedAt = time.Now()
	d.Cards = append(d.Cards, &card)
	return nil
}

// UpdateCard updates the content of a card
func (d *Deck) UpdateCard(cardID uint, front, back string) error {
	for i, c := range d.Cards {
		if c.ID == cardID {
			d.Cards[i].Front = front
			d.Cards[i].Back = back
			return nil
		}
	}
	return errors.New("card not found")
}

// RemoveCard removes a card by ID
func (d *Deck) RemoveCard(cardID uint) error {
	for i, c := range d.Cards {
		if c.ID == cardID {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
			return nil
		}
	}
	return errors.New("card not found")
}

// FindCard returns a pointer to a card by ID
func (d *Deck) FindCard(cardID uint) (*Card, error) {
	for i, c := range d.Cards {
		if c.ID == cardID {
			return d.Cards[i], nil
		}
	}
	return nil, errors.New("card not found")
}
