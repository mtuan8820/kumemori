package deck

import (
	"context"
	"kumemori/internal/application/core"
	"kumemori/internal/domain/repo"
	"kumemori/internal/domain/service"
)

type UpdateUseCase struct {
	*core.UseCaseHandler
	deckService service.IDeckService
}

func NewUpdateUseCase(
	deckService service.IDeckService, txFactory repo.TransactionFactory,
) *UpdateUseCase {
	return &UpdateUseCase{
		UseCaseHandler: core.NewUseCaseHandler(txFactory),
		deckService:    deckService,
	}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, input any) (any, error) {
	// covert and validate input
	updateInput, ok := input.(*UpdateInput)
	if !ok {
		return nil, core.ValidationError("invalid input type", nil)
	}

	if err := updateInput.Validate(); err != nil {
		return nil, err
	}

	// convert updateInput to card domain entity
	cardsToAdd, cardsToUpdate := updateInput.ToDomain()

	// execute in transaction
	result, err := uc.ExecuteInTransaction(ctx, func(ctx context.Context, tx repo.Transaction) (any, error) {
		// call domain service to update deck
		err := uc.deckService.Update(updateInput.ID, updateInput.Name, cardsToAdd, updateInput.CardsToDelete, cardsToUpdate)
		if err != nil {
			return nil, err
		}

		// get the updated deck
		_, err = uc.deckService.FindById(updateInput.ID)
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
