package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// SummonerDTO - summoner profile response
type SummonerDTO struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

// Summoner - summoner for API
type Summoner struct {
	SummonerName  string       `json:"summonerName"`
	SummonerLevel int          `json:"summonerLevel"`
	SummonerID    string       `json:"id"`
	AccountID     string       `json:"accountId"`
	Puuid         string       `json:"puuid"`
	ProfileIconID int          `json:"profileIconId"`
	RevisionDate  int          `json:"revisionDate"`
	LeagueInfo    []LeagueInfo `json:"leagueInfo"`
}

// SummonerBuilder - builder summoner
type SummonerBuilder struct {
	summonerName             string
	summonerInfo, leagueInfo bool
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

// Build - create and get data in Riot API
func (builder *SummonerBuilder) Build() (Summoner, error) {
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	var summonerDTO SummonerDTO
	if builder.summonerInfo {
		requestSummoner, errorRequestSummoner := NewRequestBuilder().TypeBuilder("summoner").WithPathParam(builder.summonerName).Build()
		if errorRequestSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print(errorRequestSummoner)
			return Summoner{}, errorRequestSummoner
		}
		responseSummoner, errResponseSummoner := client.Do(requestSummoner)
		if errResponseSummoner != nil || responseSummoner.StatusCode != 200 {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print(errResponseSummoner)
		}
		defer responseSummoner.Body.Close()
		json.NewDecoder(responseSummoner.Body).Decode(&summonerDTO)
	}
	var leagueInfos []LeagueInfo
	if builder.leagueInfo {
		requestLeague, errorRequestLeague := NewRequestBuilder().TypeBuilder("league").WithPathParam(summonerDTO.ID).Build()
		if errorRequestLeague != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print(errorRequestLeague)
			return Summoner{}, errorRequestLeague
		}
		responseLeague, errResponseLeague := client.Do(requestLeague)
		if errResponseLeague != nil || responseLeague.StatusCode != 200 {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print(errResponseLeague)
		}
		defer responseLeague.Body.Close()
		var leagueEntryDTO []LeagueEntryDTO
		json.NewDecoder(responseLeague.Body).Decode(&leagueEntryDTO)
		for _, info := range leagueEntryDTO {
			leagueInfo := LeagueInfo{
				LeagueID: info.LeagueID,
				QueueType: info.QueueType,
				Tier: info.Tier,
				Rank: info.Rank,
				LeaguePoints: info.LeaguePoints,
				Wins: info.Wins,
				Losses: info.Losses,
				Veteran: info.Veteran,
				Inactive: info.Inactive,
				FreshBlood: info.FreshBlood,
				HotStreak: info.HotStreak,
				MiniSeriesDTO: info.MiniSeriesDTO,
			}
			leagueInfos = append(leagueInfos, leagueInfo)
		}
	}
	return Summoner{
		SummonerName:  summonerDTO.Name,
		SummonerLevel: summonerDTO.SummonerLevel,
		SummonerID:    summonerDTO.ID,
		AccountID:     summonerDTO.AccountID,
		Puuid:         summonerDTO.Puuid,
		ProfileIconID: summonerDTO.ProfileIconID,
		RevisionDate:  summonerDTO.RevisionDate,
		LeagueInfo:    leagueInfos,
	}, nil
}
