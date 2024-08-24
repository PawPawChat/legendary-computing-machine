package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/pawpawchat/core/pkg/json"
	"github.com/pawpawchat/core/pkg/response"
	"github.com/pawpawchat/profile/api/pb"
)

func GetProfileByUsernameHandler(client pb.ProfileServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := strings.Trim(r.URL.Path, "/")

		profile, err := client.GetByUsername(r.Context(), &pb.GetProfileByUsernameRequest{Username: username})
		if err != nil {
			// REFACTOR: PARSE ERROR
			slog.Error(fmt.Sprintf("500 internal error: %s", err.Error()))
			response.Json(w).Code(http.StatusInternalServerError).Args("error", err.Error()).MustSend()
			return
		}

		payload := json.MustMarshal(profile)
		response.Json(w).Code(http.StatusBadRequest).Body(payload).MustSend()
	})
}

func CreateProfileHandler(client pb.ProfileServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pbReq := new(pb.CreateProfileRequest)

		if err := json.Decode(r.Body, pbReq); err != nil {
			slog.Error("400 failed to parse body")
			response.Json(w).Code(http.StatusBadRequest).Args("error", err.Error()).MustSend()
			return
		}

		if pbReq.FirstName == "" || pbReq.SecondName == "" {
			slog.Error(fmt.Sprintf("400 missing required data, req=%+v", pbReq))
			response.Json(w).Code(http.StatusBadRequest).Args("error", "mussing required fields").MustSend()
			return
		}

		profile, err := client.Create(r.Context(), pbReq)
		if err != nil {
			slog.Error(fmt.Sprintf("500 internal error: %s", err.Error()))
			response.Json(w).Code(http.StatusInternalServerError).Args("error", err.Error()).MustSend()
			return
		}

		payload := json.MustMarshal(profile)
		response.Json(w).Code(http.StatusBadRequest).Body(payload).MustSend()
	})
}
