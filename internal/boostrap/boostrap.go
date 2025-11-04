package boostrap

import (
	"context"
	"fmt"
	"kumemori/internal/adapter/repository"
	"kumemori/internal/adapter/repository/sqlite"
	"kumemori/internal/application"

	"kumemori/internal/domain/service"
)

type AppDependencies struct {
	DeckService *service.DeckService
	Factory     *application.Factory
}

func InitApp(ctx context.Context) (*AppDependencies, error) {
	db, err := sqlite.InitDb()
	if err != nil {
		return nil, fmt.Errorf("init db failed: %w", err)
	}

	deckRepo := sqlite.NewDeckRepo(db)

	deckService := service.NewDeckService(deckRepo)

	txFactory := repository.NewGormTransactionFactory(db)

	factory := application.NewFactory(ctx, deckService, txFactory)

	return &AppDependencies{
		DeckService: deckService,
		Factory:     factory,
	}, nil
}
