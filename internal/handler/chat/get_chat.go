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

type chatGetter interface {
	GetChat(context.Context, *chatpb.GetChatRequest, ...grpc.CallOption) (*chatpb.GetChatResponse, error)
}

func GetChatHandler(creator chatGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request chatpb.GetChatRequest

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

		respPb, err := creator.GetChat(r.Context(), &request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().
			Created().
			Body(map[string]any{
				"chat":     convert.MustChatPb(respPb.Chat),
				"messages": convert.MustFromPb(respPb.Messages, convert.MustMessagePb),
			}).
			MustWrite(w)
	})
}
