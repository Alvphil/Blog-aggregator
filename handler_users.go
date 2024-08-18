package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Alvphil/Blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	currentTime := time.Now()
	uuid := uuid.New()

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid,
		CreatedAt: currentTime.UTC(),
		UpdatedAt: currentTime.UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, "Couldn't create user")
	} else {
		respondWithJSON(w, http.StatusCreated, user)
	}

}

func (cfg *apiConfig) GetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithError(w, http.StatusBadRequest, "Missing auth header")
		return
	}
	const keyText = "ApiKey "
	if !strings.HasPrefix(authHeader, keyText) {
		respondWithError(w, http.StatusBadRequest, "Invalid Authorization header format")
		return
	}
	ApiKey := strings.TrimPrefix(authHeader, keyText)

	user, err := cfg.DB.GetUserByApiKey(r.Context(), ApiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't locate user")
	} else {
		respondWithJSON(w, http.StatusCreated, user)
	}
}
