package main

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go"

	"github.com/jeronimobarea/sp_predictions_bot/internal/espn"
	espnClient "github.com/jeronimobarea/sp_predictions_bot/internal/espn/client"
	"github.com/jeronimobarea/sp_predictions_bot/internal/markets"
	"github.com/jeronimobarea/sp_predictions_bot/internal/matches"
	matchesRepository "github.com/jeronimobarea/sp_predictions_bot/internal/matches/repository"
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

	var matchesSvc matches.Service
	{
		dbUrl := buildDBUrl()
		db, err := sql.Open("postgres", dbUrl)
		if err != nil {
			panic(err)
		}

		repo := matchesRepository.NewRepository(db)
		matchesSvc = matches.NewService(repo)
	}

	openaiClient := openai.NewClient()

	for _, chain := range chains {
		svc := setupMarketSvc(ctx, matchesSvc, openaiClient, chain.RPC, chain.PredictionsMarketAddress, walletPrivateKey)

		matches, err := espnSvc.GetAllScoreboards(ctx)
		if err != nil {
			panic(err)
		}

		err = svc.ProcessMatches(ctx, matches)
		if err != nil {
			panic(err)
		}
	}
}

func setupMarketSvc(ctx context.Context, matchesSvc matches.Service, openaiClient *openai.Client, rpcURI string, scAddress common.Address, pk *ecdsa.PrivateKey) markets.Service {
	chainClient, err := ethclient.DialContext(ctx, rpcURI)
	if err != nil {
		panic(err)
	}

	chainID, err := chainClient.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("chain RPC: %s <-> chain id: %d\n\n", rpcURI, chainID)

	spTransactor, err := sankoPredicts.NewSankoPredicts(scAddress, chainClient)
	if err != nil {
		panic(err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		panic(err)
	}

	return markets.NewService(transactor, spTransactor, matchesSvc, openaiClient)
}

func buildDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@postgres/%s?sslmode=disable",
		"postgres",
		"postgres",
		"sp_predictions_bot",
	)
}
