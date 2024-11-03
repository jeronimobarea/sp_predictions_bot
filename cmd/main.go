package main

import (
	"context"
	"crypto/ecdsa"
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

var (
	rpcURI           string
	walletPrivateKey *ecdsa.PrivateKey
	scAddress        common.Address
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	rpcURI = os.Getenv("CHAIN_RPC_URI")
	if rpcURI == "" {
		panic("missing rpc uri")
	}

	pk := os.Getenv("SP_BOT_WALLET_PK")
	if pk == "" {
		panic("missing wallet private key")
	}
	walletPrivateKey, err = crypto.HexToECDSA(pk)
	if err != nil {
		panic(err)
	}

	address := os.Getenv("SP_SC_ADDRESS")
	if address == "" {
		panic("missing smart contract address")
	}
	scAddress = common.HexToAddress(address)
}

func main() {
	ctx := context.Background()

	var espnSvc espn.Service
	{
		espnClient := espnClient.NewClient(http.DefaultClient)
		espnSvc = espn.NewService(espnClient)
	}

	var marketsSvc markets.Service
	{

		ethClient, err := ethclient.DialContext(ctx, rpcURI)
		if err != nil {
			panic(err)
		}

		chainID, err := ethClient.ChainID(ctx)
		if err != nil {
			panic(err)
		}
		log.Printf("chain id: %d", chainID)

		spTransactor, err := sankoPredicts.NewSankoPredicts(scAddress, ethClient)
		if err != nil {
			panic(err)
		}

		transactor, err := bind.NewKeyedTransactorWithChainID(walletPrivateKey, chainID)
		if err != nil {
			panic(err)
		}

		marketsSvc = markets.NewService(transactor, spTransactor)
	}

	approvedMarkets, err := marketsSvc.GetMarkets(ctx, markets.MarketStatusApproved)
	if err != nil {
		panic(err)
	}
	log.Printf("approved markets: %+v", approvedMarkets)

	pendingMarkets, err := marketsSvc.GetMarkets(ctx, markets.MarketStatusPending)
	if err != nil {
		panic(err)
	}
	log.Printf("pending markets: %+v", pendingMarkets)

	rejectedMarkets, err := marketsSvc.GetMarkets(ctx, markets.MarketStatusRejected)
	if err != nil {
		panic(err)
	}
	log.Printf("rejected markets: %+v", rejectedMarkets)

	nbaMatches, err := espnSvc.GetNBAScoreboard(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("nba matches: %+v", nbaMatches)

	processMatches(ctx, marketsSvc, nbaMatches)
}

func processMatches(ctx context.Context, marketsSvc markets.Service, matches []espn.ESPNMatch) {
	for _, match := range matches {
		if match.Date.Unix() <= time.Now().Unix() {
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
