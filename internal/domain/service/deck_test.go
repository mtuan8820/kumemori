package service_test

import (
	"context"
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

func (m *MockDeckRepo) Create(ctx context.Context, deck *model.Deck) error {
	args := m.Called(ctx, deck)
	return args.Error(0)
}

func (m *MockDeckRepo) Update(ctx context.Context, deck *model.Deck) error {
	args := m.Called(ctx, deck)
	return args.Error(0)
}

func (m *MockDeckRepo) FindByID(ctx context.Context, id uint) (*model.Deck, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Deck), args.Error(1)
}

func (m *MockDeckRepo) FindAll(ctx context.Context) ([]*model.Deck, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Deck), args.Error(1)
}

func (m *MockDeckRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockDeckRepo) SaveCard(ctx context.Context, card *model.Card) error {
	args := m.Called(ctx, card)
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

	mockRepo.On("FindByID", mock.Anything, uint(1)).Return((*model.Deck)(nil), errors.New("not found"))

	err := svc.Update(context.TODO(), 1, "NewName", nil, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deck not found")
	mockRepo.AssertExpectations(t)
}

func TestUpdate_AddCard(t *testing.T) {
	mockRepo := new(MockDeckRepo)
	svc := service.NewDeckService(mockRepo)

	deck := mustNewDeck("Deck1")
	mockRepo.On("FindByID", mock.Anything, uint(1)).Return(deck, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Deck")).Return(nil)

	card := model.Card{Front: "Q", Back: "A"}

	err := svc.Update(context.TODO(), 1, "Deck1", []model.Card{card}, []string{"add"})

	assert.NoError(t, err)
	assert.Len(t, deck.Cards(), 1)
	assert.Equal(t, "Q", deck.Cards()[0].Front)
	mockRepo.AssertExpectations(t)
}
