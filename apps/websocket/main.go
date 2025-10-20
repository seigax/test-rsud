package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	backendskeleton "gitlab.com/erloom.id/libraries/go/backend-skeleton"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/config"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/handler"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/websocket"
)

func init() {
	godotenv.Load()
	config.ViperConfig()
	logger.Init()
}

func main() {
	websocketHub := websocket.NewHub()

	backendSkeleton := backendskeleton.NewBackendSkeleton(websocketHub)

	handler := handler.NewHandler(&backendSkeleton)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	hub := websocket.NewHub()
	go hub.Run()

	router := chi.NewRouter()

	router.Get("/healthz", handler.Healthz)

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	server := &http.Server{
		Addr:    viper.GetString("HOST") + ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
