package store

import (
	"testing"
	"time"

	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/stretchr/testify/require"
)

func TestUserStore_Create(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		user := &model.User{
			Username:  "username",
			Password:  "password",
			Name:      "name",
			Birthdate: time.Now(),
			AvatarURL: "avatar",
			Phone:     "0123456789",
			Email:     "son@gmail.com",
			Balance:   10000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		require.NoError(t, TestDb.UserStore.Create(user))
	})
}
