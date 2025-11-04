package deck

import (
	"context"
	"kumemori/internal/application/core"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

type GetAllUseCase struct {
	ctx context.Context
	*core.UseCaseHandler
	deckService service.IDeckService
}

func NewGetAllUseCase(ctx context.Context, deckService service.IDeckService, txFactory repo.TransactionFactory) *GetAllUseCase {
	return &GetAllUseCase{
		ctx:            ctx,
		UseCaseHandler: core.NewUseCaseHandler(txFactory),
		deckService:    deckService,
	}
}

func (uc *GetAllUseCase) Execute(input any) (any, error) {
	return uc.deckService.GetDecks(uc.ctx)
}
