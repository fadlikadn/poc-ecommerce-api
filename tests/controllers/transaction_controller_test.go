package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAddTransactionDetails(t *testing.T) {
	err := refreshTable()
	if err != nil {
		log.Fatal(err)
	}

	_, _, _, _, _, _, err = SeedDataTest()
	if err != nil {
		log.Fatal(err)
	}

	samples := []struct {
		id           string
		inputJSON    string
		price        int
		quantity     int
		statusCode   int
		errorMessage string
	}{
		{
			id:           strconv.Itoa(1),
			inputJSON:    `{"transaction_id": 1, "warehouse_item_id": 4, "price": 500000, "quantity": 1 }`,
			statusCode:   422,
			errorMessage: "Item out of stock",
		},
		{
			id:         strconv.Itoa(1),
			inputJSON:  `{"transaction_id": 1, "warehouse_item_id": 2, "price": 500000, "quantity": 1 }`,
			statusCode: 201,
			quantity:   1,
			price:      500000,
		},
	}

	for _, v := range samples {
		req, err := http.NewRequest("POST", "/transaction-details", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.AddTransactionDetail)

		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 422 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
