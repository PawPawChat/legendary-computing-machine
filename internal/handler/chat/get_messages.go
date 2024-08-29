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

type messageGetter interface {
	GetMessages(context.Context, *chatpb.GetMessagesRequest, ...grpc.CallOption) (*chatpb.GetMessagesResponse, error)
}

func GetChatMessagesHandler(provider messageGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request chatpb.GetMessagesRequest

		var err error
		chat_id := mux.Vars(r)["id"]
		request.ChatId, err = strconv.ParseInt(chat_id, 0, 10)
		if err != nil {
			response.Json().BadRequest().BadRequest().Body(map[string]any{"error": "incorrect chat id val=" + chat_id}).MustWrite(w)
			return
		}

		messages, err := provider.GetMessages(r.Context(), &request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().Created().Body(messages).MustWrite(w)
	})
}
