package profile

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/core/pkg/validation"
	profilepb "github.com/pawpawchat/profile/api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AddProfileAvatar(client profilepb.ProfileServiceClient) http.Handler {
	type SetProfileAvatarRequest struct {
		ProfileID int64  `json:"profile_id"`
		OrigURL   string `json:"orig_url"`
		AddedAt   string `json:"added_at"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request SetProfileAvatarRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		addedAt, err := time.Parse(time.RFC3339, request.AddedAt)
		if err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		if zerofs := validation.GetEmptyFields(addedAt); len(zerofs) != 0 {
			response.WriteMissingFieldsError(w, zerofs)
			return
		}

		pbReq := &profilepb.AddProfileAvatarRequest{
			ProfileId: request.ProfileID,
			OrigUrl:   request.OrigURL,
			AddedAt:   timestamppb.New(addedAt),
		}

		avatar, err := client.AddProfileAvatar(r.Context(), pbReq)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().OK().Body(avatar).MustWrite(w)
	})
}
