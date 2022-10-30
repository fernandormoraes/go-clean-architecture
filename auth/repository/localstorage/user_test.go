package localstorage

import (
	"context"
	"github.com/fernandormoraes/go-clean-architecture/auth"
	"github.com/fernandormoraes/go-clean-architecture/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	s := NewUserLocalStorage()

	id1 := "id"

	user := &models.User{
		ID:       id1,
		Username: "user",
		Password: "password",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), "user", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = s.GetUser(context.Background(), "user", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
