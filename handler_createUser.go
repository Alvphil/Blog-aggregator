package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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

	ctx := context.Background()
	currentTime := time.Now()
	uuid := uuid.New()

	user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.String(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, user)
	}

}
