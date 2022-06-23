package di

import (
	user2 "github.com/8mamo10/gcl/internal/infra/user"
	"github.com/8mamo10/gcl/internal/presenter"
	"github.com/8mamo10/gcl/internal/repository/user"
	user3 "github.com/8mamo10/gcl/internal/usecase/user"
)

// Injectors from presenter.go:

func NewPresenter() *presenter.Presenter {
	useCase := NewUserUsecase()
	presenterPresenter := presenter.New(useCase)
	return presenterPresenter
}

// Injectors from repository.go:

func NewUserRepository() user.Repository {
	repository := user2.New()
	return repository
}

// Injectors from usecase.go:

func NewUserUsecase() user3.UseCase {
	repository := NewUserRepository()
	useCase := user3.New(repository)
	return useCase
}
