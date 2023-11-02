package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPIServer_RegisterUser(t *testing.T) {
	t.Parallel()

	t.Run("register success", func(t *testing.T) {
		routeInfo := TestAPIServer.AllRoutes()[RouteRegisterUser]
		body := map[string]interface{}{
			"phone": "123",
			"name":  "Son Le",
			"email": "touristversion2@gmail.com",
		}
		bodyBz, _ := json.Marshal(body)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bodyBz))
		recorder := httptest.NewRecorder()
		TestAPIServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAPIServer_VerifyRegistration(t *testing.T) {
	t.Parallel()

	t.Run("verify success", func(t *testing.T) {
		routeInfo := TestAPIServer.AllRoutes()[RouteVerifyRegisterUser]
		body := map[string]interface{}{
			"email": "touristversion2@gmail.com",
			"otp":   "864413",
		}
		bodyBz, _ := json.Marshal(body)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bodyBz))
		recorder := httptest.NewRecorder()
		TestAPIServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
}
