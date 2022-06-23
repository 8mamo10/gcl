package user

import (
	"context"

	"github.com/8mamo10/gcl/internal/domain"
	"github.com/8mamo10/gcl/internal/repository/user"
)

type UseCase interface {
	GetAll(ctx context.Context, limit, offset int) ([]*domain.User, error)
}

func New(userRepo user.Repository) UseCase {
	return &useCase{userRepo: userRepo}
}

type useCase struct {
	userRepo user.Repository
}

func (uc *useCase) GetAll(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	return uc.userRepo.GetAll(ctx, limit, offset)
}
