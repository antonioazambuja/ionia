package v1

import (
	"github.com/antonioazambuja/ionia/utils"
	"encoding/json"
	"net/http"
)


// MiniSeriesDTO -
type MiniSeriesDTO struct {
	Losses   int    `json:"losses,omitempty"`
	Progress string `json:"progress,omitempty"`
	Target   int    `json:"target,omitempty"`
	Wins     int    `json:"wins,omitempty"`
}

// LeagueEntryDTO - league data response of Riot API
type LeagueEntryDTO struct {
	LeagueID      string        `json:"leagueId,omitempty"`
	QueueType     string        `json:"queueType,omitempty"`
	Tier          string        `json:"tier,omitempty"`
	Rank          string        `json:"rank,omitempty"`
	SummonerID    string        `json:"summonerId,omitempty"`
	SummonerName  string        `json:"summonerName,omitempty"`
	LeaguePoints  int           `json:"leaguePoints,omitempty"`
	Wins          int           `json:"wins,omitempty"`
	Losses        int           `json:"losses,omitempty"`
	Veteran       bool          `json:"veteran,omitempty"`
	Inactive      bool          `json:"inactive,omitempty"`
	FreshBlood    bool          `json:"freshBlood,omitempty"`
	HotStreak     bool          `json:"hotStreak,omitempty"`
	MiniSeriesDTO MiniSeriesDTO `json:"miniSeries,omitempty"`
}

// LeagueInfo - league data for summoner in custom API
type LeagueInfo struct {
	LeagueID      string        `json:"leagueId,omitempty"`
	QueueType     string        `json:"queueType,omitempty"`
	Tier          string        `json:"tier,omitempty"`
	Rank          string        `json:"rank,omitempty"`
	LeaguePoints  int           `json:"leaguePoints,omitempty"`
	Wins          int           `json:"wins,omitempty"`
	Losses        int           `json:"losses,omitempty"`
	Veteran       bool          `json:"veteran,omitempty"`
	Inactive      bool          `json:"inactive,omitempty"`
	FreshBlood    bool          `json:"freshBlood,omitempty"`
	HotStreak     bool          `json:"hotStreak,omitempty"`
	MiniSeriesDTO MiniSeriesDTO `json:"miniSeries,omitempty"`
}

// WithLeagueInfo - add LeagueEntryDTO data in summoner
func (summoner *Summoner) WithLeagueInfo(summonerLeagueHTTPResponse *http.Response) {
	var leagueInfos []LeagueInfo
	var leagueEntryDTO []LeagueEntryDTO
	if errDecodeLeagueResponse := json.NewDecoder(summonerLeagueHTTPResponse.Body).Decode(&leagueEntryDTO); errDecodeLeagueResponse != nil {
		utils.LogOperation.Println(errDecodeLeagueResponse)
	}
	for _, info := range leagueEntryDTO {
		leagueInfo := LeagueInfo{
			LeagueID:      info.LeagueID,
			QueueType:     info.QueueType,
			Tier:          info.Tier,
			Rank:          info.Rank,
			LeaguePoints:  info.LeaguePoints,
			Wins:          info.Wins,
			Losses:        info.Losses,
			Veteran:       info.Veteran,
			Inactive:      info.Inactive,
			FreshBlood:    info.FreshBlood,
			HotStreak:     info.HotStreak,
			MiniSeriesDTO: info.MiniSeriesDTO,
		}
		leagueInfos = append(leagueInfos, leagueInfo)
	}
	summoner.LeagueInfo = leagueInfos
}
