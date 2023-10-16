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

func TestUserStore_GetByPhone(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		user := &model.User{
			Name:      "name1",
			Birthdate: time.Now(),
			AvatarURL: "avatar1",
			Phone:     "12345678910",
			Email:     "son1@gmail.com",
			Balance:   10000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		require.NoError(t, TestDb.UserStore.Create(user))

		user, err := TestDb.UserStore.GetByPhone("12345678910")
		require.NoError(t, err)

		require.Equal(t, "avatar1", user.AvatarURL)
		require.Equal(t, "name1", user.Name)
		require.Equal(t, "12345678910", user.Phone)
		require.Equal(t, "son1@gmail.com", user.Email)
		require.Equal(t, 10000, user.Balance)
	})
}
