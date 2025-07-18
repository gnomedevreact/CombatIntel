package middleware

import (
	"context"
	"errors"
	"github.com/gnomedevreact/CombatIntel/internal/api"
	"github.com/gnomedevreact/CombatIntel/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"slices"
)

type Role string

const (
	Admin   Role = "admin"
	Officer Role = "officer"
)

func isValidRole(roles []Role) bool {
	for _, role := range roles {
		if role != Admin && role != Officer {
			return false
		}
	}
	return true
}

func RolesMiddleware(next http.Handler, allowedRoles []Role, cfg *api.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId")

		userUuid, err := uuid.Parse(userId.(string))
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		pgUuid := pgtype.UUID{
			Bytes: userUuid,
			Valid: true,
		}
		user, err := cfg.Db.GetUserById(context.Background(), pgUuid)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
		}

		if !isValidRole(allowedRoles) || !slices.Contains(allowedRoles, Role(user.Role)) {
			utils.RespondWithError(w, http.StatusForbidden, errors.New("Forbidden"))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "userId", userId)))
	})
}
