package v1

// MiniSeriesDTO - mini series of response LeagueEntryDTO of Riot API
type MiniSeriesDTO struct {
	Losses   int    `json:"losses,omitempty"`
	Progress string `json:"progress,omitempty"`
	Target   int    `json:"target,omitempty"`
	Wins     int    `json:"wins,omitempty"`
}

// LeagueEntryDTO - league data response of Riot API
type LeagueEntryDTO struct {
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

// LeagueInfo - league data for summoner in custom API
type LeagueInfo struct {
	LeagueID      string        `json:"leagueId,omitempty"`
	QueueType     string        `json:"queueType,omitempty"`
	Tier          string        `json:"tier,omitempty"`
	Rank          string        `json:"rank,omitempty"`
	LeaguePoints  int           `json:"leaguePoints,omitempty"`
	Wins          int           `json:"wins,omitempty"`
	Losses        int           `json:"losses,omitempty"`
	Veteran       bool          `json:"veteran,omitempty"`
	Inactive      bool          `json:"inactive,omitempty"`
	FreshBlood    bool          `json:"freshBlood,omitempty"`
	HotStreak     bool          `json:"hotStreak,omitempty"`
	MiniSeriesDTO MiniSeriesDTO `json:"miniSeries,omitempty"`
}
