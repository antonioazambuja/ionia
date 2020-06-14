package v1

import (
	"testing"

	svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
	assert "github.com/stretchr/testify/assert"
)

func TestGetSummonerByName(test *testing.T) {
	test.Parallel()
	summoner1, errSummoner1 := svc_v1.GetByName("IsBlackPanther")
	assert.Empty(test, errSummoner1)
	assert.Equal(test, "IsBlackPanther", summoner1.SummonerName)
	assert.NotEmpty(test, summoner1.AccountID)
	assert.NotEmpty(test, summoner1.LeagueInfo)
	assert.NotEmpty(test, summoner1.MatchesInfo)
}
