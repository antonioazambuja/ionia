package v1

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
	SummonerName  string              `json:"summonerName"`
	SummonerLevel int                 `json:"summonerLevel"`
	SummonerID    string              `json:"id"`
	AccountID     string              `json:"accountId"`
	Puuid         string              `json:"puuid"`
	ProfileIconID int                 `json:"profileIconId"`
	RevisionDate  int                 `json:"revisionDate"`
	LeagueInfo    []LeagueInfo        `json:"leagueInfo"`
	TotalGames    int                 `json:"totalGames"`
	MatchesInfo   []MatchReferenceDto `json:"matchesInfo"`
}

// SummonerBuilder - builder summoner
type SummonerBuilder struct {
	summonerDTO SummonerDTO
	summoner    Summoner
	ID          string
}

// NewSummoner - build new summoner
func NewSummoner() *Summoner {
	return &Summoner{
		AccountID:     "",
		LeagueInfo:    []LeagueInfo{},
		MatchesInfo:   []MatchReferenceDto{},
		ProfileIconID: 0,
		Puuid:         "",
		RevisionDate:  0,
		SummonerID:    "",
		SummonerLevel: 0,
		SummonerName:  "",
		TotalGames:    0,
	}
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
