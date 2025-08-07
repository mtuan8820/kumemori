package repository

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"kumemori/internal/core/domain"
)

type SqliteRepository struct {
	db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) *SqliteRepository {
	return &SqliteRepository{db: db}
}

func (s *SqliteRepository) ReadDeck(id uint) (*domain.Deck, error) {
	var deck domain.Deck
	err := s.db.First(&deck, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deck with ID %d not found", id)
		}
		return nil, err
	}
	return &deck, nil

}

func (s *SqliteRepository) ReadDecks() ([]*domain.Deck, error) {
	var decks []*domain.Deck
	err := s.db.Find(&decks).Error
	if err != nil {
		return nil, fmt.Errorf("decks not found: %v", err)
	}
	return decks, nil
}

func (s *SqliteRepository) SaveDeck(deck domain.Deck) error {
	return nil
}

// InitDb opens a SQLite connection and auto-migrates the Card and Deck models.
// Returns the DB instance or an error if migration fails.
func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Card{}, &domain.Deck{})
	if err != nil {
		return nil, err
	}

	log.Println("Db migrated successfully")
	return db, nil
}
