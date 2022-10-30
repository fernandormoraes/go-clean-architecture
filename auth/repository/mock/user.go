package mock

import (
	"context"

	models "github.com/fernandormoraes/go-clean-architecture/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(ctx context.Context, user *models.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	args := s.Called(username, password)

	return args.Get(0).(*models.User), args.Error(1)
}
