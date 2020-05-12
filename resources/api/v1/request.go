package v1

import (
	"net/http"
	"os"
)

const summonerV4 string = "/lol/summoner/v4/summoners/by-name/"
const leagueV4 string = "/lol/league/v4/entries/by-summoner/"

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

type RequestBuilder struct {
	pathParam, endpoint string
	keys, values        []string
}

func (req *RequestBuilder) SetPathParam(pathParam string) *RequestBuilder {
	req.pathParam = pathParam
	return req
}

func (req *RequestBuilder) SetQueries(keys []string, values []string) *RequestBuilder {
	req.keys = keys
	req.values = values
	return req
}

func (req *RequestBuilder) Build() (*http.Request, error) {
	newRequest, errNewRequest := http.NewRequest("GET", os.Getenv("ENDPOINT_REGION")+req.endpoint+req.pathParam, nil)
	if errNewRequest != nil {
		panic(errNewRequest)
		// return nil, errNewRequest
	}
	newRequest.Header.Set(os.Getenv("HEADER_API_KEY"), os.Getenv("API_KEY"))
	// apply queries keys and values in request
	// if len(req.keys) > 0
	return newRequest, nil
}

func (req *RequestBuilder) GetBuilder(requestType string) *RequestBuilder {
	switch requestType {
	case "summoner":
		req.endpoint = summonerV4
	case "league":
		req.endpoint = leagueV4
	default:
		req.endpoint = ""
	}
	return req
}
