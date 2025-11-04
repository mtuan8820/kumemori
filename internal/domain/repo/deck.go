package repo

import (
	"context"
	"kumemori/internal/domain/model"
)

type DeckRepo interface {
	Create(ctx context.Context, deck *model.Deck) error
	Update(ctx context.Context, deck *model.Deck) error
	FindAll(ctx context.Context) ([]*model.Deck, error)
	FindByID(ctx context.Context, id uint) (*model.Deck, error)
	Delete(ctx context.Context, id uint) error
	SaveCard(ctx context.Context, card *model.Card) error
}
