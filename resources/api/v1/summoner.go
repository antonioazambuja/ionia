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
	SummonerName  string       `json:"summonerName,omitempty"`
	SummonerLevel int          `json:"summonerLevel,omitempty"`
	SummonerID    string       `json:"id,omitempty"`
	AccountID     string       `json:"accountId,omitempty"`
	Puuid         string       `json:"puuid,omitempty"`
	ProfileIconID int          `json:"profileIconId,omitempty"`
	RevisionDate  int          `json:"revisionDate,omitempty"`
	LeagueInfo    []LeagueInfo `json:"leagueInfo,omitempty"`
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
	var leagueInfo []LeagueInfo
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
		json.NewDecoder(responseLeague.Body).Decode(&leagueInfo)
	}
	return Summoner{
		SummonerName:  summonerDTO.Name,
		SummonerLevel: summonerDTO.SummonerLevel,
		SummonerID:    summonerDTO.ID,
		AccountID:     summonerDTO.AccountID,
		Puuid:         summonerDTO.Puuid,
		ProfileIconID: summonerDTO.ProfileIconID,
		RevisionDate:  summonerDTO.RevisionDate,
		LeagueInfo:    leagueInfo,
	}, nil
}
