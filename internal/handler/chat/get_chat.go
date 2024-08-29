package chat

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	chatpb "github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/pkg/response"
	"google.golang.org/grpc"
)

type chatGetter interface {
	GetChat(context.Context, *chatpb.GetChatRequest, ...grpc.CallOption) (*chatpb.GetChatResponse, error)
}

func GetChatHandler(creator chatGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request chatpb.GetChatRequest
		var err error
		chatid := mux.Vars(r)["id"]

		request.ChatId, err = strconv.ParseInt(chatid, 0, 10)
		if err != nil {
			response.Json().BadRequest().Body(map[string]any{"error": "missing chat id in path"}).MustWrite(w)
			return
		}

		chat, err := creator.GetChat(r.Context(), &request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().Created().Body(chat).MustWrite(w)
	})
}
