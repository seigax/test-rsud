package usecase

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/config"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	interface_repo "gitlab.com/erloom.id/libraries/go/backend-skeleton/repository/interface_repo"
	mock_repo "gitlab.com/erloom.id/libraries/go/backend-skeleton/repository/mock"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/storage"
)

var strg = config.NewLocalStorage()
var m = &mock_repo.MockRepository{}
var u = NewUsecase(m, strg)

type Usecase struct {
	repo    interface_repo.BaseRepository
	storage storage.Storage
}

func NewUsecase(repo interface_repo.BaseRepository, storage storage.Storage) Usecase {
	return Usecase{repo: repo, storage: storage}
}

func LogError(ctx context.Context, message string, err error) {
	logger.Error(ctx, message, map[string]interface{}{
		"error": err,
		"tags":  []string{"gorm"},
	})
}
