package deck

import (
	"context"
	"kumemori/internal/application/core"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

type GetCardsUseCase struct {
	ctx context.Context
	*core.UseCaseHandler
	deckService service.IDeckService
}

func NewGetCardsUseCase(
	ctx context.Context, deckService service.IDeckService, txFactory repo.TransactionFactory,
) *GetCardsUseCase {
	return &GetCardsUseCase{
		UseCaseHandler: core.NewUseCaseHandler(txFactory),
		deckService:    deckService,
	}
}

func (uc *GetCardsUseCase) Execute(id uint) (any, error) {
	return uc.deckService.FindAllCards(uc.ctx, id)
}
