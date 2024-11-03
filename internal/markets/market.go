package markets

import "time"

type MarketCMD struct {
	GameName   string
	CandidateA string
	CandidateB string
	ExpiryTime time.Time
	LockTime   time.Time
	ImageURL   string
}
