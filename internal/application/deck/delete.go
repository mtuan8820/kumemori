package deck

import (
	"context"
	"kumemori/internal/application/core"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

type DeleteUseCase struct {
	ctx context.Context
	*core.UseCaseHandler
	deckService service.IDeckService
}

func NewDeleteUseCase(
	ctx context.Context,
	deckService service.IDeckService, txFactory repo.TransactionFactory,
) *DeleteUseCase {
	return &DeleteUseCase{
		ctx:            ctx,
		UseCaseHandler: core.NewUseCaseHandler(txFactory),
		deckService:    deckService,
	}
}

func (uc *DeleteUseCase) Execute(deckId uint) (any, error) {
	// execute in transaction
	result, err := uc.ExecuteInTransaction(uc.ctx, func(ctx context.Context, tx repo.Transaction) (any, error) {
		if err := uc.deckService.Delete(ctx, deckId); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
