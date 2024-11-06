package markets

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/openai/openai-go"

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
}

func NewService(
	transactor *bind.TransactOpts,
	sankoPredicts *sankoPredicts.SankoPredicts,
	matchesSvc matches.Service,
	openaiClient *openai.Client,
) *service {
	return &service{
		transactor:    transactor,
		sankoPredicts: sankoPredicts,
		matchesSvc:    matchesSvc,
		openaiClient:  openaiClient,
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

			log.Printf("duplicated game: %+v\n\n", market)
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
		log.Printf("transaction sent: %s\n\n", tx.Hash().Hex())
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

			matchImageURL, err = svc.generateMatchBanner(ctx, match.Name)
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
			ExpiryTime: match.Date.Add(time.Hour),
			LockTime:   match.Date,
			ImageURL:   matchImageURL,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) generateMatchBanner(ctx context.Context, matchName string) (string, error) {
	res, err := svc.openaiClient.Images.Generate(ctx, openai.ImageGenerateParams{
		Prompt: openai.String(fmt.Sprintf("draw %s banner featuring both team logos, a dynamic stadium background, and a bold design", matchName)),
		Model:  openai.String("dall-e-3"),
	})
	if err != nil {
		return "", err
	}
	return res.Data[0].URL, nil
}
