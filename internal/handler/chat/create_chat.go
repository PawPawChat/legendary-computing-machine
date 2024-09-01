package chat

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	chatpb "github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/internal/convert"
	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/core/pkg/validation"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type chatCreator interface {
	CreateChat(context.Context, *chatpb.CreateChatRequest, ...grpc.CallOption) (*chatpb.CreateChatResponse, error)
}

func CreateChatHandler(creator chatCreator) http.Handler {
	type createChatRequest struct {
		Title         string `json:"title"`
		OwnerID       int64  `json:"owner_id"`
		OwnerUsername string `json:"owner_username"`
		CreatedAt     string `json:"created_at"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request createChatRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.WriteParseBodyError(w, err)
			return
		}
		if emfs := validation.GetEmptyFields(&request); len(emfs) != 0 {
			response.WriteMissingFieldsError(w, emfs)
			return
		}

		addedAt, err := time.Parse(time.RFC3339, request.CreatedAt)
		if err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		requestpb := chatpb.CreateChatRequest{
			Title:         request.Title,
			OwnerId:       request.OwnerID,
			OwnerUsername: request.OwnerUsername,
			CreatedAt:     timestamppb.New(addedAt),
		}

		respPb, err := creator.CreateChat(r.Context(), &requestpb)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().Created().Body(convert.MustChatPb(respPb.Chat)).MustWrite(w)
	})
}
