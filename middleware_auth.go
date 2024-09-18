package main

import (
	"net/http"
	"fmt"

	"github.com/lemmrz/rssagg/internal/auth"
	"github.com/lemmrz/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
