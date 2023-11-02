package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/stretchr/testify/require"
)

func TestAPIServer_RegisterUser(t *testing.T) {
	t.Parallel()

	apiServer := NewAPIServer(&misc.APIConfig{Host: "0.0.0.0", Port: "9090"}, TestDb, nil)

	t.Run("register success", func(t *testing.T) {
		routeInfo := apiServer.AllRoutes()[RouteRegisterUser]
		body := map[string]interface{}{
			"phone": "123",
			"name":  "Son Le",
			"email": "son1@gmail.com",
		}
		bodyBz, _ := json.Marshal(body)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bodyBz))
		recorder := httptest.NewRecorder()
		apiServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("login", func(t *testing.T) {
		routeInfo := apiServer.AllRoutes()[RouteLogin]
		body := map[string]interface{}{
			"phone": "123",
			"otp":   "1111",
		}
		bodyBz, _ := json.Marshal(body)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bodyBz))
		recorder := httptest.NewRecorder()
		apiServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)

		bz, err := io.ReadAll(recorder.Body)
		require.NoError(t, err)

		user := &model.User{}
		require.NoError(t, json.Unmarshal(bz, user))

		require.Equal(t, "123", user.Phone)
		require.Equal(t, "Son Le", user.Name)
		require.Equal(t, "son1@gmail.com", user.Email)
	})
}
