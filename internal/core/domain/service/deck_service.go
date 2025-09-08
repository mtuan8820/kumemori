package service

import (
	"kumemori/internal/core/domain/entity"
	"kumemori/internal/core/ports"
)

type DeckService struct {
	repo ports.DeckRepository
}

func NewDeckService(repo ports.DeckRepository) *DeckService {
	return &DeckService{repo: repo}
}

func (s *DeckService) CreateDeck(name string) (*entity.Deck, error) {
	deck := &entity.Deck{
		Name: name,
	}

	if err := s.repo.CreateDeck(*deck); err != nil {
		return nil, err
	}

	return deck, nil
}

func (s *DeckService) GetDeck(id uint) (*entity.Deck, error) {
	return s.repo.ReadDeck(id)
}

func (s *DeckService) ListDecks() ([]*entity.Deck, error) {
	return s.repo.ReadDecks()
}

func (s *DeckService) DeleteDeck(id uint) error {
	return s.repo.DeleteDeck(id)
}

func (s *DeckService) UpdateDeck(id uint, name string) error {
	return s.repo.UpdateDeck(id, name)
}
