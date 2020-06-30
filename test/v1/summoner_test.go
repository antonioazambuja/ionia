package v1

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/stretchr/testify/assert"
	// assert "github.com/stretchr/testify/assert"
)

func TestServiceGetInfoByName(test *testing.T) {
	test.Parallel()
	bodyServiceInfoByName := []byte(`{
		"name": "IsBlackPanther",
		"summonerLevel": 1,
		"id": "OSIlasjdsanc",
		"accountId": "IsolsmdhYDHadKBDA-9fM",
		"puuid": "jClaj-2S4ZsbjRgIItf1PtjL7-FXbqyDeC",
		"profileIconId": 1,
		"revisionDate": 1592855746000
	}`)
	responseSummonerMock := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(bodyServiceInfoByName)),
		ContentLength: int64(len(bodyServiceInfoByName)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
	defer responseSummonerMock.Body.Close()
	summoner := rsc_v1.NewSummoner(responseSummonerMock)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.Equal(test, 1, summoner.SummonerLevel)
	assert.Equal(test, "OSIlasjdsanc", summoner.SummonerID)
	assert.Equal(test, "IsolsmdhYDHadKBDA-9fM", summoner.AccountID)
	assert.Equal(test, "jClaj-2S4ZsbjRgIItf1PtjL7-FXbqyDeC", summoner.Puuid)
	assert.Equal(test, 1, summoner.ProfileIconID)
	assert.Equal(test, 1592855746000, summoner.RevisionDate)
}

// func TestGetSummonerInfoByName(test *testing.T) {
// 	test.Parallel()
// 	summoner, errSummoner := svc_v1.GetInfoByName("IsBlackPanther")
// 	assert.Empty(test, errSummoner)
// 	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
// 	assert.NotEmpty(test, summoner.AccountID)
// 	assert.Empty(test, summoner.LeagueInfo)
// 	assert.Empty(test, summoner.MatchesInfo)
// }

// func TestGetSummonerLeagueByName(test *testing.T) {
// 	test.Parallel()
// 	summoner, errSummoner := svc_v1.GetLeagueByName("IsBlackPanther")
// 	assert.Empty(test, errSummoner)
// 	assert.Empty(test, summoner.AccountID)
// 	assert.NotEmpty(test, summoner.LeagueInfo)
// 	assert.Empty(test, summoner.MatchesInfo)
// }

// func TestGetSummonerMatchesByName(test *testing.T) {
// 	test.Parallel()
// 	summoner, errSummoner := svc_v1.GetMatchesByName("IsBlackPanther")
// 	assert.Empty(test, errSummoner)
// 	assert.Empty(test, summoner.AccountID)
// 	assert.Empty(test, summoner.LeagueInfo)
// 	assert.NotEmpty(test, summoner.MatchesInfo)
// }

// func TestInvalidSummonerName(test *testing.T) {
// 	test.Parallel()
// 	summoner1, errSummoner1 := svc_v1.GetByName("&*|ª")
// 	assert.NotEmpty(test, errSummoner1)
// 	assert.Empty(test, summoner1)
// 	summoner2, errSummoner2 := svc_v1.GetInfoByName("&*|ª")
// 	assert.NotEmpty(test, errSummoner2)
// 	assert.Empty(test, summoner2)
// 	summoner3, errSummoner3 := svc_v1.GetLeagueByName("&*|ª")
// 	assert.NotEmpty(test, errSummoner3)
// 	assert.Empty(test, summoner3)
// 	summoner4, errSummoner4 := svc_v1.GetMatchesByName("&*|ª")
// 	assert.NotEmpty(test, errSummoner4)
// 	assert.Empty(test, summoner4)
// }
