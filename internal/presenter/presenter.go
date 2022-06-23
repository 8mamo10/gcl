package presenter

import "github.com/8mamo10/gcl/internal/usecase/user"

type Presenter struct {
	userUC user.UseCase
}

func New(userUC user.UseCase) *Presenter {
	return &Presenter{
		userUC: userUC,
	}
}

const (
	defaultLimit  = 100
	defaultOffset = 0
)
