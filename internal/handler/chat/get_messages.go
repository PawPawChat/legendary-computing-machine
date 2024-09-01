package chat

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	chatpb "github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/internal/convert"
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
		chatID := mux.Vars(r)["id"]
		request.ChatId, err = strconv.ParseInt(chatID, 0, 10)
		if err != nil {
			response.Json().
				BadRequest().
				Body(map[string]any{
					"error": map[string]any{
						"message": "cannot parse chat id",
						"value":   chatID,
					}}).
				MustWrite(w)
			return
		}

		respPb, err := provider.GetMessages(r.Context(), &request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().
			Created().
			Body(map[string]any{
				"chat_id":  respPb.ChatId,
				"messages": convert.MustFromPb(respPb.Messages, convert.MustMessagePb),
			}).
			MustWrite(w)
	})
}
