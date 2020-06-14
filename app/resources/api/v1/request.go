package v1

import (
	"net/http"
	"os"
	"time"

	utils "github.com/antonioazambuja/ionia/utils"
)

// NewRequestBuilder - initialize SummonerBuilder
func NewRequestBuilder(path string) *RequestBuilder {
	newRequest, errNewRequest := http.NewRequest("GET", os.Getenv("ENDPOINT_REGION")+path, nil)
	if errNewRequest != nil {
		utils.LogOperation.Print("Failed build request")
		return nil
	}
	newRequest.Header.Set(os.Getenv("HEADER_API_KEY"), os.Getenv("API_KEY"))
	return &RequestBuilder{
		requestBuilded: newRequest,
	}
}

// RequestBuilder - builder request
type RequestBuilder struct {
	pathParam, endpoint string
	keys, values        []string
	requestBuilded      *http.Request
}

// WithPathParam - add path parameter in request
func (req *RequestBuilder) WithPathParam(pathParam string) *RequestBuilder {
	req.requestBuilded.URL.Path = req.requestBuilded.URL.Path + pathParam
	return req
}

// WithQueries - add queries in request
func (req *RequestBuilder) WithQueries(keys []string, values []string) *RequestBuilder {
	query := req.requestBuilded.URL.Query()
	if len(req.keys) == len(req.values) {
		for i := range req.keys {
			query.Add(req.keys[i], req.values[i])
		}
	}
	req.requestBuilded.URL.RawQuery = query.Encode()
	return req
}

// Run - perform request builded
func (req *RequestBuilder) Run() (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	responseLeague, errResponseLeague := client.Do(req.requestBuilded)
	if errResponseLeague != nil {
		utils.LogOperation.Print("Failed perform request: " + req.endpoint + " - " + errResponseLeague.Error())
		return nil, errResponseLeague
	} else if responseLeague.StatusCode != 200 {
		utils.LogOperation.Print("Failed get request with following info: '" + req.pathParam + "' in: '" + req.endpoint + "' - Invalid status code: " + responseLeague.Status)
		return nil, errResponseLeague
	}
	return responseLeague, nil
}
