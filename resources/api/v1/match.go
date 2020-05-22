package v1

// MiniSeriesDTO -
type MiniSeriesDTO struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

// MatchlistDto - list of matches from summoner
type MatchlistDto struct {
	StartIndex int                 `json:"startIndex"`
	TotalGames int                 `json:"totalGames"`
	EndIndex   int                 `json:"endIndex"`
	Matches    []MatchReferenceDto `json:"matches"`
}

// MatchReferenceDto - reference of the match
type MatchReferenceDto struct {
	GameID     int64  `json:"gameId"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
	PlatformID string `json:"platformId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Lane       string `json:"lane"`
	Timestamp  int64  `json:"timestamp"`
}
