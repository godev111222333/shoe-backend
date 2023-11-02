package handler

import (
	"testing"

	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/stretchr/testify/require"
)

func TestOTPService_SendOTP(t *testing.T) {
	t.Parallel()

	t.Run("send OTP", func(t *testing.T) {
		otpService := NewOTPService(TestDb, "shoestoresevice111@gmail.com", "lvlc hnck fpxl ibwi")
		require.NoError(t, otpService.SendOTP(model.OTPTypeRegister, "touristversion2@gmail.com"))
	})
}
