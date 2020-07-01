package v1

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/antonioazambuja/ionia/utils"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetSummonerByName(test *testing.T) {
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
	summoner := rsc_v1.NewSummoner(responseSummonerMock)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.Equal(test, 1, summoner.SummonerLevel)
	assert.Equal(test, "OSIlasjdsanc", summoner.SummonerID)
	assert.Equal(test, "IsolsmdhYDHadKBDA-9fM", summoner.AccountID)
	assert.Equal(test, "jClaj-2S4ZsbjRgIItf1PtjL7-FXbqyDeC", summoner.Puuid)
	assert.Equal(test, 1, summoner.ProfileIconID)
	assert.Equal(test, 1592855746000, summoner.RevisionDate)
	bodyServiceLeagueByName := []byte(`
		[
			{
				"leagueId": "b80abef5-4fc6-ba48",
				"queueType": "RANKED_SOLO_5x5",
				"tier": "PLATINUM",
				"rank": "II",
				"summonerId": "OSIlasjdsanc",
				"summonerName": "IsBlackPanther",
				"leaguePoints": 25,
				"wins": 17,
				"losses": 23,
				"veteran": false,
				"inactive": false,
				"freshBlood": false,
				"hotStreak": false
			}
		]
	`)
	responseSummonerLeagueMock := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(bodyServiceLeagueByName)),
		ContentLength: int64(len(bodyServiceLeagueByName)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
	summoner.WithLeagueInfo(responseSummonerLeagueMock)
	assert.Equal(test, "b80abef5-4fc6-ba48", summoner.LeagueInfo[0].LeagueID)
	assert.Equal(test, "RANKED_SOLO_5x5", summoner.LeagueInfo[0].QueueType)
	assert.Equal(test, "PLATINUM", summoner.LeagueInfo[0].Tier)
	assert.Equal(test, "II", summoner.LeagueInfo[0].Rank)
	assert.Equal(test, 25, summoner.LeagueInfo[0].LeaguePoints)
	assert.Equal(test, 17, summoner.LeagueInfo[0].Wins)
	assert.Equal(test, 23, summoner.LeagueInfo[0].Losses)
	assert.Equal(test, false, summoner.LeagueInfo[0].Veteran)
	assert.Equal(test, false, summoner.LeagueInfo[0].Inactive)
	assert.Equal(test, false, summoner.LeagueInfo[0].FreshBlood)
	assert.Equal(test, false, summoner.LeagueInfo[0].HotStreak)
	assert.Equal(test, 0, summoner.LeagueInfo[0].MiniSeriesDTO.Losses)
	assert.Equal(test, "", summoner.LeagueInfo[0].MiniSeriesDTO.Progress)
	assert.Equal(test, 0, summoner.LeagueInfo[0].MiniSeriesDTO.Target)
	assert.Equal(test, 0, summoner.LeagueInfo[0].MiniSeriesDTO.Wins)
	bodyServiceMatchesByName := []byte(`
		{
			"matches": [
				{
					"platformId": "BR1",
					"gameId": 1982082824,
					"champion": 39,
					"queue": 420,
					"season": 13,
					"timestamp": 1592878663522,
					"role": "SOLO",
					"lane": "TOP"
				},
				{
					"platformId": "BR1",
					"gameId": 1981939620,
					"champion": 39,
					"queue": 420,
					"season": 13,
					"timestamp": 1592876962973,
					"role": "SOLO",
					"lane": "TOP"
				}
			],
			"startIndex": 0,
			"endIndex": 50,
			"totalGames": 158
		}
	`)
	responseSummonerMatchesMock := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(bodyServiceMatchesByName)),
		ContentLength: int64(len(bodyServiceMatchesByName)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
	summoner.WithMatchesInfo(responseSummonerMatchesMock)
	assert.Equal(test, "BR1", summoner.MatchesInfo[0].PlatformID)
	assert.Equal(test, 1982082824, summoner.MatchesInfo[0].GameID)
	assert.Equal(test, 39, summoner.MatchesInfo[0].Champion)
	assert.Equal(test, 420, summoner.MatchesInfo[0].Queue)
	assert.Equal(test, 13, summoner.MatchesInfo[0].Season)
	assert.Equal(test, 1592878663522, summoner.MatchesInfo[0].Timestamp)
	assert.Equal(test, "SOLO", summoner.MatchesInfo[0].Role)
	assert.Equal(test, "TOP", summoner.MatchesInfo[0].Lane)
	assert.Equal(test, "BR1", summoner.MatchesInfo[1].PlatformID)
	assert.Equal(test, 1981939620, summoner.MatchesInfo[1].GameID)
	assert.Equal(test, 39, summoner.MatchesInfo[1].Champion)
	assert.Equal(test, 420, summoner.MatchesInfo[1].Queue)
	assert.Equal(test, 13, summoner.MatchesInfo[1].Season)
	assert.Equal(test, 1592876962973, summoner.MatchesInfo[1].Timestamp)
	assert.Equal(test, "SOLO", summoner.MatchesInfo[1].Role)
	assert.Equal(test, "TOP", summoner.MatchesInfo[1].Lane)
	assert.Equal(test, 158, summoner.TotalGames)
}

