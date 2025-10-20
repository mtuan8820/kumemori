package service_test

import (
	"errors"
	"kumemori/internal/domain/model"
	"kumemori/internal/domain/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeckRepo struct {
	mock.Mock
}

func (m *MockDeckRepo) Save(deck *model.Deck) error {
	args := m.Called(deck)
	return args.Error(0)
}

func (m *MockDeckRepo) FindByID(id uint) (*model.Deck, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Deck), args.Error(1)
}

func (m *MockDeckRepo) FindAll() ([]*model.Deck, error) {
	args := m.Called()
	return args.Get(0).([]*model.Deck), args.Error(1)
}

func (m *MockDeckRepo) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDeckRepo) SaveCard(card *model.Card) error {
	args := m.Called(card)
	return args.Error(0)
}

// --- Helper for model ---
func mustNewDeck(name string) *model.Deck {
	d, err := model.NewDeck(name)
	if err != nil {
		panic(err)
	}
	return d
}

// --- Tests ---
func TestUpdate_DeckNotFound(t *testing.T) {
	mockRepo := new(MockDeckRepo)
	svc := service.NewDeckService(mockRepo)

	mockRepo.On("FindByID", uint(1)).Return((*model.Deck)(nil), errors.New("not found"))

	err := svc.Update(1, "NewName", nil, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deck not found")
	mockRepo.AssertExpectations(t)
}

func TestUpdate_AddCard(t *testing.T) {
	mockRepo := new(MockDeckRepo)
	svc := service.NewDeckService(mockRepo)

	deck := mustNewDeck("Deck1")

	mockRepo.On("FindByID", uint(1)).Return(deck, nil)
	mockRepo.On("Save", mock.AnythingOfType("*model.Deck")).Return(nil)

	card := model.Card{Front: "Q", Back: "A"}
	err := svc.Update(1, "Deck1", []model.Card{card}, []string{"add"})

	assert.NoError(t, err)
	assert.Len(t, deck.Cards(), 1)
	assert.Equal(t, "Q", deck.Cards()[0].Front)
	mockRepo.AssertExpectations(t)
}
