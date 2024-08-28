package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pawpawchat/core/config"
	"github.com/pawpawchat/core/internal/handler"
	"github.com/pawpawchat/core/internal/router"
	"github.com/pawpawchat/core/pkg/middleware"
)

func settingUpRouter(env config.Environment) *router.MuxRouter {
	profile_client := profileServiceConn(env.Addr.Profile())

	mux := mux.NewRouter()
	rcfg := &router.Config{
		Routes: []router.Route{
			{
				Path:    "/api/profiles/{username}",
				Methods: []string{"GET"},
				Handler: handler.GetProfileByUsernameHandler(profile_client),
			},
			{
				Path:    "/api/profiles",
				Methods: []string{"GET"},
				Handler: handler.GetProfileByIdHandler(profile_client),
			},
			{
				Path:    "/api/profiles",
				Methods: []string{"POST"},
				Handler: handler.CreateProfileHandler(profile_client),
			},
			{
				Path:    "/api/profiles/{username}/avatars",
				Methods: []string{"POST"},
				Handler: handler.SetProfileAvatar(profile_client),
			},
		},
		Middlewares: []func(http.Handler) http.Handler{
			middleware.LogMiddleware(mux),
		},
	}
	return router.NewMuxRouter(mux, rcfg)
}
