package service

import (
	"context"
	"fmt"
	"kumemori/internal/domain/model"
	"kumemori/internal/domain/repo"
)

// make sure DeckService implements IDeckService
var _ IDeckService = (*DeckService)(nil)

// DeckService handle business logic for Deck and Card entity
type DeckService struct {
	Repository repo.DeckRepo
}

func NewDeckService(repository repo.DeckRepo) *DeckService {
	return &DeckService{
		Repository: repository,
	}
}

func (s *DeckService) CreateDeck(ctx context.Context, name string, cards []*model.Card) (*model.Deck, error) {
	deck, err := model.NewDeck(name)
	if err != nil {
		return nil, fmt.Errorf("invalid deck data: %w", err)
	}

	// persist the entity (create only)
	if err := s.Repository.Create(ctx, deck); err != nil {
		return nil, fmt.Errorf("failed to create deck: %w", err)
	}

	return deck, nil
}

func (s *DeckService) GetDecks(ctx context.Context) ([]*model.Deck, error) {
	return s.Repository.FindAll(ctx)
}

func (s *DeckService) Delete(ctx context.Context, id uint) error {
	return s.Repository.Delete(ctx, id)
}

func (s *DeckService) Save(ctx context.Context, deck *model.Deck) error {
	return s.Repository.Update(ctx, deck)
}

func (s *DeckService) FindById(ctx context.Context, id uint) (*model.Deck, error) {
	return s.Repository.FindByID(ctx, id)
}

// CRUD card

func (s *DeckService) AddCard(ctx context.Context, deckID uint, card model.Card) error {
	deck, err := s.Repository.FindByID(ctx, deckID)
	if err != nil {
		return err
	}

	card.DeckID = deckID

	if err := deck.AddCard(card); err != nil {
		return err
	}

	if err := s.Repository.Update(ctx, deck); err != nil {
		return err
	}

	return nil
}

func (s *DeckService) DeleteCard(ctx context.Context, deckID uint, cardID uint) error {
	deck, err := s.Repository.FindByID(ctx, deckID)
	if err != nil {
		return err
	}

	if err := deck.RemoveCard(cardID); err != nil {
		return err
	}

	if err := s.Repository.Update(ctx, deck); err != nil {
		return err
	}

	return nil
}

func (s *DeckService) UpdateCard(ctx context.Context, deckID uint, cardID uint, front string, back string) error {
	deck, err := s.Repository.FindByID(ctx, deckID)
	if err != nil {
		return err
	}

	if err := deck.UpdateCard(cardID, front, back); err != nil {
		return err
	}

	if err := s.Repository.Update(ctx, deck); err != nil {
		return err
	}

	return nil
}

// FindAllCards retrieves all cards for the given deckID.
func (s *DeckService) FindAllCards(ctx context.Context, deckID uint) ([]*model.Card, error) {
	deck, err := s.Repository.FindByID(ctx, deckID)
	if err != nil {
		return nil, err
	}

	return deck.Cards(), nil
}

func (s *DeckService) Update(ctx context.Context, deckID uint, name string, cards []model.Card, actions []string) error {

	deck, err := s.Repository.FindByID(ctx, deckID)
	if err != nil {
		return fmt.Errorf("deck not found: %w", err)
	}

	if deck.Name != name {
		if err := deck.Rename(name); err != nil {
			return err
		}
	}

	for index, card := range cards {
		switch actions[index] {
		case "not changed":
			continue
		case "add":
			if _, err := deck.CreateCard(card.Front, card.Back); err != nil {
				return err
			}
		case "delete":
			if err := deck.RemoveCard(card.ID); err != nil {
				return err
			}
		case "update":
			if err := deck.UpdateCard(card.ID, card.Front, card.Back); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid action: %s", actions[index])
		}
	}

	if err := s.Repository.Update(ctx, deck); err != nil {
		return fmt.Errorf("failed to save deck: %w", err)
	}

	return nil
}
