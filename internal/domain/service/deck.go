package service

import (
	"kumemori/internal/domain/model"
	"kumemori/internal/domain/repo"
	"time"
)

type DeckService struct {
	Repository repo.DeckRepo
}

func NewDeckService(repository repo.DeckRepo) *DeckService {
	return &DeckService{
		Repository: repository,
	}
}

func (s *DeckService) CreateDeck(name string, cards []*model.Card) (*model.Deck, error) {
	deck := &model.Deck{
		Name:  name,
		Cards: cards,
	}

	if err := s.Repository.Save(deck); err != nil {
		return nil, err
	}

	return deck, nil
}

func (s *DeckService) GetDecks() ([]*model.Deck, error) {
	return s.Repository.FindAll()
}

func (s *DeckService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

func (s *DeckService) Save(deck *model.Deck) error {
	return s.Repository.Save(deck)
}

func (s *DeckService) FindById(id uint) (*model.Deck, error) {
	return s.Repository.FindByID(id)
}

// CRUD card

func (s *DeckService) AddCard(deckID uint, card model.Card) error {
	deck, err := s.Repository.FindByID(deckID)
	if err != nil {
		return err
	}

	card.DeckID = deckID
	card.CreatedAt = time.Now()

	if err := deck.AddCard(card); err != nil {
		return err
	}

	if err := s.Repository.Save(deck); err != nil {
		return err
	}

	return nil
}

func (s *DeckService) DeleteCard(deckID uint, cardID uint) error {
	deck, err := s.Repository.FindByID(deckID)
	if err != nil {
		return err
	}

	if err := deck.RemoveCard(cardID); err != nil {
		return err
	}

	if err := s.Repository.Save(deck); err != nil {
		return err
	}

	return nil
}

func (s *DeckService) UpdateCard(deckID uint, cardID uint, front string, back string) error {
	deck, err := s.Repository.FindByID(deckID)
	if err != nil {
		return err
	}

	if err := deck.UpdateCard(cardID, front, back); err != nil {
		return err
	}

	if err := s.Repository.Save(deck); err != nil {
		return err
	}

	return nil
}

func (s *DeckService) FindAllCards(deckID uint) ([]*model.Card, error) {
	deck, err := s.Repository.FindByID(deckID)
	if err != nil {
		return nil, err
	}
	return deck.Cards, nil
}
