package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/profile/api/pb"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
)

func GetProfileByUsernameHandler(client pb.ProfileServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]

		if username == "" {
			response.WriteMissingFieldsError(w, []string{"username"})
			return
		}

		request := &pb.GetProfileRequest{
			SearchBy: &pb.GetProfileRequest_Username{
				Username: username,
			},
		}

		profile, err := client.GetProfile(r.Context(), request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().OK().Body(profile).MustWrite(w)
	})
}

func GetProfileByIdHandler(client pb.ProfileServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.URL.Query().Get("id")
		if idRaw == "" {
			response.Json().BadRequest().Body(map[string]any{"error": "missing id in query params"}).MustWrite(w)
			return
		}

		profileID, err := strconv.ParseInt(idRaw, 0, 10)
		if err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		request := &pb.GetProfileRequest{
			SearchBy: &pb.GetProfileRequest_Id{
				Id: profileID,
			},
		}

		profile, err := client.GetProfile(r.Context(), request)
		if err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().OK().Body(profile).MustWrite(w)
	})
}
