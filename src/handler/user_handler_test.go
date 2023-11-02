package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func TestAPIServer_UserLogin(t *testing.T) {
	t.Parallel()

	t.Run("login", func(t *testing.T) {
		routeInfo := TestAPIServer.AllRoutes()[RouteLoginUser]
		body := map[string]interface{}{
			"email": "touristversion2@gmail.com",
		}
		bodyBz, _ := json.Marshal(body)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bodyBz))
		recorder := httptest.NewRecorder()
		TestAPIServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAPIServer_VerifyLogin(t *testing.T) {
	t.Parallel()

	t.Run("verify login", func(t *testing.T) {
		routeInfo := TestAPIServer.AllRoutes()[RouteVerifyLoginUser]
		body := map[string]interface{}{
			"email": "touristversion2@gmail.com",
			"otp":   "727421",
		}
		bodyBz, _ := json.Marshal(body)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bodyBz))
		recorder := httptest.NewRecorder()
		TestAPIServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)

		respBody, err := io.ReadAll(recorder.Body)
		require.NoError(t, err)

		r := &VerifyLoginResponse{}
		require.NoError(t, json.Unmarshal(respBody, &r))

		require.Equal(t, "Son Le", r.Name)
		require.Equal(t, "touristversion2@gmail.com", r.Email)
		fmt.Println(r.AccessToken)
	})
}
