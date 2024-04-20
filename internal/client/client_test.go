package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rornic/starlingexporter/internal/client"
	"github.com/rornic/starlingexporter/internal/model"
)

func TestStarlingHttpClientBalance(t *testing.T) {
	type test struct {
		response    string
		statusCode  int
		wantBalance *model.BalanceResponse
		wantErr     bool
	}
	tests := []test{
		{
			response:    `{"totalClearedBalance":{"currency": "GBP", "minorUnits": 100}}`,
			statusCode:  http.StatusOK,
			wantBalance: &model.BalanceResponse{TotalClearedBalance: model.CurrencyAndAmount{Currency: "GBP", MinorUnits: 100}},
			wantErr:     false,
		},
		{
			response:    `{"/{}£{£}89*("*($*(l)))", "minorUnits": 100}}`,
			statusCode:  http.StatusOK,
			wantBalance: nil,
			wantErr:     true,
		},
		{
			response:    `{"totalClearedBalance":{"currency": "GBP", "minorUnits": 100}}`,
			statusCode:  http.StatusBadRequest,
			wantBalance: nil,
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/accounts/my-account/balance" {
				t.Errorf("expected to request '/api/v2/accounts/my-account/balance', got: %s", r.URL.Path)
			}
			w.WriteHeader(tc.statusCode)
			w.Write([]byte(tc.response))
		}))
		defer server.Close()

		client := client.NewStarlingHttpClient("", server.URL)
		balance, err := client.Balance("my-account")

		if tc.wantErr && err == nil {
			t.Errorf("expected error, got none")
		}

		if !tc.wantErr && err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if err == nil && *balance != *tc.wantBalance {
			t.Errorf("expected %v, got %v", tc.wantBalance, balance)
		}
	}

}
