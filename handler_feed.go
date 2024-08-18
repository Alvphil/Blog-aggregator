package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Alvphil/Blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
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

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid,
		CreatedAt: currentTime.UTC(),
		UpdatedAt: currentTime.UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusNotAcceptable, "Couldn't add feed")
	} else {
		respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
	}

}
