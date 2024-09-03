package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Alvphil/Blog-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPostFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_id uuid.UUID `json:"feed_id"`
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

	feed, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid,
		CreatedAt: currentTime.UTC(),
		UpdatedAt: currentTime.UTC(),
		UserID:    user.ID,
		FeedID:    params.Feed_id,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't add feed")
	} else {
		respondWithJSON(w, http.StatusCreated, databaseFollowFeed(feed))
	}
}

func (cfg *apiConfig) handlerDeleteFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	string_Feed_id := chi.URLParam(r, "feedFollowId")
	Feed_id, err := uuid.Parse(string_Feed_id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feedID: %v", Feed_id.String()))
		return
	}

	err = cfg.DB.DeleteFollowFeed(r.Context(), database.DeleteFollowFeedParams{
		ID:     Feed_id,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not delete follow")
	} else {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}

}

func (cfg *apiConfig) handlerGetAllFollowFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsUnsanitized, err := cfg.DB.GetAllFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Found no feeds")
	}
	var followFeeds = []FollowFeed{}
	for _, f := range feedsUnsanitized {
		followFeeds = append(followFeeds, databaseFollowFeed(f))
	}
	respondWithJSON(w, http.StatusOK, followFeeds)
}
