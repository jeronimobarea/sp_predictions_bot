package matches

import "time"

type MatchCMD struct {
	ID        string
	Date      time.Time
	Name      string
	TeamA     string
	TeamB     string
	BannerURL string
}

type Match struct {
	ID        string
	Date      time.Time
	Name      string
	TeamA     string
	TeamB     string
	BannerURL string
}
