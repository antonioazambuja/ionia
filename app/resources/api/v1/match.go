package v1

// MatchlistDto - list of matches from summoner
type MatchlistDto struct {
	StartIndex int                 `json:"startIndex"`
	TotalGames int                 `json:"totalGames"`
	EndIndex   int                 `json:"endIndex"`
	Matches    []MatchReferenceDto `json:"matches"`
}

// MatchReferenceDto - reference of the match
type MatchReferenceDto struct {
	GameID     int    `json:"gameId"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
	PlatformID string `json:"platformId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Lane       string `json:"lane"`
	Timestamp  int    `json:"timestamp"`
}
