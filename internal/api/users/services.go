package users

import (
	"context"
	"github.com/gnomedevreact/CombatIntel/internal/database"
)

type usersService struct {
	db *database.Queries
}

func (s *usersService) GetAllUsers(ctx context.Context) (*[]database.User, error) {
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return &users, nil
}
