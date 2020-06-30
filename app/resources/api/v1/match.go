package v1

import (
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

// MatchlistDto - list of matches from summoner
type MatchlistDto struct {
	StartIndex int                 `json:"startIndex,omitempty"`
	TotalGames int                 `json:"totalGames,omitempty"`
	EndIndex   int                 `json:"endIndex,omitempty"`
	Matches    []MatchReferenceDto `json:"matches,omitempty"`
}

// MatchReferenceDto - reference of the match
type MatchReferenceDto struct {
	GameID     int64  `json:"gameId,omitempty"`
	Role       string `json:"role,omitempty"`
	Season     int    `json:"season,omitempty"`
	PlatformID string `json:"platformId,omitempty"`
	Champion   int    `json:"champion,omitempty"`
	Queue      int    `json:"queue,omitempty"`
	Lane       string `json:"lane,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
}

// WithMatchesInfo - add MatchReferenceDto data in summoner
func (summoner *Summoner) WithMatchesInfo(summonerHTTPMatchesResponse *http.Response) {
	var matchlistDto MatchlistDto
	defer summonerHTTPMatchesResponse.Body.Close()
	json.NewDecoder(summonerHTTPMatchesResponse.Body).Decode(&matchlistDto)
	summoner.MatchesInfo = matchlistDto.Matches
	summoner.TotalGames = matchlistDto.TotalGames
}
