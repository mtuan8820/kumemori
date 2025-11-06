package deck

import (
	"context"
	"encoding/json"
	"kumemori/internal/application/core"
	"kumemori/internal/domain/model"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

type CreateUseCase struct {
	ctx context.Context
	*core.UseCaseHandler
	deckService service.IDeckService
}

func NewCreateUseCase(
	ctx context.Context,
	deckService service.IDeckService, txFactory repo.TransactionFactory,
) *CreateUseCase {
	return &CreateUseCase{
		ctx:            ctx,
		UseCaseHandler: core.NewUseCaseHandler(txFactory),
		deckService:    deckService,
	}
}

func (uc *CreateUseCase) Execute(input any) (any, error) {
	var createInput CreateInput

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonBytes, &createInput); err != nil {
		return nil, err
	}

	// execute in transaction
	result, err := uc.ExecuteInTransaction(uc.ctx, func(ctx context.Context, tx repo.Transaction) (any, error) {
		ptrCards := make([]*model.Card, len(createInput.Cards))
		for i := range createInput.Cards {
			ptrCards[i] = &createInput.Cards[i]
		}

		deck, err := uc.deckService.CreateDeck(ctx, createInput.Name, ptrCards)
		if err != nil {
			return nil, err
		}

		return deck, nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
