package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pawpawchat/core/config"
	"github.com/pawpawchat/core/internal/handler/chat"
	"github.com/pawpawchat/core/internal/handler/profile"
	"github.com/pawpawchat/core/internal/router"
	"github.com/pawpawchat/core/pkg/middleware"
)

func settingUpRouter(env config.Environment) *router.MuxRouter {
	profile_client := profileServiceConn(env.Addr.Profile())
	chat_client := chatServiceConn(env.Addr.Chat())
	mux := mux.NewRouter()

	rcfg := &router.Config{
		Routes: []router.Route{
			{
				Path:    "/api/profiles",
				Methods: []string{"POST"},
				Handler: profile.CreateProfileHandler(profile_client),
			},
			{
				Path:    "/api/profiles",
				Methods: []string{"GET"},
				Handler: profile.GetProfileByIdHandler(profile_client),
			},
			{
				Path:    "/api/profiles/{username}",
				Methods: []string{"GET"},
				Handler: profile.GetProfileByUsernameHandler(profile_client),
			},
			{
				Path:    "/api/profiles/{username}/avatars",
				Methods: []string{"POST"},
				Handler: profile.AddProfileAvatar(profile_client),
			},
			{
				Path:    "/api/chats",
				Methods: []string{"POST"},
				Handler: chat.CreateChatHandler(chat_client),
			},
			{
				Path:    "/api/chats/{id}",
				Methods: []string{"GET"},
				Handler: chat.GetChatHandler(chat_client),
			},
			{
				Path:    "/api/chats/{id}/members",
				Methods: []string{"POST"},
				Handler: chat.AddChatMembersHandler(chat_client),
			},
			{
				Path:    "/api/chats/{id}/members",
				Methods: []string{"GET"},
				Handler: chat.GetChatMembersHandler(chat_client),
			},
			{
				Path:    "/api/chats/{id}/messages",
				Methods: []string{"POST"},
				Handler: chat.SendChatMessageHandler(chat_client),
			},
			{
				Path:    "/api/chats/{id}/messages",
				Methods: []string{"GET"},
				Handler: chat.GetChatMessagesHandler(chat_client),
			},
		},
		Middlewares: []func(http.Handler) http.Handler{
			middleware.LogMiddleware(mux),
		},
	}
	return router.NewMuxRouter(mux, rcfg)
}
