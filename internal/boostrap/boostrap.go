package boostrap

import (
	"fmt"
	"kumemori/internal/adapters/repository"
	"log"
)

func InitApp() error {
	db, err := repository.InitDb()
	if err != nil {
		return fmt.Errorf("init db failed: %w", err)
	}

	deckRepo := repository.NewSqliteRepository(db)

	decks, err := deckRepo.ReadDecks()
	if err != nil {
		log.Println("Deck not found:", err)
	} else {
		log.Println("Deck loaded:", len(decks))
	}
	return nil
}
