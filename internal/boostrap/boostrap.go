package boostrap

import (
	"fmt"
	"kumemori/internal/adapter/repository"
	"kumemori/internal/adapter/repository/sqlite"
	"kumemori/internal/core/domain/service"
)

type AppDependencies struct {
	DeckService *service.DeckService
}

func InitApp() (*AppDependencies, error) {
	db, err := repository.InitDb()
	if err != nil {
		return nil, fmt.Errorf("init db failed: %w", err)
	}

	deckRepo := sqlite.NewDeckRepo(db)

	deckService := service.NewDeckService(deckRepo)

	return &AppDependencies{
		DeckService: deckService,
	}, nil
}
