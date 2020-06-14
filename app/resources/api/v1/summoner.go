package v1

import (
	"encoding/json"

	utils "github.com/antonioazambuja/ionia/utils"
)

const summonerV4 string = "/lol/summoner/v4/summoners/by-name/"
const leagueV4 string = "/lol/league/v4/entries/by-summoner/"
const matchesV4 string = "/lol/match/v4/matchlists/by-account/"

// SummonerDTO - summoner profile response
type SummonerDTO struct {
	ID            string `json:"id,omitempty"`
	AccountID     string `json:"accountId,omitempty"`
	Puuid         string `json:"puuid,omitempty"`
	Name          string `json:"name,omitempty"`
	ProfileIconID int    `json:"profileIconId,omitempty"`
	RevisionDate  int    `json:"revisionDate,omitempty"`
	SummonerLevel int    `json:"summonerLevel,omitempty"`
}

// Summoner - summoner for API
type Summoner struct {
	SummonerName  string              `json:"summonerName,omitempty"`
	SummonerLevel int                 `json:"summonerLevel,omitempty"`
	SummonerID    string              `json:"id,omitempty"`
	AccountID     string              `json:"accountId,omitempty"`
	Puuid         string              `json:"puuid,omitempty"`
	ProfileIconID int                 `json:"profileIconId,omitempty"`
	RevisionDate  int                 `json:"revisionDate,omitempty"`
	LeagueInfo    []LeagueInfo        `json:"leagueInfo,omitempty"`
	TotalGames    int                 `json:"totalGames,omitempty"`
	MatchesInfo   []MatchReferenceDto `json:"matchesInfo,omitempty"`
}

// SummonerBuilder - builder summoner
type SummonerBuilder struct {
	summonerDTO SummonerDTO
	summoner    Summoner
}

// NewSummonerBuilder - initialize SummonerBuilder
func NewSummonerBuilder(summonerName string) *SummonerBuilder {
	var summonerDTO SummonerDTO
	responseSummoner, errorResponseSummoner := NewRequestBuilder(summonerV4).WithPathParam(summonerName).Run()
	if errorResponseSummoner != nil {
		utils.LogOperation.Print("Failed build summoner, get summoners info")
		return &SummonerBuilder{}
	}
	defer responseSummoner.Body.Close()
	json.NewDecoder(responseSummoner.Body).Decode(&summonerDTO)
	return &SummonerBuilder{
		summonerDTO: summonerDTO,
	}
}

// WithSummonerInfo - add SummonerDTO data in summoner
func (builder *SummonerBuilder) WithSummonerInfo() *SummonerBuilder {
	builder.summoner.SummonerName = builder.summonerDTO.Name
	builder.summoner.SummonerLevel = builder.summonerDTO.SummonerLevel
	builder.summoner.SummonerID = builder.summonerDTO.ID
	builder.summoner.AccountID = builder.summonerDTO.AccountID
	builder.summoner.Puuid = builder.summonerDTO.Puuid
	builder.summoner.ProfileIconID = builder.summonerDTO.ProfileIconID
	builder.summoner.RevisionDate = builder.summonerDTO.RevisionDate
	return builder
}

// WithLeagueInfo - add LeagueEntryDTO data in summoner
func (builder *SummonerBuilder) WithLeagueInfo() *SummonerBuilder {
	var leagueInfos []LeagueInfo
	responseLeague, errorResponseLeague := NewRequestBuilder(leagueV4).WithPathParam(builder.summonerDTO.ID).Run()
	if errorResponseLeague != nil {
		utils.LogOperation.Print("Failed build summoner, get league info")
		return &SummonerBuilder{}
	}
	defer responseLeague.Body.Close()
	var leagueEntryDTO []LeagueEntryDTO
	json.NewDecoder(responseLeague.Body).Decode(&leagueEntryDTO)
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
	builder.summoner.LeagueInfo = leagueInfos
	return builder
}

// WithMatchesInfo - add MatchReferenceDto data in summoner
func (builder *SummonerBuilder) WithMatchesInfo() *SummonerBuilder {
	var matchlistDto MatchlistDto
	responseMatches, errorResponseMatches := NewRequestBuilder(matchesV4).WithPathParam(builder.summonerDTO.AccountID).WithQueries([]string{"beginIndex", "endIndex"}, []string{"0", "15"}).Run()
	if errorResponseMatches != nil {
		utils.LogOperation.Print("Failed build summoner, get matches info")
		return &SummonerBuilder{}
	}
	defer responseMatches.Body.Close()
	json.NewDecoder(responseMatches.Body).Decode(&matchlistDto)
	builder.summoner.MatchesInfo = matchlistDto.Matches
	builder.summoner.TotalGames = matchlistDto.TotalGames
	return builder
}

// Build - create and get data in Riot API
func (builder *SummonerBuilder) Build() (Summoner, error) {
	return builder.summoner, nil
}
