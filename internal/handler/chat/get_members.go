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

type membersGetter interface {
	GetMembers(context.Context, *chatpb.GetMembersRequest, ...grpc.CallOption) (*chatpb.GetMembersResponse, error)
}

func GetChatMembersHandler(provider membersGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request chatpb.GetMembersRequest
		var err error
		chatid := mux.Vars(r)["id"]

		request.ChatId, err = strconv.ParseInt(chatid, 0, 10)
		if err != nil {
			response.Json().BadRequest().Body(map[string]any{"error": "missing chat id in path"}).MustWrite(w)
			return
		}

		members, err := provider.GetMembers(r.Context(), &request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().Created().Body(members).MustWrite(w)
	})
}
