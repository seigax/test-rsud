package backendskeleton

import (
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/config"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/repository"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/usecase"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/websocket"
)

type BackendSkeleton struct {
	Usecase      *usecase.Usecase
	JWT          *lib.JWT
	WebsocketHub *websocket.Hub
	GoogleOauth2 *lib.GoogleOauth2
}

func NewBackendSkeleton(websocketHub *websocket.Hub) BackendSkeleton {
	db, _ := config.NewPG()
	storage := config.NewStorage()
	smtpClient := config.NewSMTPClient()
	googleOauth2 := config.NewGoogleOauth2()
	redisCon := config.NewRedisConnection()

	repository := repository.NewRepository(db, &smtpClient, redisCon)
	usecase := usecase.NewUsecase(&repository, storage)
	jwt := lib.NewJWT()

	return BackendSkeleton{
		Usecase:      &usecase,
		JWT:          jwt,
		WebsocketHub: websocketHub,
		GoogleOauth2: googleOauth2,
	}
}
