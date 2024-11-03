package client

type (
	ScoreboardResponse struct {
		Events []Event `json:"events"`
	}

	Event struct {
		ID           string        `json:"id"`
		Date         string        `json:"date"`
		Name         string        `json:"name"`
		Competitions []Competition `json:"competitions"`
	}

	Competition struct {
		Competitors []Competitor `json:"competitors"`
	}

	Competitor struct {
		Team Team `json:"team"`
	}

	Team struct {
		DisplayName string `json:"displayName"`
		Logo        string `json:"logo"`
	}
)
