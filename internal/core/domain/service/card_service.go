package service

import (
	"kumemori/internal/core/domain/entity"
	"kumemori/internal/core/ports"
)

type CardService struct {
	repo ports.CardRepository
}

func NewCardService(repo ports.CardRepository) *CardService {
	return &CardService{repo: repo}
}

func (s *CardService) CreateCard(front string, back string) (*entity.Card, error) {
	card := &entity.Card{
		Front: front,
		Back:  back,
	}

	if err := s.repo.CreateCard(*card); err != nil {
		return nil, err
	}

	return card, nil
}

func (s *CardService) GetCard(id uint) (*entity.Card, error) {
	return s.repo.ReadCard(id)
}

func (s *CardService) ListCards() ([]*entity.Card, error) {
	return s.repo.ReadCards()
}

func (s *CardService) DeleteCard(id uint) error {
	return s.repo.DeleteCard(id)
}

func (s *CardService) ReadCardsByDeck(deckId uint) ([]*entity.Card, error) {
	return s.repo.ReadCardsByDeck(deckId)
}
