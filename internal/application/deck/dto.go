package deck

import "kumemori/internal/application/core"

const DeckNameTextLength = 100
const MaxCardTextLength = 1000
const DeckCardLimit = 100

// EditInput represents input for updating a deck
type EditInput struct {
	core.BaseInput
	ID            string            `json:"id"`             // deck id
	Name          string            `json:"name,omitempty"` // deck name
	CurrLength    int               `json:"curLength"`      // current number of cards
	CardsToAdd    []CreateCardInput `json:"cardsToAdd,omitempty"`
	CardsToUpdate []UpdateCardInput `json:"cardsToUpdate,omitempty"`
	CardsToDelete []uint            `json:"cardsToDelete,omitempty"`
}

type CreateCardInput struct {
	Front string `json:"front"`
	Back  string `json:"back,omitempty"`
}

type UpdateCardInput struct {
	ID    uint    `json:"id"`
	Front *string `json:"front"`
	Back  *string `json:"back,omitempty"`
}

func (i *EditInput) Validate() error {
	if i.Name == "" {
		return core.ValidationError("name is required", map[string]any{
			"name": "required",
		})
	}

	if i.CurrLength-len(i.CardsToDelete)+len(i.CardsToUpdate) > DeckCardLimit {
		return core.ValidationError("exceed card limit", map[string]any{"card limit": "exceed"})
	}

	for _, card := range i.CardsToAdd {
		if err := validateCard(card.Front, card.Back); err != nil {
			return err
		}
	}

	for _, card := range i.CardsToUpdate {
		if err := validateCard(*card.Front, *card.Back); err != nil {
			return err
		}
	}

	return nil
}

func validateCard(front string, back string) error {
	if back == "" {
		return core.ValidationError("back is required", map[string]any{"back": "required"})
	}
	if len(front) > MaxCardTextLength {
		return core.ValidationError("card front too long", map[string]any{"front": "too long"})
	}
	if len(back) > MaxCardTextLength {
		return core.ValidationError("card back too long", map[string]any{"back": "too long"})
	}
	return nil
}
