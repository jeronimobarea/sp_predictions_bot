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
	ID        string
	Date      time.Time
	Esport    string
	Name      string
	TeamA     string
	TeamB     string
	BannerURL string
}

func newESPNMatch(
	id string,
	date time.Time,
	esport string,
	name string,
	teamA string,
	teamB string,
	bannerURL string,
) ESPNMatch {
	return ESPNMatch{
		ID:        id,
		Date:      date,
		Esport:    esport,
		Name:      name,
		TeamA:     teamA,
		TeamB:     teamB,
		BannerURL: bannerURL,
	}
}
