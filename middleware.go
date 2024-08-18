package main

import (
	"fmt"
	"net/http"

	"github.com/Alvphil/Blog-aggregator/internal/auth"
)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("auth error: %v", err))
			return
		}
		handler(w, r, user)
	})
}
