package markets

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/openai/openai-go"
	"go.uber.org/zap"

	"github.com/jeronimobarea/sp_predictions_bot/internal/espn"
	"github.com/jeronimobarea/sp_predictions_bot/internal/matches"
	sankoPredicts "github.com/jeronimobarea/sp_predictions_bot/pkg/sankopredicts"
)

type Service interface {
	CreateMarket(ctx context.Context, cmd MarketCMD) error
	GetMarkets(ctx context.Context, status int64) ([]sankoPredicts.SankoPredictsGame, error)
	ProcessMatches(ctx context.Context, matches []espn.ESPNMatch) error
}

type service struct {
	transactor    *bind.TransactOpts
	sankoPredicts *sankoPredicts.SankoPredicts
	matchesSvc    matches.Service
	openaiClient  *openai.Client
	logger        *zap.SugaredLogger
}

func NewService(
	transactor *bind.TransactOpts,
	sankoPredicts *sankoPredicts.SankoPredicts,
	matchesSvc matches.Service,
	openaiClient *openai.Client,
	logger *zap.SugaredLogger,
) *service {
	return &service{
		transactor:    transactor,
		sankoPredicts: sankoPredicts,
		matchesSvc:    matchesSvc,
		openaiClient:  openaiClient,
		logger:        logger,
	}
}

const (
	ElectionGameType = "Election"

	MarketStatusPending  int64 = 1
	MarketStatusRejected int64 = 2
	MarketStatusApproved int64 = 3
)

func (svc *service) CreateMarket(ctx context.Context, cmd MarketCMD) error {
	pendingMarkets, err := svc.GetMarkets(ctx, MarketStatusPending)
	if err != nil {
		return err
	}

	rejectedMarkets, err := svc.GetMarkets(ctx, MarketStatusRejected)
	if err != nil {
		return err
	}

	approvedMarkets, err := svc.GetMarkets(ctx, MarketStatusApproved)
	if err != nil {
		return err
	}

	var markets []sankoPredicts.SankoPredictsGame
	markets = append(markets, pendingMarkets...)
	markets = append(markets, rejectedMarkets...)
	markets = append(markets, approvedMarkets...)

	for _, market := range markets {
		if market.CandidateA == cmd.CandidateA &&
			market.CandidateB == cmd.CandidateB &&
			market.GameName == cmd.GameName &&
			market.Creator == svc.transactor.From {

			svc.logger.Warnw("duplicated market",
				"market", market,
			)
			continue
		}

		tx, err := svc.sankoPredicts.CreateElectionGame(
			svc.transactor,
			cmd.GameName,
			big.NewInt(cmd.ExpiryTime.Unix()),
			big.NewInt(cmd.LockTime.Unix()),
			ElectionGameType,
			cmd.CandidateA,
			cmd.CandidateB,
			cmd.ImageURL,
			svc.transactor.From,
		)
		if err != nil {
			return err
		}
		svc.logger.Infow("market created",
			"TransactionID", tx.Hash().Hex(),
			"Market", cmd,
		)
	}
	return nil
}

func (svc *service) GetMarkets(ctx context.Context, status int64) ([]sankoPredicts.SankoPredictsGame, error) {
	return svc.sankoPredicts.GetGames(
		&bind.CallOpts{
			Pending:     false,
			From:        svc.transactor.From,
			BlockNumber: nil, // latest
			Context:     ctx,
		},
		false,
		big.NewInt(status),
		svc.transactor.From,
	)
}

func (svc *service) ProcessMatches(ctx context.Context, matchList []espn.ESPNMatch) error {
	for _, match := range matchList {
		if match.Date.Unix() <= time.Now().Unix() ||
			match.Date.Unix() > time.Now().AddDate(0, 0, 7).Unix() {
			continue
		}

		var matchImageURL string
		savedMatch, err := svc.matchesSvc.GetMatchByID(ctx, match.ID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}

			matchImageURL, err = svc.generateMatchBanner(ctx, match.Esport, match.Name)
			if err != nil {
				return err
			}

			err = svc.matchesSvc.CreateMatch(ctx, matches.MatchCMD{
				ID:        match.ID,
				Date:      match.Date,
				Name:      match.Name,
				TeamA:     match.TeamA,
				TeamB:     match.TeamB,
				BannerURL: matchImageURL,
			})
			if err != nil {
				return err
			}
		}
		matchImageURL = savedMatch.BannerURL

		err = svc.CreateMarket(ctx, MarketCMD{
			GameName:   match.Name,
			CandidateA: match.TeamA,
			CandidateB: match.TeamB,
			ExpiryTime: match.Date.Add(match.Duration),
			LockTime:   match.Date,
			ImageURL:   matchImageURL,
		})
		if err != nil {
			return err
		}
	}
	return nil
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
