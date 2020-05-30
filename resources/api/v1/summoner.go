package v1

import (
	"encoding/json"
	"log"
	"os"
)

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
	summonerName                          string
	summonerInfo, leagueInfo, matchesInfo bool
	keys, values                          []string
}

// NewSummonerBuilder - initialize SummonerBuilder
func NewSummonerBuilder(summonerName string) *SummonerBuilder {
	return &SummonerBuilder{
		summonerName: summonerName,
	}
}

// WithSummonerInfo - add SummonerDTO data in summoner
func (builder *SummonerBuilder) WithSummonerInfo() *SummonerBuilder {
	builder.summonerInfo = true
	// builder.summonerLevel = summonerDTO.SummonerLevel
	return builder
}

// WithLeagueInfo - add LeagueEntryDTO data in summoner
func (builder *SummonerBuilder) WithLeagueInfo() *SummonerBuilder {
	builder.leagueInfo = true
	return builder
}

// WithMatchesInfo - add MatchReferenceDto data in summoner
func (builder *SummonerBuilder) WithMatchesInfo() *SummonerBuilder {
	builder.matchesInfo = true
	return builder
}

// Build - create and get data in Riot API
func (builder *SummonerBuilder) Build() (Summoner, error) {
	var summoner Summoner
	var summonerDTO SummonerDTO
	responseSummoner, errorResponseSummoner := NewRequestBuilder().TypeBuilder("summoner").WithPathParam(builder.summonerName).Build().Run()
	if errorResponseSummoner != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed build summoner, get summoners info")
		return Summoner{}, errorResponseSummoner
	}
	defer responseSummoner.Body.Close()
	json.NewDecoder(responseSummoner.Body).Decode(&summonerDTO)
	if builder.summonerInfo {
		summoner.SummonerName = summonerDTO.Name
		summoner.SummonerLevel = summonerDTO.SummonerLevel
		summoner.SummonerID = summonerDTO.ID
		summoner.AccountID = summonerDTO.AccountID
		summoner.Puuid = summonerDTO.Puuid
		summoner.ProfileIconID = summonerDTO.ProfileIconID
		summoner.RevisionDate = summonerDTO.RevisionDate
	}
	var leagueInfos []LeagueInfo
	if builder.leagueInfo {
		responseLeague, errorResponseLeague := NewRequestBuilder().TypeBuilder("league").WithPathParam(summonerDTO.ID).Build().Run()
		if errorResponseLeague != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Failed build summoner, get league info")
			return Summoner{}, errorResponseLeague
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
		summoner.LeagueInfo = leagueInfos
	}
	var matchlistDto MatchlistDto
	if builder.matchesInfo {
		responseMatches, errorResponseMatches := NewRequestBuilder().TypeBuilder("matches").WithPathParam(summonerDTO.AccountID).WithQueries([]string{"beginIndex", "endIndex"}, []string{"0", "15"}).Build().Run()
		if errorResponseMatches != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Failed build summoner, get matches info")
			return Summoner{}, errorResponseMatches
		}
		defer responseMatches.Body.Close()
		json.NewDecoder(responseMatches.Body).Decode(&matchlistDto)
		summoner.MatchesInfo = matchlistDto.Matches
		summoner.TotalGames = matchlistDto.TotalGames
	}
	return summoner, nil
}
