package matches

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jeronimobarea/sp_predictions_bot/internal/espn"
	"github.com/openai/openai-go"
)

type Service interface {
	CreateMatch(ctx context.Context, cmd MatchCMD) error
	UpdateMatch(ctx context.Context, match Match) error
	GetMatchByID(ctx context.Context, matchID string) (*Match, error)
	ProcessMatches(ctx context.Context, matchList []espn.ESPNMatch) ([]Match, error)
}

type service struct {
	repo         Repository
	openaiClient *openai.Client
}

func NewService(repo Repository, opeopenaiClient *openai.Client) *service {
	return &service{
		repo:         repo,
		openaiClient: opeopenaiClient,
	}
}

func (svc *service) CreateMatch(ctx context.Context, cmd MatchCMD) error {
	return svc.repo.PersistMatch(ctx, Match{
		ID:        cmd.ID,
		Date:      cmd.Date,
		Name:      cmd.Name,
		TeamA:     cmd.TeamA,
		TeamB:     cmd.TeamB,
		BannerURL: cmd.BannerURL,
	})
}

func (svc *service) GetMatchByID(ctx context.Context, matchID string) (*Match, error) {
	return svc.repo.GetMatchByID(ctx, matchID)
}

func (svc *service) UpdateMatch(ctx context.Context, match Match) error {
	return svc.repo.UpdateMatch(ctx, match)
}

func (svc *service) ProcessMatches(ctx context.Context, matchList []espn.ESPNMatch) ([]Match, error) {
	var matches []Match
	for _, match := range matchList {
		if match.Date.Unix() <= time.Now().Unix() ||
			match.Date.Unix() > time.Now().AddDate(0, 0, 7).Unix() {
			continue
		}

		_, err := svc.GetMatchByID(ctx, match.ID)
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		matchImageURL, err := svc.generateMatchBanner(ctx, match.Esport, match.Name)
		if err != nil {
			return nil, err
		}

		err = svc.CreateMatch(ctx, MatchCMD{
			ID:        match.ID,
			Date:      match.Date,
			Name:      match.Name,
			TeamA:     match.TeamA,
			TeamB:     match.TeamB,
			BannerURL: matchImageURL,
		})
		if err != nil {
			return nil, err
		}

		match, err := svc.GetMatchByID(ctx, match.ID)
		if err != nil {
			return nil, err
		}
		matches = append(matches, *match)
	}
	return matches, nil
}

func (svc *service) generateMatchBanner(ctx context.Context, esport, matchName string) (string, error) {
	var prompt string
	switch esport {
	case espn.BaseballSport:
		prompt = fmt.Sprintf("Design a high-energy baseball match banner featuring %s with both team logos prominently displayed. The background shows a packed baseball stadium under bright lights, with a detailed view of the field and scoreboard. Incorporate baseballs, bats, gloves, and home plate to highlight the baseball theme. Use dynamic lighting, bold colors, and a sense of motion to capture the excitement of a baseball rivalry.", matchName)
	case espn.BasketballSport:
		prompt = fmt.Sprintf("Create an intense sports match banner featuring %s with both team logos displayed prominently. The background shows a vibrant basketball arena, with spotlights and the hardwood court visible. Include basketballs, hoops, and other relevant elements that highlight the basketball theme. Use bold, vivid colors and dynamic design to convey the thrill of a high-stakes game.", matchName)
	case espn.FootballSport:
		prompt = fmt.Sprintf("Create a dynamic sports match banner featuring %s with both players’ names and national flags (if applicable) prominently displayed. The background should show a tennis court with vibrant lighting and a clear view of the net. Add tennis rackets, balls, and sport-specific elements to emphasize the tennis theme. Use vibrant, contrasting colors and an energetic design to reflect the intensity of a high-level match.", matchName)
	case espn.HockeySport:
		prompt = fmt.Sprintf("Create a dynamic sports match banner featuring %s with both players’ names and national flags (if applicable) prominently displayed. The background should show a tennis court with vibrant lighting and a clear view of the net. Add tennis rackets, balls, and sport-specific elements to emphasize the tennis theme. Use vibrant, contrasting colors and an energetic design to reflect the intensity of a high-level match.", matchName)
	case espn.MMASport:
		prompt = fmt.Sprintf("Design an intense hockey match banner featuring %s with both team logos prominently displayed. The background should showcase a packed hockey arena with bright, cool-toned lights reflecting off the ice. Include dynamic elements like hockey sticks, pucks, and goal nets to emphasize the hockey theme. Capture the speed and intensity of the game with bold colors, icy textures, and a sense of movement, creating a high-energy, competitive atmosphere perfect for a hockey rivalry.", matchName)
	default:
		return "", fmt.Errorf("esport type prompt not supported: %s", esport)
	}

	res, err := svc.openaiClient.Images.Generate(ctx, openai.ImageGenerateParams{
		Prompt:  openai.String(prompt),
		Model:   openai.String("dall-e-3"),
		Quality: openai.F(openai.ImageGenerateParamsQualityHD),
	})
	if err != nil {
		return "", err
	}
	return res.Data[0].URL, nil
}
