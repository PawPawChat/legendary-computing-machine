package profile

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pawpawchat/core/internal/convert"
	"github.com/pawpawchat/core/internal/model"
	"github.com/pawpawchat/core/pkg/response"
	profilepb "github.com/pawpawchat/profile/api/pb"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
)

func GetProfileByUsernameHandler(client profilepb.ProfileServiceClient) http.Handler {
	type GetProfileResponse struct {
		Profile *model.Profile  `json:"profile"`
		Avatars []*model.Avatar `json:"avatars"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]

		if username == "" {
			response.WriteMissingFieldsError(w, []string{"username"})
			return
		}

		request := &profilepb.GetProfileRequest{
			SearchBy: &profilepb.GetProfileRequest_Username{
				Username: username,
			},
		}

		respPb, err := client.GetProfile(r.Context(), request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().
			OK().
			Body(&GetProfileResponse{
				Profile: convert.MustProfilePb(respPb.Profile),
				Avatars: convert.MustFromPb(respPb.Avatars, convert.MustAvatarPb),
			}).
			MustWrite(w)
	})
}

func GetProfileByIdHandler(client profilepb.ProfileServiceClient) http.Handler {
	type GetProfileResponse struct {
		Profile *model.Profile  `json:"profile"`
		Avatars []*model.Avatar `json:"avatars"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 0, 10)
		if err != nil {
			response.Json().
				BadRequest().
				Body(map[string]any{
					"error": map[string]any{
						"message": "cannot parse profile id",
						"value":   idStr,
					}}).
				MustWrite(w)
			return
		}

		requestPb := &profilepb.GetProfileRequest{
			SearchBy: &profilepb.GetProfileRequest_Id{
				Id: id,
			},
		}

		respPb, err := client.GetProfile(r.Context(), requestPb)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().
			OK().
			Body(&GetProfileResponse{
				Profile: convert.MustProfilePb(respPb.Profile),
				Avatars: convert.MustFromPb(respPb.Avatars, convert.MustAvatarPb),
			}).
			MustWrite(w)
	})
}
