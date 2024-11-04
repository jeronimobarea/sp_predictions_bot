package espn

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jeronimobarea/sp_predictions_bot/internal/espn/client"
)

type Service interface {
	GetNBAScoreboard(ctx context.Context) ([]ESPNMatch, error)
	GetMLBScoreboard(ctx context.Context) ([]ESPNMatch, error)
	GetNFLScoreboard(ctx context.Context) ([]ESPNMatch, error)
	GetUFCScoreboard(ctx context.Context) ([]ESPNMatch, error)
}

type service struct {
	client client.Client
}

func NewService(client client.Client) *service {
	return &service{
		client: client,
	}
}

const RFC3339Z = "2006-01-02T15:04Z07:00"

func (svc *service) GetNBAScoreboard(ctx context.Context) ([]ESPNMatch, error) {
	resp, err := svc.client.GetScoreboard(ctx, BasketballSport, NBALeague)
	if err != nil {
		return nil, err
	}

	return svc.parseMatches(resp.Events)
}

func (svc *service) GetMLBScoreboard(ctx context.Context) ([]ESPNMatch, error) {
	resp, err := svc.client.GetScoreboard(ctx, BaseballSport, MLBLeague)
	if err != nil {
		return nil, err
	}

	return svc.parseMatches(resp.Events)
}

func (svc *service) GetNFLScoreboard(ctx context.Context) ([]ESPNMatch, error) {
	resp, err := svc.client.GetScoreboard(ctx, FootballSport, NFLLeague)
	if err != nil {
		return nil, err
	}

	return svc.parseMatches(resp.Events)
}

func (svc *service) GetUFCScoreboard(ctx context.Context) ([]ESPNMatch, error) {
	resp, err := svc.client.GetScoreboard(ctx, MMASport, UFCLeague)
	if err != nil {
		return nil, err
	}

	var matches []ESPNMatch
	for _, event := range resp.Events {
		for _, fightNight := range event.Competitions {
			date, err := time.Parse(RFC3339Z, fightNight.Date)
			if err != nil {
				return nil, err
			}

			if len(fightNight.Competitors) != 2 {
				return nil, fmt.Errorf("error competitors amount missmatch: %+v", event.Competitions)
			}

			matchName := fmt.Sprintf(
				"%s: %s VS %s",
				event.ShortName,
				fightNight.Competitors[0].Athlete.DisplayName,
				fightNight.Competitors[1].Athlete.DisplayName,
			)

			match := newESPNMatch(
				date,
				matchName,
				fightNight.Competitors[0].Athlete.DisplayName,
				fightNight.Competitors[0].Athlete.Logo,
				fightNight.Competitors[1].Athlete.DisplayName,
				fightNight.Competitors[1].Athlete.Logo,
			)
			matches = append(matches, match)
		}
	}

	return matches, nil
}

func (svc *service) parseMatches(events []client.Event) ([]ESPNMatch, error) {
	var matches []ESPNMatch
	for _, event := range events {
		date, err := time.Parse(RFC3339Z, event.Date)
		if err != nil {
			return nil, err
		}

		if len(event.Competitions) != 1 {
			return nil, fmt.Errorf("error competitions amount missmatch: %+v", event.Competitions)
		}

		if len(event.Competitions[0].Competitors) != 2 {
			return nil, fmt.Errorf("error competitors amount missmatch: %+v", event.Competitions[0].Competitors)
		}

		matchName := strings.Replace(event.Name, " at ", " VS ", 1)

		match := newESPNMatch(
			date,
			matchName,
			event.Competitions[0].Competitors[0].Team.DisplayName,
			event.Competitions[0].Competitors[0].Team.Logo,
			event.Competitions[0].Competitors[1].Team.DisplayName,
			event.Competitions[0].Competitors[1].Team.Logo,
		)
		matches = append(matches, match)
	}

	return matches, nil
}
