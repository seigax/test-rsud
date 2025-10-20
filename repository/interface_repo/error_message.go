package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type ErrorMessageRepository interface {
	CreateErrorMessage(ctx context.Context, errorMessage model.ErrorMessage) (model.ErrorMessage, error)
	GetErrorMessages(ctx context.Context, query request.GetErrorMessageQuery) ([]model.ErrorMessage, error)
	GetErrorMessageTotal(ctx context.Context, query request.GetErrorMessageQuery) (uint, error)
	UpdateErrorMessage(ctx context.Context, errorMessage model.ErrorMessage) (model.ErrorMessage, error)
	GetErrorMessage(ctx context.Context, query request.GetErrorMessageQuery) (model.ErrorMessage, error)
	DeleteErrorMessage(ctx context.Context, ID uint) error
}
