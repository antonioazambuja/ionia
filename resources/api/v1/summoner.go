package v1

// SummonerResponse - summoner profile response
type SummonerResponse struct {
	ID            string `json:"id,omitempty"`
	AccountID     string `json:"accountId,omitempty"`
	Puuid         string `json:"puuid,omitempty"`
	Name          string `json:"name,omitempty"`
	ProfileIconID int    `json:"profileIconId,omitempty"`
	RevisionDate  int    `json:"revisionDate,omitempty"`
	SummonerLevel int    `json:"summonerLevel,omitempty"`
}

// LeagueEntriesResponse - data league of summoner
type LeagueEntriesResponse struct {
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

// MiniSeriesDTO -
type MiniSeriesDTO struct {
	Losses   int    `json:"losses,omitempty"`
	Progress string `json:"progress,omitempty"`
	Target   int    `json:"target,omitempty"`
	Wins     int    `json:"wins,omitempty"`
}

// Summoner - summoner for API
type Summoner struct {
	SummonerName  string `json:"summonerName,omitempty"`
	SummonerLevel int    `json:"summonerLevel,omitempty"`
	CurrentTier   string `json:"currentTier,omitempty"`
	CurrentRank   string `json:"currentRank,omitempty"`
	Wins          int    `json:"wins,omitempty"`
	Losses        int    `json:"losses,omitempty"`
	QueueType     string `json:"queueType,omitempty"`
	LeaguePoints  int    `json:"leaguePoints,omitempty"`
}
