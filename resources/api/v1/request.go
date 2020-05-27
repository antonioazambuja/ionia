package v1

import (
	"log"
	"net/http"
	"os"
	"time"
)

const summonerV4 string = "/lol/summoner/v4/summoners/by-name/"
const leagueV4 string = "/lol/league/v4/entries/by-summoner/"
const matchesV4 string = "/lol/match/v4/matchlists/by-account/"

// NewRequestBuilder - initialize SummonerBuilder
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

// RequestBuilder - builder request
type RequestBuilder struct {
	pathParam, endpoint string
	keys, values        []string
	requestBuilded      *http.Request
}

// WithPathParam - add path parameter in request
func (req *RequestBuilder) WithPathParam(pathParam string) *RequestBuilder {
	req.pathParam = pathParam
	return req
}

// WithQueries - add queries in request
func (req *RequestBuilder) WithQueries(keys []string, values []string) *RequestBuilder {
	for _, item := range keys {
		req.keys = append(req.keys, item)
	}
	for _, item := range values {
		req.values = append(req.values, item)
	}
	return req
}

// Build - create request with API_KEY and HEADER
func (req *RequestBuilder) Build() *RequestBuilder {
	newRequest, errNewRequest := http.NewRequest("GET", os.Getenv("ENDPOINT_REGION")+req.endpoint+req.pathParam, nil)
	if errNewRequest != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed build request")
		return nil
	}
	newRequest.Header.Set(os.Getenv("HEADER_API_KEY"), os.Getenv("API_KEY"))
	if len(req.keys) > 0 && len(req.values) > 0 {
		query := newRequest.URL.Query()
		if len(req.keys) == len(req.values) {
			for i := range req.keys {
				query.Add(req.keys[i], req.values[i])
			}
		}
		newRequest.URL.RawQuery = query.Encode()
	}
	req.requestBuilded = newRequest
	return req
}

// Run - perform request builded
func (req *RequestBuilder) Run() (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	responseLeague, errResponseLeague := client.Do(req.requestBuilded)
	if errResponseLeague != nil || responseLeague.StatusCode != 200 {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed perform request")
		return nil, errResponseLeague
	}
	return responseLeague, nil
}

// TypeBuilder - set type of builder
func (req *RequestBuilder) TypeBuilder(requestType string) *RequestBuilder {
	switch requestType {
	case "summoner":
		req.endpoint = summonerV4
	case "league":
		req.endpoint = leagueV4
	case "matches":
		req.endpoint = matchesV4
	default:
		req.endpoint = ""
	}
	return req
}
