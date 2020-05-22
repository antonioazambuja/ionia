package v1

// LeagueEntryDTO - league data response of Riot API
type LeagueEntryDTO struct {
	LeagueID      string        `json:"leagueId"`
	QueueType     string        `json:"queueType"`
	Tier          string        `json:"tier"`
	Rank          string        `json:"rank"`
	SummonerID    string        `json:"summonerId"`
	SummonerName  string        `json:"summonerName"`
	LeaguePoints  int           `json:"leaguePoints"`
	Wins          int           `json:"wins"`
	Losses        int           `json:"losses"`
	Veteran       bool          `json:"veteran"`
	Inactive      bool          `json:"inactive"`
	FreshBlood    bool          `json:"freshBlood"`
	HotStreak     bool          `json:"hotStreak"`
	MiniSeriesDTO MiniSeriesDTO `json:"miniSeries"`
}

// LeagueInfo - league data for summoner in custom API
type LeagueInfo struct {
	LeagueID      string        `json:"leagueId"`
	QueueType     string        `json:"queueType"`
	Tier          string        `json:"tier"`
	Rank          string        `json:"rank"`
	LeaguePoints  int           `json:"leaguePoints"`
	Wins          int           `json:"wins"`
	Losses        int           `json:"losses"`
	Veteran       bool          `json:"veteran"`
	Inactive      bool          `json:"inactive"`
	FreshBlood    bool          `json:"freshBlood"`
	HotStreak     bool          `json:"hotStreak"`
	MiniSeriesDTO MiniSeriesDTO `json:"miniSeries"`
}