func TestServiceGetSummonerInfoByName(test *testing.T) {
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
	summoner := rsc_v1.NewSummoner(responseSummonerMock)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.Equal(test, 1, summoner.SummonerLevel)
	assert.Equal(test, "OSIlasjdsanc", summoner.SummonerID)
	assert.Equal(test, "IsolsmdhYDHadKBDA-9fM", summoner.AccountID)
	assert.Equal(test, "jClaj-2S4ZsbjRgIItf1PtjL7-FXbqyDeC", summoner.Puuid)
	assert.Equal(test, 1, summoner.ProfileIconID)
	assert.Equal(test, 1592855746000, summoner.RevisionDate)
}

func TestServiceGetSummonerLeagueByName(test *testing.T) {
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
	summoner := rsc_v1.NewSummoner(responseSummonerMock)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.Equal(test, 1, summoner.SummonerLevel)
	assert.Equal(test, "OSIlasjdsanc", summoner.SummonerID)
	assert.Equal(test, "IsolsmdhYDHadKBDA-9fM", summoner.AccountID)
	assert.Equal(test, "jClaj-2S4ZsbjRgIItf1PtjL7-FXbqyDeC", summoner.Puuid)
	assert.Equal(test, 1, summoner.ProfileIconID)
	assert.Equal(test, 1592855746000, summoner.RevisionDate)
	bodyServiceLeagueByName := []byte(`
		[
			{
				"leagueId": "b80abef5-4fc6-ba48",
				"queueType": "RANKED_SOLO_5x5",
				"tier": "PLATINUM",
				"rank": "II",
				"summonerId": "OSIlasjdsanc",
				"summonerName": "IsBlackPanther",
				"leaguePoints": 25,
				"wins": 17,
				"losses": 23,
				"veteran": false,
				"inactive": false,
				"freshBlood": false,
				"hotStreak": false
			}
		]
	`)
	responseSummonerLeagueMock := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(bodyServiceLeagueByName)),
		ContentLength: int64(len(bodyServiceLeagueByName)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
	summoner.WithLeagueInfo(responseSummonerLeagueMock)
	assert.Equal(test, "b80abef5-4fc6-ba48", summoner.LeagueInfo[0].LeagueID)
	assert.Equal(test, "RANKED_SOLO_5x5", summoner.LeagueInfo[0].QueueType)
	assert.Equal(test, "PLATINUM", summoner.LeagueInfo[0].Tier)
	assert.Equal(test, "II", summoner.LeagueInfo[0].Rank)
	assert.Equal(test, 25, summoner.LeagueInfo[0].LeaguePoints)
	assert.Equal(test, 17, summoner.LeagueInfo[0].Wins)
	assert.Equal(test, 23, summoner.LeagueInfo[0].Losses)
	assert.Equal(test, false, summoner.LeagueInfo[0].Veteran)
	assert.Equal(test, false, summoner.LeagueInfo[0].Inactive)
	assert.Equal(test, false, summoner.LeagueInfo[0].FreshBlood)
	assert.Equal(test, false, summoner.LeagueInfo[0].HotStreak)
	assert.Equal(test, 0, summoner.LeagueInfo[0].MiniSeriesDTO.Losses)
	assert.Equal(test, "", summoner.LeagueInfo[0].MiniSeriesDTO.Progress)
	assert.Equal(test, 0, summoner.LeagueInfo[0].MiniSeriesDTO.Target)
	assert.Equal(test, 0, summoner.LeagueInfo[0].MiniSeriesDTO.Wins)
}

