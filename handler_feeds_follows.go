package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lemmrz/rssagg/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedsFollow(r.Context(), database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt create feeds follows: %s", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feeds follows: %s", err))
		return
	}
	respondWithJSON(w, 200, feeds)
}

func (apiCfg *apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt parse feeds follows id: %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't delete feeds follows: %s", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
