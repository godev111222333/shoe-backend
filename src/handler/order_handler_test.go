package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/stretchr/testify/require"
)

func TestAPIServer_CreateOrder(t *testing.T) {
	t.Parallel()

	t.Run("create order success", func(t *testing.T) {
		routeInfo := TestAPIServer.AllRoutes()[RouteCreateOrder]
		body := map[string]interface{}{
			"user_id": 1,
			"products": []struct {
				ProductID int `json:"product_id"`
				Quantity  int `json:"quantity"`
				AtPrice   int `json:"at_price"`
			}{
				{
					ProductID: 13,
					Quantity:  5,
					AtPrice:   1_000_000,
				},
				{
					ProductID: 14,
					Quantity:  5,
					AtPrice:   1_000_000,
				},
				{
					ProductID: 15,
					Quantity:  5,
					AtPrice:   1_000_000,
				},
				{
					ProductID: 16,
					Quantity:  5,
					AtPrice:   1_000_000,
				},
			},
		}

		bz, err := json.Marshal(body)
		require.NoError(t, err)
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, bytes.NewReader(bz))

		recorder := httptest.NewRecorder()
		TestAPIServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAPIServer_GetAllOrders(t *testing.T) {
	t.Parallel()

	t.Run("Get all orders by user", func(t *testing.T) {
		t.Parallel()

		routeInfo := TestAPIServer.AllRoutes()[RouteGetAllOrders]
		req, _ := http.NewRequest(routeInfo.Method, routeInfo.Path, nil)
		query := req.URL.Query()
		query.Add("user_id", "1")
		query.Add("payment_status", fmt.Sprintf("%d", int(model.PaymentStatusPending)))
		req.URL.RawQuery = query.Encode()
		fmt.Println(req.URL.String())

		recorder := httptest.NewRecorder()
		TestAPIServer.route.ServeHTTP(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
}
