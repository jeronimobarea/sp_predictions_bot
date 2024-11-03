package espn

import "time"

const (
	BasketballSport = "basketball"
	BaseballSport   = "baseball"
	FootballSport   = "football"

	NBALeague = "nba"
	MLBLeague = "mlb"
	NFLLeague = "nfl"
)

type ESPNMatch struct {
	ID        string
	Date      time.Time
	Name      string
	TeamA     string
	TeamALogo string
	TeamB     string
	TeamBLogo string
}

func newESPNMatch(
	id string,
	date time.Time,
	name string,
	teamA string,
	teamALogo string,
	teamB string,
	teamBLogo string,
) ESPNMatch {
	return ESPNMatch{
		ID:        id,
		Date:      date,
		Name:      name,
		TeamA:     teamA,
		TeamALogo: teamALogo,
		TeamB:     teamB,
		TeamBLogo: teamBLogo,
	}
}
