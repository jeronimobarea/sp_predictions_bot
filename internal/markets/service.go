package markets

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"

	"github.com/jeronimobarea/sp_predictions_bot/internal/matches"
	sankoPredicts "github.com/jeronimobarea/sp_predictions_bot/pkg/sankopredicts"
)

type Service interface {
	CreateMarket(ctx context.Context, cmd MarketCMD) (*sankoPredicts.SankoPredictsGame, error)
	GetMarketsFromChainData(ctx context.Context, status int64) ([]sankoPredicts.SankoPredictsGame, error)
	SeedMarket(ctx context.Context, match matches.Match, marketID, amount *big.Int, option uint8) error
}

type service struct {
	transactor    *bind.TransactOpts
	sankoPredicts *sankoPredicts.SankoPredicts
	matchesSvc    matches.Service
	repo          Repository
	chainName     string
	logger        *zap.SugaredLogger
}

func NewService(
	transactor *bind.TransactOpts,
	sankoPredicts *sankoPredicts.SankoPredicts,
	matchesSvc matches.Service,
	repo Repository,
	chainName string,
	logger *zap.SugaredLogger,
) *service {
	return &service{
		transactor:    transactor,
		sankoPredicts: sankoPredicts,
		matchesSvc:    matchesSvc,
		repo:          repo,
		chainName:     chainName,
		logger:        logger,
	}
}

const (
	ElectionGameType = "Election"

	MarketStatusPending  int64 = 1
	MarketStatusRejected int64 = 2
	MarketStatusApproved int64 = 3

	MarketASeedingOption uint8 = 1
	MarketBSeedingOption uint8 = 2
)

func (svc *service) CreateMarket(ctx context.Context, cmd MarketCMD) (*sankoPredicts.SankoPredictsGame, error) {
	m, err := svc.marketHasBeenCreated(ctx, cmd.CandidateA, cmd.CandidateB, cmd.GameName)
	if err != nil {
		return nil, err
	}
	if m != nil {
		return m, nil
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
		return nil, err
	}
	svc.logger.Infow("market created on chain",
		"TransactionID", tx.Hash().Hex(),
		"Market", cmd,
	)

	var market Market
	err = json.Unmarshal(tx.Data(), &market)
	if err != nil {
		return nil, err
	}
	market.Chain = svc.chainName

	err = svc.repo.PersistMarket(ctx, market)
	if err != nil {
		return nil, err
	}
	svc.logger.Infow("market saved",
		"Market", cmd,
	)
	return &sankoPredicts.SankoPredictsGame{Id: &market.ID}, nil
}

func (svc *service) GetMarketsFromChainData(ctx context.Context, status int64) ([]sankoPredicts.SankoPredictsGame, error) {
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

func (svc *service) SeedMarket(
	ctx context.Context,
	match matches.Match,
	marketID,
	amount *big.Int,
	option uint8,
) error {
	m, err := svc.repo.GetMarketByID(ctx, *marketID)
	if err != nil {
		return err
	}
	if m.isSeeded(option) {
		svc.logger.Infow("market already seeded",
			"saved market data", m,
		)
		return nil
	}

	tx, err := svc.sankoPredicts.Deposit(
		svc.transactor,
		marketID,
		option,
		amount,
	)
	if err != nil {
		return err
	}

	svc.logger.Infow("market seeded",
		"market", m,
		"transaxtion", tx,
	)

	m.setAmount(option, amount)

	err = svc.repo.UpdateMarket(ctx, *m)
	if err != nil {
		return err
	}

	match.SetTx(svc.chainName, m.ID)

	return svc.matchesSvc.UpdateMatch(ctx, match)
}

func (svc *service) marketHasBeenCreated(ctx context.Context, candidateA, candidateB, gameName string) (*sankoPredicts.SankoPredictsGame, error) {
	pendingMarkets, err := svc.GetMarketsFromChainData(ctx, MarketStatusPending)
	if err != nil {
		return nil, err
	}

	rejectedMarkets, err := svc.GetMarketsFromChainData(ctx, MarketStatusRejected)
	if err != nil {
		return nil, err
	}

	approvedMarkets, err := svc.GetMarketsFromChainData(ctx, MarketStatusApproved)
	if err != nil {
		return nil, err
	}

	var markets []sankoPredicts.SankoPredictsGame
	markets = append(markets, pendingMarkets...)
	markets = append(markets, rejectedMarkets...)
	markets = append(markets, approvedMarkets...)

	for _, market := range markets {
		if market.CandidateA == candidateA &&
			market.CandidateB == candidateB &&
			market.GameName == gameName &&
			market.Creator == svc.transactor.From {

			svc.logger.Warnw("duplicated market",
				"market", market,
			)
			return &market, nil
		}
	}
	return nil, nil
}
