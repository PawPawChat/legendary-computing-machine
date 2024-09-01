package chat

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	chatpb "github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/internal/convert"
	"github.com/pawpawchat/core/internal/model"
	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/core/pkg/validation"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type messageSender interface {
	SendMessage(context.Context, *chatpb.SendMessageRequest, ...grpc.CallOption) (*chatpb.SendMessageResponse, error)
}

func SendChatMessageHandler(provider messageSender) http.Handler {
	type SendMessageRequest struct {
		model.Message
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request SendMessageRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		var err error
		chatID := mux.Vars(r)["id"]
		request.ChatID, err = strconv.ParseInt(chatID, 0, 10)
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

		if emfs := validation.GetEmptyFields(&request); len(emfs) != 0 {
			response.WriteMissingFieldsError(w, emfs)
			return
		}

		sentAt, err := time.Parse(time.RFC3339, request.SentAt)
		if err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		request_pb := chatpb.SendMessageRequest{
			ChatId:         request.ChatID,
			SenderId:       request.SenderID,
			SenderUsername: request.SenderUsername,
			Body:           request.Body,
			SentAt:         timestamppb.New(sentAt),
		}

		respPb, err := provider.SendMessage(r.Context(), &request_pb)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().Created().Body(convert.MustMessagePb(respPb.Message)).MustWrite(w)
	})
}
