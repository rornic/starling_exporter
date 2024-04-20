package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rornic/starlingexporter/internal/model"
)

type StarlingClient interface {
	Accounts() (*model.AccountsResponse, error)
	Balance(account string) (*model.BalanceResponse, error)
	FeedItems(account string, category string, since time.Time) (*model.FeedItemsResponse, error)
}

type StarlingHttpClient struct {
	endpoint    string
	httpClient  http.Client
	accessToken string
}

func NewStarlingHttpClient(accessToken string, endpoint string) StarlingHttpClient {
	return StarlingHttpClient{
		endpoint:    endpoint,
		httpClient:  http.Client{Timeout: 10 * time.Second},
		accessToken: accessToken,
	}
}

func (starlingClient *StarlingHttpClient) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", starlingClient.endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	resp, err := starlingClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad http response code: %d", resp.StatusCode)
	}

	return resp, nil
}

func (starlingClient *StarlingHttpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", starlingClient.accessToken))
	return starlingClient.httpClient.Do(req)
}

func (starlingHttpClient *StarlingHttpClient) Accounts() (*model.AccountsResponse, error) {
	resp, err := starlingHttpClient.Get("/accounts")
	if err != nil {
		return nil, err
	}

	accountsResponse := model.AccountsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&accountsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &accountsResponse, nil
}

func (starlingHttpClient *StarlingHttpClient) Balance(account string) (*model.BalanceResponse, error) {
	resp, err := starlingHttpClient.Get(fmt.Sprintf("/accounts/%s/balance", account))
	if err != nil {
		return nil, err
	}

	balanceResponse := model.BalanceResponse{}
	err = json.NewDecoder(resp.Body).Decode(&balanceResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &balanceResponse, nil
}

func (starlingHttpClient *StarlingHttpClient) FeedItems(account string, category string, since time.Time) (*model.FeedItemsResponse, error) {
	resp, err := starlingHttpClient.Get(fmt.Sprintf("/feed/account/%s/category/%s?changesSince=%s", account, category, since.Format("2006-01-02T15:04:05.999Z")))
	if err != nil {
		return nil, err
	}

	feedItemsResponse := model.FeedItemsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&feedItemsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &feedItemsResponse, nil
}
