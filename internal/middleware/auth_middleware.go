package middleware

import (
	"context"
	"github.com/gnomedevreact/CombatIntel/internal/auth"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		userId, err := auth.ValidateJWT(token)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
