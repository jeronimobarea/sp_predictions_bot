package markets

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	sankoPredicts "github.com/jeronimobarea/sp_predictions_bot/pkg/sankopredicts"
)

type Service interface {
	CreateMarket(ctx context.Context, cmd MarketCMD) error
	GetMarkets(ctx context.Context, status int64) ([]sankoPredicts.SankoPredictsGame, error)
}

type service struct {
	transactor    *bind.TransactOpts
	sankoPredicts *sankoPredicts.SankoPredicts
}

func NewService(
	transactor *bind.TransactOpts,
	sankoPredicts *sankoPredicts.SankoPredicts,
) *service {
	return &service{
		transactor:    transactor,
		sankoPredicts: sankoPredicts,
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
			market.GameName == cmd.GameName {

			log.Printf("duplicated game: %+v", market)
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
		log.Printf("Transaction sent: %s", tx.Hash().Hex())
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
