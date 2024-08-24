package app

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/pawpawchat/core/config"
	"github.com/pawpawchat/core/internal/infrastructure/handler"
	"github.com/pawpawchat/core/internal/infrastructure/router"
	"github.com/pawpawchat/core/pkg/middleware"
	profile "github.com/pawpawchat/profile/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(ctx context.Context, config config.Environment) {
	srv := settingUpServer(config)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("http server is running", "addr", config.ServerAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("http server error = ", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		slog.Debug("initiate a server shutdown")
		if err := srv.Shutdown(shutdownctx); err != nil {
			slog.Error("shutdown", "error", err)
		}
	}()

	wg.Wait()
}

func settingUpServer(env config.Environment) *http.Server {
	httpServer := &http.Server{
		BaseContext: func(listener net.Listener) context.Context {
			return context.Background()
		},
		Addr:    env.ServerAddr,
		Handler: settingUpRouter(env),
	}
	return httpServer
}

func settingUpRouter(env config.Environment) *router.MuxRouter {
	profile := profileServiceConn(env.GRPCPeers.Profile())

	mux := mux.NewRouter()
	rcfg := &router.Config{
		Routes: []router.Route{
			{
				Path:    "/{username}",
				Methods: []string{"GET"},
				Handler: handler.GetProfileByUsernameHandler(profile),
			},
			{
				Path:    "/{username}",
				Methods: []string{"POST"},
				Handler: handler.CreateProfileHandler(profile),
			},
		},
		Middlewares: []func(http.Handler) http.Handler{
			middleware.LogMiddleware(mux),
		},
	}
	return router.NewMuxRouter(mux, rcfg)
}

func profileServiceConn(addr string) profile.ProfileServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return profile.NewProfileServiceClient(conn)
}
