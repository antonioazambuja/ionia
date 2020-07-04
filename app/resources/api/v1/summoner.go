package v1

import (
	"context"
	"encoding/json"
	"fmt"

	utils "github.com/antonioazambuja/ionia/utils"
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
	summonerDTO SummonerDTO
	summoner    Summoner
	ID          string
}

// NewCacheSummoner - initialize SummonerBuilder
func NewCacheSummoner(summoner *Summoner, serviceID string) error {
	summonerJSON, errParseStructToJSON := json.Marshal(summoner)
	if errParseStructToJSON != nil {
		utils.LogOperation.Println("Failed cached summoner of id service: " + serviceID)
		utils.LogOperation.Println(errParseStructToJSON.Error())
		return errParseStructToJSON
	}
	setSummonerRedisResult, errSetSummonerRedisResult := GetConn().Do(context.TODO(), "SET", fmt.Sprint(summoner.SummonerName+"_"+serviceID), summonerJSON).Result()
	if errSetSummonerRedisResult != nil {
		utils.LogOperation.Println("Failed cached summoner of id service: " + serviceID)
		utils.LogOperation.Println(errSetSummonerRedisResult.Error())
		return errSetSummonerRedisResult
	}
	utils.LogOperation.Printf("Succesfull cached summoner of id service: %s. Result Redis: %s\n", summoner.SummonerName+"_"+serviceID, setSummonerRedisResult)
	return nil
}

// GetCacheSummoner - initialize SummonerBuilder
func GetCacheSummoner(summonerName string, serviceID string) (Summoner, error) {
	var summoner Summoner
	summonerCacheRedis, errGetSummonerCacheRedis := GetConn().Get(context.TODO(), summonerName+"_"+serviceID).Result()
	if errGetSummonerCacheRedis != nil {
		utils.LogOperation.Print("Not found cache data in Redis - errGetSummonerCacheRedis. Result Redis: " + errGetSummonerCacheRedis.Error())
		return Summoner{}, errGetSummonerCacheRedis
	}
	errParseJSONToStruct := json.Unmarshal([]byte(summonerCacheRedis), &summoner)
	if errParseJSONToStruct != nil {
		utils.LogOperation.Println("Error found cache data in Redis - errParseJSONToStruct. Result Redis: " + errParseJSONToStruct.Error())
		return Summoner{}, errParseJSONToStruct
	}
	utils.LogOperation.Print("Found cache data in Redis")
	return summoner, nil
}

func (summoner *Summoner) WithSummonerInfo(summonerDTO *SummonerDTO) {
	summoner.SummonerName = summonerDTO.Name
	summoner.SummonerLevel = summonerDTO.SummonerLevel
	summoner.SummonerID = summonerDTO.ID
	summoner.AccountID = summonerDTO.AccountID
	summoner.Puuid = summonerDTO.Puuid
	summoner.ProfileIconID = summonerDTO.ProfileIconID
	summoner.RevisionDate = summonerDTO.RevisionDate
}

// WithLeagueInfo - add LeagueEntryDTO data in summoner
func (summoner *Summoner) WithLeagueInfo(leagueEntryDTO []LeagueEntryDTO) {
	var leagueInfos []LeagueInfo
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

// WithMatchesInfo - add MatchReferenceDto data in summoner
func (summoner *Summoner) WithMatchesInfo(matchlistDto *MatchlistDto) {
	summoner.MatchesInfo = matchlistDto.Matches
	summoner.TotalGames = matchlistDto.TotalGames
}
