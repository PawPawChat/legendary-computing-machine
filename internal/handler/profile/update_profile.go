package profile

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pawpawchat/core/internal/convert"
	"github.com/pawpawchat/core/internal/model"
	"github.com/pawpawchat/core/pkg/response"
	profilepb "github.com/pawpawchat/profile/api/pb"
)

func UpdateProfileHandler(profileClient profilepb.ProfileServiceClient) http.Handler {
	type UpdateProfileRequest struct {
		Username    string           `json:"username"`
		Description string           `json:"deescription"`
		Biography   *model.Biography `json:"biography"`
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

		var request UpdateProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		// filling of biographical data if obtained
		bioPb, err := convert.Biography(request.Biography)
		if err != nil {
			response.WriteParseBodyError(w, err)
			return
		}

		if bioPb == nil && request.Description == "" && request.Username == "" {
			response.Json().BadRequest().Body(map[string]any{"error": "has no fields to update"}).MustWrite(w)
			return
		}

		requestPb := &profilepb.UpdateProfileRequest{
			ProfileId:   id,
			Username:    request.Username,
			Description: request.Description,
			Biography:   bioPb,
		}

		// rpc
		if _, err = profileClient.UpdateProfile(r.Context(), requestPb); err != nil {
			response.WriteProtoError(w, err)
			return
		}

		response.Json().OK().MustWrite(w)
	})
}
