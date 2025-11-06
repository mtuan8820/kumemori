package model

import (
	"time"
)

const DeckNameTextLength = 100
const MaxCardTextLength = 1000
const DeckCardLimit = 100

// Card entity
type Card struct {
	ID           uint
	DeckID       uint `gorm:"not null;index"`
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
	if err := validateDeckName(name); err != nil {
		return nil, err
	}

	deck := &Deck{
		Name:          name,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		NewCardLimit:  DeckCardLimit, //temporally set equal 200
		ReviewLimit:   0,             //temporally set equal 0
		LastStudiedAt: time.Time{},
		cards:         make([]*Card, 0),
	}

	return deck, nil
}

func validateDeckName(name string) error {
	if name == "" {
		return ErrEmptyDeckName
	}

	if len(name) > DeckNameTextLength {
		return ErrDeckNameTooLong
	}

	return nil
}

// ---- Aggregate behavior methods ----
// CreateCard creates a new card entity and assign to deck
func (d *Deck) CreateCard(front string, back string) (*Card, error) {
	if err := d.validateCard(front, back); err != nil {
		return nil, err
	}

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

	d.cards = append(d.cards, &card)

	return &card, nil
}

// AddCard adds a card to the deck
func (d *Deck) AddCard(card Card) error {
	if d.NewCardLimit > 0 && len(d.cards) >= d.NewCardLimit {
		return ErrDeckCardLimit
	}
	card.DeckID = d.ID
	card.CreatedAt = time.Now()
	d.cards = append(d.cards, &card)
	return nil
}

// UpdateCard updates the content of a card
func (d *Deck) UpdateCard(cardID uint, front, back string) error {
	if cardID == 0 {
		return ErrInvalidID
	}

	if err := d.validateCard(front, back); err != nil {
		return err
	}

	for i, c := range d.cards {
		if c.ID == cardID {
			d.cards[i].Front = front
			d.cards[i].Back = back
			return nil
		}
	}
	return ErrCardNotFound
}

func (d *Deck) RemoveCard(cardID uint) error {
	if len(d.cards) == 1 {
		return ErrDeckHaveAtLeastOneCard
	}

	if cardID == 0 {
		return ErrInvalidID
	}

	for i, c := range d.cards {
		if c.ID == cardID {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return nil
		}
	}
	return ErrCardNotFound
}

// FindCard returns a pointer to a card by ID
func (d *Deck) FindCard(cardID uint) (*Card, error) {
	for i, c := range d.cards {
		if c.ID == cardID {
			return d.cards[i], nil
		}
	}
	return nil, ErrCardNotFound
}

func (d *Deck) Rename(name string) error {
	if err := validateDeckName(name); err != nil {
		return err
	}
	d.Name = name
	return nil
}

// Cards return a copy of cards belong to deck
func (d *Deck) Cards() []*Card {
	cardsCopy := make([]*Card, len(d.cards))
	copy(cardsCopy, d.cards)
	return cardsCopy
}

func (d *Deck) LoadCards(cards []*Card) {
	d.cards = make([]*Card, len(cards))
	copy(d.cards, cards)
}

func (d *Deck) validateCard(front string, back string) error {
	if front == "" {
		return ErrEmptyCardFront
	}

	if len(front) > MaxCardTextLength {
		return ErrCardFrontTooLong
	}

	if len(back) > MaxCardTextLength {
		return ErrCardBackTooLong
	}

	return nil
}
