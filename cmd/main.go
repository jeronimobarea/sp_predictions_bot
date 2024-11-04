package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/jeronimobarea/sp_predictions_bot/internal/espn"
	espnClient "github.com/jeronimobarea/sp_predictions_bot/internal/espn/client"
	"github.com/jeronimobarea/sp_predictions_bot/internal/markets"
	sankoPredicts "github.com/jeronimobarea/sp_predictions_bot/pkg/sankopredicts"
)

type Chain struct {
	RPC                      string
	PredictionsMarketAddress common.Address
}

const (
	SankoChain    = "SANKO"
	ArbitrumChain = "ARBITRUM"
	BaseChain     = "BASE"
	EthereumChain = "ETHEREUM"
)

var (
	chains           map[string]Chain
	marketSvcs       map[string]markets.Service
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
	marketSvcs = make(map[string]markets.Service)

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

	var espnSvc espn.Service
	{
		espnClient := espnClient.NewClient(http.DefaultClient)
		espnSvc = espn.NewService(espnClient)
	}

	nbaMatches, err := espnSvc.GetNBAScoreboard(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("nba matches: %+v", nbaMatches)

	mlbMatches, err := espnSvc.GetMLBScoreboard(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("mlb matches: %+v", mlbMatches)

	nflMatches, err := espnSvc.GetNFLScoreboard(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("nfl matches: %+v", nflMatches)

	ufcMatches, err := espnSvc.GetUFCScoreboard(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("ufc matches: %+v", ufcMatches)

	for chainName, chain := range chains {
		svc := setupMarketSvc(ctx, chain.RPC, chain.PredictionsMarketAddress, walletPrivateKey)
		marketSvcs[chainName] = svc

		approvedMarkets, err := svc.GetMarkets(ctx, markets.MarketStatusApproved)
		if err != nil {
			panic(err)
		}
		log.Printf("approved markets: %+v", approvedMarkets)

		pendingMarkets, err := svc.GetMarkets(ctx, markets.MarketStatusPending)
		if err != nil {
			panic(err)
		}
		log.Printf("pending markets: %+v", pendingMarkets)

		rejectedMarkets, err := svc.GetMarkets(ctx, markets.MarketStatusRejected)
		if err != nil {
			panic(err)
		}
		log.Printf("rejected markets: %+v", rejectedMarkets)

		// processMatches(ctx, svc, nbaMatches)
		// processMatches(ctx, svc, mlbMatches)
		// processMatches(ctx, svc, nflMatches)
		// processMatches(ctx, svc, ufcMatches)
	}
}

func setupMarketSvc(ctx context.Context, rpcURI string, scAddress common.Address, pk *ecdsa.PrivateKey) markets.Service {
	chainClient, err := ethclient.DialContext(ctx, rpcURI)
	if err != nil {
		panic(err)
	}

	chainID, err := chainClient.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("chain RPC: %s <-> chain id: %d", rpcURI, chainID)

	spTransactor, err := sankoPredicts.NewSankoPredicts(scAddress, chainClient)
	if err != nil {
		panic(err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		panic(err)
	}

	return markets.NewService(transactor, spTransactor)
}

func processMatches(ctx context.Context, marketsSvc markets.Service, matches []espn.ESPNMatch) {
	for _, match := range matches {
		if match.Date.Unix() <= time.Now().Unix() ||
			match.Date.Unix() > time.Now().AddDate(0, 0, 7).Unix() {
			continue
		}

		err := marketsSvc.CreateMarket(ctx, markets.MarketCMD{
			GameName:   match.Name,
			CandidateA: match.TeamA,
			CandidateB: match.TeamB,
			ExpiryTime: match.Date.Add(time.Hour),
			LockTime:   match.Date,
			ImageURL:   match.TeamALogo,
		})
		if err != nil {
			panic(err)
		}
	}
}
