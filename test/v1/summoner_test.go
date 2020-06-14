package v1

import (
	"testing"

	svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
	assert "github.com/stretchr/testify/assert"
)

func TestGetSummonerByName(test *testing.T) {
	test.Parallel()
	summoner, errSummoner := svc_v1.GetByName("IsBlackPanther")
	assert.Empty(test, errSummoner)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.NotEmpty(test, summoner.AccountID)
	assert.NotEmpty(test, summoner.LeagueInfo)
	assert.NotEmpty(test, summoner.MatchesInfo)
}

func TestGetSummonerInfoByName(test *testing.T) {
	test.Parallel()
	summoner, errSummoner := svc_v1.GetInfoByName("IsBlackPanther")
	assert.Empty(test, errSummoner)
	assert.Equal(test, "IsBlackPanther", summoner.SummonerName)
	assert.NotEmpty(test, summoner.AccountID)
	assert.Empty(test, summoner.LeagueInfo)
	assert.Empty(test, summoner.MatchesInfo)
}

func TestGetSummonerLeagueByName(test *testing.T) {
	test.Parallel()
	summoner, errSummoner := svc_v1.GetLeagueByName("IsBlackPanther")
	assert.Empty(test, errSummoner)
	assert.Empty(test, summoner.AccountID)
	assert.NotEmpty(test, summoner.LeagueInfo)
	assert.Empty(test, summoner.MatchesInfo)
}

func TestGetSummonerMatchesByName(test *testing.T) {
	test.Parallel()
	summoner, errSummoner := svc_v1.GetMatchesByName("IsBlackPanther")
	assert.Empty(test, errSummoner)
	assert.Empty(test, summoner.AccountID)
	assert.Empty(test, summoner.LeagueInfo)
	assert.NotEmpty(test, summoner.MatchesInfo)
}

func TestInvalidSummonerName(test *testing.T) {
	test.Parallel()
	summoner1, errSummoner1 := svc_v1.GetByName("&*|ª")
	assert.NotEmpty(test, errSummoner1)
	assert.Empty(test, summoner1)
	summoner2, errSummoner2 := svc_v1.GetInfoByName("&*|ª")
	assert.NotEmpty(test, errSummoner2)
	assert.Empty(test, summoner2)
	summoner3, errSummoner3 := svc_v1.GetLeagueByName("&*|ª")
	assert.NotEmpty(test, errSummoner3)
	assert.Empty(test, summoner3)
	summoner4, errSummoner4 := svc_v1.GetMatchesByName("&*|ª")
	assert.NotEmpty(test, errSummoner4)
	assert.Empty(test, summoner4)
}
