package client

type (
	ScoreboardResponse struct {
		Events []Event `json:"events"`
	}

	Event struct {
		Date         string        `json:"date"`
		Name         string        `json:"name"`
		ShortName    string        `json:"shortName"`
		Competitions []Competition `json:"competitions"`
	}

	Competition struct {
		ID          string       `json:"id"`
		Date        string       `json:"date,omitempty"`
		Competitors []Competitor `json:"competitors"`
	}

	Competitor struct {
		Team    Team    `json:"team,omitempty"`
		Athlete Athlete `json:"athlete,omitempty"`
	}

	Team struct {
		DisplayName string `json:"displayName"`
		Logo        string `json:"logo"`
	}

	Athlete struct {
		DisplayName string `json:"displayName"`
		Logo        string `json:"logo"`
	}
)