func TestServiceGetSummonerMatchesByName(test *testing.T) {
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
	summoner := rsc_v1.NewSummoner(responseSummonerMock)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.Equal(test, 1, summoner.SummonerLevel)
	assert.Equal(test, "OSIlasjdsanc", summoner.SummonerID)
	assert.Equal(test, "IsolsmdhYDHadKBDA-9fM", summoner.AccountID)
	assert.Equal(test, "jClaj-2S4ZsbjRgIItf1PtjL7-FXbqyDeC", summoner.Puuid)
	assert.Equal(test, 1, summoner.ProfileIconID)
	assert.Equal(test, 1592855746000, summoner.RevisionDate)
	bodyServiceMatchesByName := []byte(`
		{
			"matches": [
				{
					"platformId": "BR1",
					"gameId": 1982082824,
					"champion": 39,
					"queue": 420,
					"season": 13,
					"timestamp": 1592878663522,
					"role": "SOLO",
					"lane": "TOP"
				},
				{
					"platformId": "BR1",
					"gameId": 1981939620,
					"champion": 39,
					"queue": 420,
					"season": 13,
					"timestamp": 1592876962973,
					"role": "SOLO",
					"lane": "TOP"
				}
			],
			"startIndex": 0,
			"endIndex": 50,
			"totalGames": 158
		}
	`)
	responseSummonerMatchesMock := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(bodyServiceMatchesByName)),
		ContentLength: int64(len(bodyServiceMatchesByName)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
	summoner.WithMatchesInfo(responseSummonerMatchesMock)
	assert.Equal(test, "BR1", summoner.MatchesInfo[0].PlatformID)
	assert.Equal(test, 1982082824, summoner.MatchesInfo[0].GameID)
	assert.Equal(test, 39, summoner.MatchesInfo[0].Champion)
	assert.Equal(test, 420, summoner.MatchesInfo[0].Queue)
	assert.Equal(test, 13, summoner.MatchesInfo[0].Season)
	assert.Equal(test, 1592878663522, summoner.MatchesInfo[0].Timestamp)
	assert.Equal(test, "SOLO", summoner.MatchesInfo[0].Role)
	assert.Equal(test, "TOP", summoner.MatchesInfo[0].Lane)
	assert.Equal(test, "BR1", summoner.MatchesInfo[1].PlatformID)
	assert.Equal(test, 1981939620, summoner.MatchesInfo[1].GameID)
	assert.Equal(test, 39, summoner.MatchesInfo[1].Champion)
	assert.Equal(test, 420, summoner.MatchesInfo[1].Queue)
	assert.Equal(test, 13, summoner.MatchesInfo[1].Season)
	assert.Equal(test, 1592876962973, summoner.MatchesInfo[1].Timestamp)
	assert.Equal(test, "SOLO", summoner.MatchesInfo[1].Role)
	assert.Equal(test, "TOP", summoner.MatchesInfo[1].Lane)
	assert.Equal(test, 158, summoner.TotalGames)
}

func TestInvalidSummonerName(test *testing.T) {
	test.Parallel()
	errSummoner1 := svc_v1.CheckSummonerName("&*|ª")
	assert.Nil(test, errSummoner1)
}

func TestMain(test *testing.M) {
	utils.LogOperation.Println("TEstMain")
	os.Exit(test.Run())
}
