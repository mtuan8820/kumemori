package boostrap

import (
	"fmt"
	"kumemori/internal/adapters/repository"
	"kumemori/internal/core/domain/service"
)

type AppDependencies struct {
	DeckService *service.DeckService
	CardService *service.CardService
}

func InitApp() (*AppDependencies, error) {
	db, err := repository.InitDb()
	if err != nil {
		return nil, fmt.Errorf("init db failed: %w", err)
	}

	deckRepo := repository.NewDeckSqliteRepository(db)
	cardRepo := repository.NewCardSqliteRepository(db)

	deckService := service.NewDeckService(deckRepo)
	cardService := service.NewCardService(cardRepo)

	return &AppDependencies{
		DeckService: deckService,
		CardService: cardService,
	}, nil
}
