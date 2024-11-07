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
	Duration  time.Duration
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
	var duration time.Duration
	switch esport {
	case BaseballSport:
		duration = 3 * time.Hour
	case FootballSport:
		duration = 72 * time.Minute
	case BasketballSport:
		duration = 3 * time.Hour
	case MMASport:
		duration = 30 * time.Minute
	default:
		duration = time.Hour
	}

	return ESPNMatch{
		ID:        id,
		Date:      date,
		Esport:    esport,
		Name:      name,
		TeamA:     teamA,
		TeamB:     teamB,
		Duration:  duration,
		BannerURL: bannerURL,
	}
}
