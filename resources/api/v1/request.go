package v1

import (
	"log"
	"net/http"
	"os"
	"time"
)

// var channelResponse chan *http.Response

// NewRequestBuilder - initialize SummonerBuilder
func NewRequestBuilder(path string) *RequestBuilder {
	newRequest, errNewRequest := http.NewRequest("GET", os.Getenv("ENDPOINT_REGION")+path, nil)
	if errNewRequest != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed build request")
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
	if errResponseLeague != nil || responseLeague.StatusCode != 200 {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed perform request")
		return nil, errResponseLeague
	}
	return responseLeague, nil
}
