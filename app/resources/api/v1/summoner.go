package v1

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

// WithSummonerInfo - add SummonerDTO data in summoner
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
