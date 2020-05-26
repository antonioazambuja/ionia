package v1

import (
	"log"
	"net/http"
	"os"
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
}

// WithPathParam - add path parameter in request
func (req *RequestBuilder) WithPathParam(pathParam string) *RequestBuilder {
	req.pathParam = pathParam
	return req
}

// WithQueries - add queries in request
func (req *RequestBuilder) WithQueries(keys []string, values []string) *RequestBuilder {
	req.keys = keys
	req.values = values
	return req
}

// Build - create request with API_KEY and HEADER
func (req *RequestBuilder) Build() (*http.Request, error) {
	newRequest, errNewRequest := http.NewRequest("GET", os.Getenv("ENDPOINT_REGION")+req.endpoint+req.pathParam, nil)
	if errNewRequest != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print(errNewRequest)
		return nil, errNewRequest
	}
	newRequest.Header.Set(os.Getenv("HEADER_API_KEY"), os.Getenv("API_KEY"))
	if len(req.keys) > 0 && len(req.values) > 0 {
		query := newRequest.URL.Query()
		if len(req.keys) == len(req.values) {
			for i := range req.keys {
				query.Add(req.keys[i], req.values[i])
			}
		}
	}
	return newRequest, nil
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
