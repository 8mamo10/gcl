package user

import (
	"context"

	"github.com/8mamo10/gcl/internal/domain"
)

type repository interface {
	GetAll(ctx context.Context, limit, offset int) ([]*domain.User, error)
}
