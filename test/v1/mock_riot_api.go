package v1

import (
	"net/http"
	"time"

	svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
)

// RestClientRiotAPI - Client interface
type RestClientRiotAPI interface {
	GetInfoByName(summonerName string) (*http.Response, error)
}

// RiotAPI - Client interface
type RiotAPI struct {
	clientAPIRiot http.Client
	baseURL       string
}

// NewMockRiotAPI - New client with defaults
func NewMockRiotAPI(url string) *RiotAPI {
	return &RiotAPI{
		clientAPIRiot: http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		baseURL: url,
	}
}

// GetInfoByName returns a user
func (mockRestRiotAPI *RiotAPI) GetInfoByName(id string) (*http.Response, error) {
	req, err := http.NewRequest("GET", mockRestRiotAPI.baseURL+svc_v1.SummonerV4+id, nil)

	resp, err := mockRestRiotAPI.clientAPIRiot.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
