package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	backendskeleton "gitlab.com/erloom.id/libraries/go/backend-skeleton"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/config"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/handler"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/otel"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/websocket"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func init() {
	godotenv.Load()
	config.ViperConfig()
	logger.Init()
}

func main() {
	// clientBuilder := lib.NewMyHttpClientBuilder().
	// 	SetCtx(context.Background()).
	// 	SetUrl("https://localhost").
	// 	SetMethod(lib.GET).
	// 	SetVerifySsl(true).
	// 	SetVerifySslCaCrtPath("ssl/ca.crt").
	// 	SetVerifySslClientCrtPath("ssl/client.crt").
	// 	SetVerifySslClientKeyPath("ssl/client.key")

	// client := clientBuilder.Build()
	// client.Execute()

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := otel.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	websocketHub := websocket.NewHub()

	go websocketHub.Run()

	backendSkeleton := backendskeleton.NewBackendSkeleton(websocketHub)
	handler := handler.NewHandler(&backendSkeleton)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	router := chi.NewRouter()

	router.Get("/healthz", handler.Healthz)

	router.HandleFunc("/ws/{RoomID}", handler.ServeWS)

	router.Group(func(r chi.Router) {
		r.Use(handler.StandardMiddleware)
		r.Use(handler.PanicMiddlewares)
		r.Use(handler.OtelMiddleware)

		r.Route("/auth", func(r chi.Router) {
			// r.Post("/register-farmer", handler.RegisterFarmer)
			r.Post("/login", handler.Login)
			r.Post("/logout", handler.AuthorizeSessionTokenMiddleware(handler.Logout))
			// r.Post("/check-otp-login", handler.CheckOTPLoginFarmer)
			// r.Post("/get-farmer-by-phone", handler.GetFarmerPhone)
			// r.Post("/change-password", handler.AuthorizeSessionTokenFarmerMiddleware(handler.ChangePassword))
		})

		r.Route("/system", func(r chi.Router) {
			r.Route("/system-parameter", func(r chi.Router) {
				// r.Use(handler.MenuParameterSystemMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateSystemParameter))
				r.Get("/", handler.AuthorizeSessionTokenMiddleware(handler.GetSystemParameters))
				r.Get("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.GetSystemParameterDetail))
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateSystemParameter))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteSystemParameter))
			})

			r.Route("/error-message", func(r chi.Router) {
				// r.Use(handler.MenuErrorMessageMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateErrorMessage))
				r.Get("/", handler.AuthorizeSessionTokenMiddleware(handler.GetErrorMessages))
				r.Get("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.GetErrorMessageDetail))
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateErrorMessage))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteErrorMessage))
			})
		})

		r.Route("/security", func(r chi.Router) {
			r.Route("/menu", func(r chi.Router) {
				// r.Use(handler.MenuMenuMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateMenu))
				r.Get("/", handler.AuthorizeSessionTokenMiddleware(handler.GetMenus))
				r.Get("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.GetMenuDetail))
				r.Get("/{ID}/with-child", handler.AuthorizeSessionTokenMiddleware(handler.GetMenuDetailWithChild))
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateMenu))
				r.Put("/update-order", handler.AuthorizeSessionTokenMiddleware(handler.UpdateOrderMenu))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteMenu))
			})

			r.Route("/role", func(r chi.Router) {
				// r.Use(handler.MenuRoleMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateRole))
				r.Get("/", handler.AuthorizeSessionTokenMiddleware(handler.GetRoles))
				r.Get("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.GetRoleDetail))
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateRole))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteRole))
			})

			r.Route("/role-menu", func(r chi.Router) {
				// r.Use(handler.MenuRoleMiddleware)

				r.Get("/get-all-menu", handler.AuthorizeSessionTokenMiddleware(handler.GetMenusWithChild))
				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.SaveRoleMenu))
			})

			r.Route("/user", func(r chi.Router) {
				// r.Use(handler.MenuUserMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateUser))
				r.Get("/", handler.AuthorizeSessionTokenMiddleware(handler.GetUsers))
				r.Get("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.GetUserDetail))
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateUser))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteUser))
			})
		})

		r.Route("/master", func(r chi.Router) {
			r.Route("/province", func(r chi.Router) {
				// r.Use(handler.MenuRegionMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateProvince))
				r.Get("/", handler.GetProvinces)
				r.Get("/{ID}", handler.GetProvinceDetail)
				r.Get("/{ID}/tree", handler.GetProvinceDetailWithTree)
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateProvince))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteProvince))
			})

			r.Route("/city", func(r chi.Router) {
				// r.Use(handler.MenuRegionMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateCity))
				r.Get("/", handler.GetCitys)
				r.Get("/{ID}", handler.GetCityDetail)
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateCity))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteCity))
			})

			r.Route("/district", func(r chi.Router) {
				// r.Use(handler.MenuRegionMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateDistrict))
				r.Get("/", handler.GetDistricts)
				r.Get("/{ID}", handler.GetDistrictDetail)
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateDistrict))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteDistrict))
			})

			r.Route("/village", func(r chi.Router) {
				// r.Use(handler.MenuRegionMiddleware)

				r.Post("/", handler.AuthorizeSessionTokenMiddleware(handler.CreateVillage))
				r.Get("/", handler.GetVillages)
				r.Get("/{ID}", handler.GetVillageDetail)
				r.Put("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.UpdateVillage))
				r.Delete("/{ID}", handler.AuthorizeSessionTokenMiddleware(handler.DeleteVillage))
			})
		})

	})

	h := otelhttp.NewHandler(router, "/")

	server := &http.Server{
		Addr: viper.GetString("HOST") + ":" + viper.GetString("PORT"),
		// Handler: router,
		Handler: h,
	}

	// server.ListenAndServe()

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- server.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		log.Fatal(err)
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = server.Shutdown(context.Background())
}
