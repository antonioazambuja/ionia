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

func TestGetSummonerByNameInvalidSummonerName(test *testing.T) {
	test.Parallel()
	summoner2, errSummoner2 := svc_v1.GetByName("&*|ª")
	assert.NotEmpty(test, errSummoner2)
	assert.Empty(test, summoner2)
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
