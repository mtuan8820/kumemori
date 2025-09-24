package model

import (
	"kumemori/internal/util/errors"
)

var (

	// Deck validation errors
	ErrEmptyDeckName   = errors.New(errors.ErrorTypeValidation, "deck name cannot be empty")
	ErrDeckNameTooLong = errors.New(errors.ErrorTypeValidation, "deck name is too long")
	ErrDeckNameTaken   = errors.New(errors.ErrorTypeConflict, "deck name already taken")

	// Card validation errors
	ErrEmptyCardFront     = errors.New(errors.ErrorTypeValidation, "card front cannot be empty")
	ErrCardDeckIDRequired = errors.New(errors.ErrorTypeValidation, "card's deckId must not be empty")
	ErrCardFrontTooLong   = errors.New(errors.ErrorTypeValidation, "card front is too long")
	ErrCardBackTooLong    = errors.New(errors.ErrorTypeValidation, "card back is too long")

	ErrInvalidID = errors.New(errors.ErrorTypeValidation, "invalid ID")
)
