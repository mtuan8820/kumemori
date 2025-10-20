package deck

import (
	"kumemori/internal/application/core"
	"kumemori/internal/domain/model"
)

const DeckNameTextLength = 100
const MaxCardTextLength = 1000
const DeckCardLimit = 100

// UpdateInput represents input for updating a deck
type UpdateInput struct {
	core.BaseInput
	ID            uint              `json:"id"`             // deck id
	Name          string            `json:"name,omitempty"` // deck name
	CurrLength    int               `json:"curLength"`      // current number of cards
	CardsToUpdate []UpdateCardInput `json:"cardsToUpdate,omitempty"`
}

type UpdateCardInput struct {
	ID     uint   `json:"id"`
	Front  string `json:"front"`
	Back   string `json:"back,omitempty"`
	Action string `json:"action"`
}

func (i *UpdateInput) Validate() error {
	if i.Name == "" {
		return core.ValidationError("name is required", map[string]any{
			"name": "required",
		})
	}

	for _, card := range i.CardsToUpdate {
		if err := validateCard(card.Front, card.Back); err != nil {
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

func (i *UpdateInput) ToDomain() ([]model.Card, []string) {

	cardsToUpdate := make([]model.Card, 0, len(i.CardsToUpdate))
	actions := make([]string, 0, len(i.CardsToUpdate))
	for _, card := range i.CardsToUpdate {
		cardsToUpdate = append(cardsToUpdate, model.Card{
			DeckID: i.ID,
			Front:  card.Front,
			Back:   card.Back,
		})
		actions = append(actions, card.Action)
	}

	return cardsToUpdate, actions
}
