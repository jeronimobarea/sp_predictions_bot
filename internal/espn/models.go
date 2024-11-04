package espn

import "time"

const (
	BasketballSport = "basketball"
	BaseballSport   = "baseball"
	FootballSport   = "football"
	MMASport        = "mma"

	NBALeague = "nba"
	MLBLeague = "mlb"
	NFLLeague = "nfl"
	UFCLeague = "ufc"
)

type ESPNMatch struct {
	Date      time.Time
	Name      string
	TeamA     string
	TeamALogo string
	TeamB     string
	TeamBLogo string
}

func newESPNMatch(
	date time.Time,
	name string,
	teamA string,
	teamALogo string,
	teamB string,
	teamBLogo string,
) ESPNMatch {
	return ESPNMatch{
		Date:      date,
		Name:      name,
		TeamA:     teamA,
		TeamALogo: teamALogo,
		TeamB:     teamB,
		TeamBLogo: teamBLogo,
	}
}
