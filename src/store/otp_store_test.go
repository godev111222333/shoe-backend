package store

import (
	"testing"
	"time"

	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/stretchr/testify/require"
)

func TestOTPStore_GetLastByOTPType(t *testing.T) {
	t.Parallel()

	t.Run("Create and get last", func(t *testing.T) {
		t.Parallel()

		otp := &model.OTP{
			Type:      model.OTPTypeRegister,
			Email:     "touristversion2@gmail.com",
			Code:      "123456",
			Status:    model.OTPStatusSent,
			ExpiresAt: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		require.NoError(t, TestDb.OTPStore.Create(otp))

		found, err := TestDb.OTPStore.GetLastByOTPType("touristversion2@gmail.com", model.OTPTypeRegister)
		require.NoError(t, err)
		require.NotEmpty(t, found)
		require.Equal(t, "123456", found.Code)
		require.Equal(t, model.OTPStatusSent, found.Status)

		require.NoError(t, TestDb.OTPStore.UpdateStatus("touristversion2@gmail.com", model.OTPTypeRegister, model.OTPStatusVerified))
		found1, err := TestDb.OTPStore.GetLastByOTPType("touristversion2@gmail.com", model.OTPTypeRegister)
		require.NoError(t, err)
		require.NotEmpty(t, found1)
		require.Equal(t, model.OTPStatusVerified, found1.Status)
	})
}
