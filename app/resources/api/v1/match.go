package v1

// MatchlistDto - list of matches from summoner
type MatchlistDto struct {
	StartIndex int                 `json:"startIndex,omitempty"`
	TotalGames int                 `json:"totalGames,omitempty"`
	EndIndex   int                 `json:"endIndex,omitempty"`
	Matches    []MatchReferenceDto `json:"matches,omitempty"`
}

// MatchReferenceDto - reference of the match
type MatchReferenceDto struct {
	GameID     int    `json:"gameId,omitempty"`
	Role       string `json:"role,omitempty"`
	Season     int    `json:"season,omitempty"`
	PlatformID string `json:"platformId,omitempty"`
	Champion   int    `json:"champion,omitempty"`
	Queue      int    `json:"queue,omitempty"`
	Lane       string `json:"lane,omitempty"`
	Timestamp  int    `json:"timestamp,omitempty"`
}
