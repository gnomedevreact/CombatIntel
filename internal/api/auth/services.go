package auth

import (
	"context"
	"github.com/gnomedevreact/CombatIntel/internal/api"
	"github.com/gnomedevreact/CombatIntel/internal/auth"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	cfg *api.ApiConfig
}

func (s *authService) register(ctx context.Context, reqData regsiterRequest) (*authResponse, error) {
	hashesPassword, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), 15)
	if err != nil {
		return nil, err
	}

	var userResp models.User
	user, err := s.cfg.Db.CreateUser(ctx, database.CreateUserParams{
		Username:       reqData.Username,
		Password:       string(hashesPassword),
		ClearanceLevel: "C",
	})
	if err != nil {
		return nil, err
	}
	copier.Copy(&userResp, &user)

	token, err := auth.GenerateJWT(user.ID.String(), s.cfg.Secret)
	if err != nil {
		return nil, err
	}

	return &authResponse{
		User:        userResp,
		AccessToken: token,
	}, nil
}

func (s *authService) login(ctx context.Context, reqData loginRequest) (*authResponse, error) {
	var userResp models.User
	user, err := s.cfg.Db.GetUserByUsername(ctx, reqData.Username)
	if err != nil {
		return nil, err
	}
	copier.Copy(&userResp, &user)

	hashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(reqData.Password))
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateJWT(user.ID.String(), s.cfg.Secret)
	if err != nil {
		return nil, err
	}

	return &authResponse{
		User:        userResp,
		AccessToken: token,
	}, nil
}
