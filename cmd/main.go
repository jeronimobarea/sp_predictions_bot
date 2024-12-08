package main

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"go.uber.org/zap"

	"github.com/jeronimobarea/sp_predictions_bot/internal/espn"
	espnClient "github.com/jeronimobarea/sp_predictions_bot/internal/espn/client"
	"github.com/jeronimobarea/sp_predictions_bot/internal/markets"
	marketsRepository "github.com/jeronimobarea/sp_predictions_bot/internal/markets/repository"
	"github.com/jeronimobarea/sp_predictions_bot/internal/matches"
	matchesRepository "github.com/jeronimobarea/sp_predictions_bot/internal/matches/repository"
	sankoPredicts "github.com/jeronimobarea/sp_predictions_bot/pkg/sankopredicts"
)

type Chain struct {
	Name                     string
	RPC                      string
	PredictionsMarketAddress common.Address
}

const (
	SankoChain    = "SANKO"
	ArbitrumChain = "ARBITRUM"
	EthereumChain = "ETHEREUM"
	BaseChain     = "BASE"
	ApeChain      = "APE"
)

var (
	chains           map[string]Chain
	walletPrivateKey *ecdsa.PrivateKey

	// validChains = []string{SankoChain, ArbitrumChain, BaseChain, EthereumChain}
	validChains = []string{SankoChain}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	chains = make(map[string]Chain)
	for _, chain := range validChains {
		rpcURI := os.Getenv(chain + "_RPC_URI")
		if rpcURI == "" {
			panic(fmt.Errorf("missing %s rpc uri", chain))
		}

		scAddress := os.Getenv(chain + "_PREDICTIONS_ADDRESS")
		if scAddress == "" {
			panic(fmt.Errorf("missing %s smart contract address", chain))
		}
		predictionsMarketAddress := common.HexToAddress(scAddress)

		chains[chain] = Chain{
			Name:                     chain,
			RPC:                      rpcURI,
			PredictionsMarketAddress: predictionsMarketAddress,
		}
	}

	pk := os.Getenv("SP_BOT_WALLET_PK")
	if pk == "" {
		panic("missing wallet private key")
	}
	walletPrivateKey, err = crypto.HexToECDSA(pk)
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	var logger *zap.SugaredLogger
	{
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		logger = l.Sugar()
	}

	var espnSvc espn.Service
	{
		espnClient := espnClient.NewClient(http.DefaultClient)
		espnSvc = espn.NewService(espnClient, logger)
	}

	var matchesSvc matches.Service
	{
		dbUrl := buildDBUrl()
		db, err := sql.Open("postgres", dbUrl)
		if err != nil {
			logger.Fatal(err)
		}

		repo := matchesRepository.NewRepository(db)

		openaiClient := openai.NewClient()

		matchesSvc = matches.NewService(repo, openaiClient)
	}

	var marketsRepo markets.Repository
	{
		dbUrl := buildDBUrl()
		db, err := sql.Open("postgres", dbUrl)
		if err != nil {
			logger.Fatal(err)
		}

		marketsRepo = marketsRepository.NewRepository(db)
	}

	matches, err := espnSvc.GetAllScoreboards(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	processedMatches, err := matchesSvc.ProcessMatches(ctx, matches)
	if err != nil {
		logger.Fatal(err)
	}

	for _, chain := range chains {
		marketsSvc := setupMarketSvc(ctx, chain.Name, logger, matchesSvc, marketsRepo, chain.RPC, chain.PredictionsMarketAddress, walletPrivateKey)

		for _, match := range processedMatches {
			market, err := marketsSvc.CreateMarket(ctx, markets.MarketCMD{
				GameName:   match.Name,
				CandidateA: match.TeamA,
				CandidateB: match.TeamB,
				ExpiryTime: match.Date.Add(time.Hour),
				LockTime:   match.Date,
				ImageURL:   match.BannerURL,
			})
			if err != nil {
				logger.Error(err)
			}

			err = marketsSvc.SeedMarket(ctx, match, market.Id, big.NewInt(1), markets.MarketASeedingOption)
			if err != nil {
				logger.Error(err)
			}

			err = marketsSvc.SeedMarket(ctx, match, market.Id, big.NewInt(1), markets.MarketBSeedingOption)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func setupMarketSvc(
	ctx context.Context,
	chain string,
	logger *zap.SugaredLogger,
	matchesSvc matches.Service,
	repo markets.Repository,
	rpcURI string,
	scAddress common.Address,
	pk *ecdsa.PrivateKey,
) markets.Service {
	chainClient, err := ethclient.DialContext(ctx, rpcURI)
	if err != nil {
		logger.Fatal(err)
	}

	chainID, err := chainClient.ChainID(ctx)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infow("chain connection made",
		"RPC", rpcURI,
		"ChainID", chainID,
	)

	spTransactor, err := sankoPredicts.NewSankoPredicts(scAddress, chainClient)
	if err != nil {
		logger.Fatal(err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		logger.Fatal(err)
	}

	return markets.NewService(transactor, spTransactor, matchesSvc, repo, chain, logger)
}

func buildDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@postgres/%s?sslmode=disable",
		"postgres",
		"postgres",
		"sp_predictions_bot",
	)
}
