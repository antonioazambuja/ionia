package v1

// Summoner - gamer profile
type Summoner struct {
	ID            string `json:"id,omitempty"`
	AccountID     string `json:"accountId,omitempty"`
	Puuid         string `json:"puuid,omitempty"`
	Name          string `json:"name,omitempty"`
	ProfileIconID int    `json:"profileIconID,omitempty"`
	RevisionDate  int    `json:"revisionDate,omitempty"`
	SummonerLevel int    `json:"summonerLevel,omitempty"`
}
