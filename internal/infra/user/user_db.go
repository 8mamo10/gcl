package user

import (
	"context"

	"github.com/8mamo10/gcl/internal/domain"
	"github.com/8mamo10/gcl/internal/infra"
	userRepo "github.com/8mamo10/gcl/internal/repository/user"
)

type user struct {
	ID        int64
	AccountID int64
	Name      string
	Email     string
}

func New() userRepo.Repository {
	return &repository{
		converter: &converter{},
	}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	db := infra.GetDB()
	rows, err := db.Model(&user{}).Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for rows.Next() {
		var user user
		err := db.ScanRows(rows, &user)
		if err != nil {
			return nil, err
		}

		users = append(users, r.converter.To(&user))
	}

	return users, nil
}

type repository struct {
	converter domainConverter
}

type domainConverter interface {
	To(o *user) *domain.User
}

type converter struct{}

func (c *converter) To(o *user) *domain.User {
	return &domain.User{
		ID:        int64(o.ID),
		AccountID: o.AccountID,
		Name:      o.Name,
		Email:     o.Email,
	}
}
