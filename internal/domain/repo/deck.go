package repo

import "kumemori/internal/domain/model"

type DeckRepo interface {
	Save(deck *model.Deck) error
	FindAll() ([]*model.Deck, error)
	FindByID(id uint) (*model.Deck, error)
	Delete(id uint) error
	SaveCard(card *model.Card) error
}
