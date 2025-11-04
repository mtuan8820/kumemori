package deck

// import (
// 	"context"
// 	"kumemori/internal/application/core"
// 	"kumemori/internal/domain/repo"
// 	"kumemori/internal/domain/service"
// )

// type CreateUseCase struct {
// 	*core.UseCaseHandler
// 	deckService service.IDeckService
// }

// func NewCreateUseCase(
// 	deckService service.IDeckService, txFactory repo.TransactionFactory,
// ) *CreateUseCase {
// 	return &CreateUseCase{
// 		UseCaseHandler: core.NewUseCaseHandler(txFactory),
// 		deckService:    deckService,
// 	}
// }

// func (uc *CreateUseCase) Execute(ctx context.Context, input any) (any, error) {
// 	// validate input
// 	createInput, ok := input.(*CreateInput)

// 	// execute in transaction

// 	return nil, nil

// }
