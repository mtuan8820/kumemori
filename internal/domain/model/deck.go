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

	//internal fields
	cards []*Card
}

// NewDeck creates a new deck entity with validation
func NewDeck(name string) (*Deck, error) {
	if name == "" {
		return nil, ErrEmptyDeckName
	}

	deck := &Deck{
		Name:          name,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		NewCardLimit:  200, //temporally set equal 200
		ReviewLimit:   0,   //temporally set equal 0
		LastStudiedAt: time.Time{},
		cards:         make([]*Card, 0),
	}

	return deck, nil
}

// ---- Aggregate behavior methods ----

func (d *Deck) CreateCard(front string, back string) (*Card, error) {
	card := Card{
		DeckID:       d.ID,
		Front:        front,
		Back:         back,
		CreatedAt:    time.Now(),
		Repetitions:  0,
		Lapses:       0,
		EaseFactor:   0,
		Interval:     0,
		Due:          time.Now(),
		LastReviewed: time.Time{},
	}

	// Add to collection nếu Deck giữ cards trong memory
	d.cards = append(d.cards, &card)

	return &card, nil
}

// AddCard adds a card to the deck
func (d *Deck) AddCard(card Card) error {
	if d.NewCardLimit > 0 && len(d.cards) >= d.NewCardLimit {
		return errors.New("deck has reached new card limit")
	}
	card.DeckID = d.ID
	card.CreatedAt = time.Now()
	d.cards = append(d.cards, &card)
	return nil
}

// UpdateCard updates the content of a card
func (d *Deck) UpdateCard(cardID uint, front, back string) error {
	for i, c := range d.cards {
		if c.ID == cardID {
			d.cards[i].Front = front
			d.cards[i].Back = back
			return nil
		}
	}
	return errors.New("card not found")
}

// RemoveCard removes a card by ID
func (d *Deck) RemoveCard(cardID uint) error {
	for i, c := range d.cards {
		if c.ID == cardID {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return nil
		}
	}
	return errors.New("card not found")
}

// FindCard returns a pointer to a card by ID
func (d *Deck) FindCard(cardID uint) (*Card, error) {
	for i, c := range d.cards {
		if c.ID == cardID {
			return d.cards[i], nil
		}
	}
	return nil, errors.New("card not found")
}

func (d *Deck) UpdateName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	d.Name = name
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func (d *Deck) Cards() []*Card {
	cardsCopy := make([]*Card, len(d.cards))
	copy(cardsCopy, d.cards)
	return cardsCopy
}
