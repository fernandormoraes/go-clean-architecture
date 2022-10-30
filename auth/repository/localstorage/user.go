package localstorage

import (
	"context"
	"sync"

	"github.com/fernandormoraes/go-clean-architecture/auth"
	models "github.com/fernandormoraes/go-clean-architecture/domain/entities"
)

type UserLocalStorage struct {
	users map[string]*models.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*models.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *models.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
