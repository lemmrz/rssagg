package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lemmrz/rssagg/internal/auth"
	"github.com/lemmrz/rssagg/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:name`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt create user: %s", err))
		return
	}
	respondWithJSON(w, 201, databseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	key, err := auth.GetApiKeyFromHeader(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Couldn't extract api key: %s", err))
		return
	}
	
	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), key)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Error getting user from db: %s", err))
		return
	}

	respondWithJSON(w, 201, databseUserToUser(user))
}
