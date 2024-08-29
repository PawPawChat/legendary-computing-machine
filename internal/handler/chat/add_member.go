package chat

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	chatpb "github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/core/pkg/validation"
	"google.golang.org/grpc"
)

type membersAdder interface {
	AddMember(context.Context, *chatpb.AddMemberRequest, ...grpc.CallOption) (*chatpb.AddMemberResponse, error)
}

func AddChatMembersHandler(provider membersAdder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request chatpb.AddMemberRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		var err error
		id := mux.Vars(r)["id"]
		request.ChatId, err = strconv.ParseInt(id, 0, 10)
		if err != nil {
			response.Json().BadRequest().BadRequest().Body(map[string]any{"error": "incorrect chat id val=" + id}).MustWrite(w)
			return
		}

		if emfs := validation.GetEmptyFields(&request); len(emfs) != 0 {
			response.WriteMissingFieldsError(w, emfs)
			return
		}

		member, err := provider.AddMember(r.Context(), &request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().Created().Body(member).MustWrite(w)
	})
}
