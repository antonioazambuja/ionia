package v1

import (
	"encoding/json"
	"log"
	"net/http"
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
	SummonerName  string           `json:"summonerName,omitempty"`
	SummonerLevel int              `json:"summonerLevel,omitempty"`
	LeagueInfo    []LeagueEntryDTO `json:"LeagueInfo,omitempty"`
	// CurrentTier   string `json:"currentTier,omitempty"`
	// CurrentRank   string `json:"currentRank,omitempty"`
	// Wins          int    `json:"wins,omitempty"`
	// Losses        int    `json:"losses,omitempty"`
	// QueueType     string `json:"queueType,omitempty"`
	// LeaguePoints  int    `json:"leaguePoints,omitempty"`
}

type SummonerBuilder struct {
	summonerName             string
	summonerInfo, leagueInfo bool
}

func NewSummonerBuilder(summonerName string) *SummonerBuilder {
	return &SummonerBuilder{
		summonerName: summonerName,
	}
}

// func (builder *SummonerBuilder) SetSummonerInfo(summonerDTO *SummonerDTO) *SummonerBuilder {
func (builder *SummonerBuilder) WithSummonerInfo() *SummonerBuilder {
	builder.summonerInfo = true
	// builder.summonerLevel = summonerDTO.SummonerLevel
	return builder
}

// func (builder *SummonerBuilder) setLeagueInfo(leagueEntryDTO *LeagueEntryDTO) *SummonerBuilder {
func (builder *SummonerBuilder) WithLeagueInfo() *SummonerBuilder {
	builder.leagueInfo = true
	return builder
}

func (builder *SummonerBuilder) Build() (Summoner, error) {
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	var summonerDTO SummonerDTO
	if builder.summonerInfo {
		requestSummoner, errorRequestSummoner := NewRequestBuilder().GetBuilder("summoner").SetPathParam(builder.summonerName).Build()
		if errorRequestSummoner != nil {
			// log.Fatal("Error build request...")
			panic(errorRequestSummoner)
			// return Summoner{}, errorRequestSummoner
		}
		responseSummoner, errResponseSummoner := client.Do(requestSummoner)
		if errResponseSummoner != nil && responseSummoner.StatusCode != 200 {
			log.Fatal("Error build request...")
			panic(errResponseSummoner)
		}
		defer responseSummoner.Body.Close()
		json.NewDecoder(responseSummoner.Body).Decode(&summonerDTO)
	}
	var leagueEntryDTO []LeagueEntryDTO
	if builder.leagueInfo {
		requestLeague, errorRequestLeague := NewRequestBuilder().GetBuilder("league").SetPathParam(summonerDTO.ID).Build()
		if errorRequestLeague != nil {
			log.Fatal("Error build request...")
			panic(errorRequestLeague)
			// return Summoner{}, errorRequestLeague
		}
		responseLeague, errResponseLeague := client.Do(requestLeague)
		if errResponseLeague != nil && responseLeague.StatusCode != 200 {
			log.Fatal("Error build request...")
			panic(errResponseLeague)
		}
		defer responseLeague.Body.Close()
		json.NewDecoder(responseLeague.Body).Decode(&leagueEntryDTO)
	}
	return Summoner{
		SummonerName:  summonerDTO.Name,
		SummonerLevel: summonerDTO.SummonerLevel,
		LeagueInfo:    leagueEntryDTO,
		// CurrentRank:   leagueEntryDTO[0].Rank,
		// CurrentTier:   leagueEntryDTO[0].Tier,
		// LeaguePoints:  leagueEntryDTO[0].LeaguePoints,
		// Wins:          leagueEntryDTO[0].Wins,
		// Losses:        leagueEntryDTO[0].Losses,
		// QueueType:     leagueEntryDTO[0].QueueType,
	}, nil
}
