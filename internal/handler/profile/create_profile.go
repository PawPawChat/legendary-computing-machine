package profile

import (
	"encoding/json"
	"net/http"

	"github.com/pawpawchat/core/internal/convert"
	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/core/pkg/validation"
	profilepb "github.com/pawpawchat/profile/api/pb"
)

func CreateProfileHandler(client profilepb.ProfileServiceClient) http.Handler {
	type CreateProfileRequest struct {
		FirstName  string `json:"first_name"`
		SecondName string `json:"second_name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request CreateProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.WriteParseBodyError(w, err)
			return
		}
		if zerofs := validation.GetEmptyFields(request); len(zerofs) != 0 {
			response.WriteMissingFieldsError(w, zerofs)
			return
		}

		pbReq := &profilepb.CreateProfileRequest{
			FirstName:  request.FirstName,
			SecondName: request.SecondName,
		}

		respData, err := client.CreateProfile(r.Context(), pbReq)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		profile, err := convert.ProfilePb(respData.Profile)
		if err != nil {
			response.Json().InternalError().Body(map[string]any{"error": err.Error()}).MustWrite(w)
			return
		}

		response.Json().Created().Body(profile).MustWrite(w)
	})
}
