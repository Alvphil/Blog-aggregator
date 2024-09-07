package main

import (
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

func (cfg *apiConfig) GetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) GetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusNoContent, "Couldn't fetch posts")
	}
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
