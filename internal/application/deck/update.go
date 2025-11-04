package deck

import (
	"context"
	"kumemori/internal/application/core"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"

	"github.com/mitchellh/mapstructure"
)

type UpdateUseCase struct {
	ctx context.Context
	*core.UseCaseHandler
	deckService service.IDeckService
}

// make sure UpdateUseCase implement UseCase interface
var _ core.UseCase = (*UpdateUseCase)(nil)

func NewUpdateUseCase(
	ctx context.Context, deckService service.IDeckService, txFactory repo.TransactionFactory,
) *UpdateUseCase {
	return &UpdateUseCase{
		ctx:            ctx,
		UseCaseHandler: core.NewUseCaseHandler(txFactory),
		deckService:    deckService,
	}
}

func (uc *UpdateUseCase) Execute(input any) (any, error) {
	var updateInput UpdateInput

	// Decode map[string]interface{} → struct
	if err := mapstructure.Decode(input, &updateInput); err != nil {
		return nil, core.ValidationError("invalid update input data", nil)
	}

	if err := updateInput.Validate(); err != nil {
		return nil, err
	}

	// convert updateInput to card domain entity
	cardsToUpdate, actions := updateInput.ToDomain()

	// execute in transaction
	result, err := uc.ExecuteInTransaction(uc.ctx, func(ctx context.Context, tx repo.Transaction) (any, error) {
		// call domain service to update deck
		err := uc.deckService.Update(uc.ctx, updateInput.ID, updateInput.Name, cardsToUpdate, actions)
		if err != nil {
			return nil, err
		}

		// get the updated deck
		_, err = uc.deckService.FindById(uc.ctx, updateInput.ID)
		if err != nil {
			return nil, err
		}

		//create output dto
		return nil, nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
