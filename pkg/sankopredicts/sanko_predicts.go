// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sanko_predicts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SankoPredictsGame is an auto generated low-level Go binding around an user-defined struct.
type SankoPredictsGame struct {
	Id           *big.Int
	GameName     string
	Banner       string
	GameType     uint8
	CandidateA   string
	CandidateB   string
	InitialPrice *big.Int
	FinalPrice   *big.Int
	ExpiryTime   *big.Int
	LockTime     *big.Int
	TotalPoolA   *big.Int
	TotalPoolB   *big.Int
	Resolved     bool
	WinnerPool   *big.Int
	FeePercent   *big.Int
	Market       string
	Creator      common.Address
	Status       *big.Int
}

// SankoPredictsGameDeposit is an auto generated low-level Go binding around an user-defined struct.
type SankoPredictsGameDeposit struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Pool      uint8
}

// SankoPredictsLeaderboardEntry is an auto generated low-level Go binding around an user-defined struct.
type SankoPredictsLeaderboardEntry struct {
	Player        common.Address
	TotalDeposits *big.Int
	TotalWinnings *big.Int
}

// SankoPredictsMetaData contains all meta data concerning the SankoPredicts contract.
var SankoPredictsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_spToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"option\",\"type\":\"uint8\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeePercent\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newfeePercentage\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gameName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSankoPredicts.GameType\",\"name\":\"gameType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"candidateA\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"candidateB\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"banner\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"GameCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"winner\",\"type\":\"uint8\"}],\"name\":\"GameResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_gameName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_market\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_banner\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"createCoinGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_gameName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_market\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_candidateA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_candidateB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_banner\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"createElectionGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFeePercentile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"option\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pool\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositsA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositsB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"banner\",\"type\":\"string\"},{\"internalType\":\"enumSankoPredicts.GameType\",\"name\":\"gameType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"candidateA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"candidateB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoolB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"winnerPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"market\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimableOrPendingGames\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"claimableGameIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingGameIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getDepositsForGame\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pool\",\"type\":\"uint8\"}],\"internalType\":\"structSankoPredicts.GameDeposit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"banner\",\"type\":\"string\"},{\"internalType\":\"enumSankoPredicts.GameType\",\"name\":\"gameType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"candidateA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"candidateB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoolB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"winnerPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"market\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structSankoPredicts.Game\",\"name\":\"game\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pool\",\"type\":\"uint8\"}],\"internalType\":\"structSankoPredicts.GameDeposit[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"includeResolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getGames\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gameName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"banner\",\"type\":\"string\"},{\"internalType\":\"enumSankoPredicts.GameType\",\"name\":\"gameType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"candidateA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"candidateB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoolA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoolB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"winnerPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"market\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structSankoPredicts.Game[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLeaderboard\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWinnings\",\"type\":\"uint256\"}],\"internalType\":\"structSankoPredicts.LeaderboardEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"leaderboard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWinnings\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"winner\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_finalPrice\",\"type\":\"uint256\"}],\"name\":\"resolveGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercentile\",\"type\":\"uint256\"}],\"name\":\"updateCreatorFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newfeePercentage\",\"type\":\"uint256\"}],\"name\":\"updateFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newStatus\",\"type\":\"uint256\"}],\"name\":\"updateGameStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SankoPredictsABI is the input ABI used to generate the binding from.
// Deprecated: Use SankoPredictsMetaData.ABI instead.
var SankoPredictsABI = SankoPredictsMetaData.ABI

// SankoPredicts is an auto generated Go binding around an Ethereum contract.
type SankoPredicts struct {
	SankoPredictsCaller     // Read-only binding to the contract
	SankoPredictsTransactor // Write-only binding to the contract
	SankoPredictsFilterer   // Log filterer for contract events
}

// SankoPredictsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SankoPredictsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SankoPredictsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SankoPredictsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SankoPredictsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SankoPredictsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SankoPredictsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SankoPredictsSession struct {
	Contract     *SankoPredicts    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SankoPredictsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SankoPredictsCallerSession struct {
	Contract *SankoPredictsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SankoPredictsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SankoPredictsTransactorSession struct {
	Contract     *SankoPredictsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SankoPredictsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SankoPredictsRaw struct {
	Contract *SankoPredicts // Generic contract binding to access the raw methods on
}

// SankoPredictsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SankoPredictsCallerRaw struct {
	Contract *SankoPredictsCaller // Generic read-only contract binding to access the raw methods on
}

// SankoPredictsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SankoPredictsTransactorRaw struct {
	Contract *SankoPredictsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSankoPredicts creates a new instance of SankoPredicts, bound to a specific deployed contract.
func NewSankoPredicts(address common.Address, backend bind.ContractBackend) (*SankoPredicts, error) {
	contract, err := bindSankoPredicts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SankoPredicts{SankoPredictsCaller: SankoPredictsCaller{contract: contract}, SankoPredictsTransactor: SankoPredictsTransactor{contract: contract}, SankoPredictsFilterer: SankoPredictsFilterer{contract: contract}}, nil
}

// NewSankoPredictsCaller creates a new read-only instance of SankoPredicts, bound to a specific deployed contract.
func NewSankoPredictsCaller(address common.Address, caller bind.ContractCaller) (*SankoPredictsCaller, error) {
	contract, err := bindSankoPredicts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SankoPredictsCaller{contract: contract}, nil
}

// NewSankoPredictsTransactor creates a new write-only instance of SankoPredicts, bound to a specific deployed contract.
func NewSankoPredictsTransactor(address common.Address, transactor bind.ContractTransactor) (*SankoPredictsTransactor, error) {
	contract, err := bindSankoPredicts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SankoPredictsTransactor{contract: contract}, nil
}

// NewSankoPredictsFilterer creates a new log filterer instance of SankoPredicts, bound to a specific deployed contract.
func NewSankoPredictsFilterer(address common.Address, filterer bind.ContractFilterer) (*SankoPredictsFilterer, error) {
	contract, err := bindSankoPredicts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SankoPredictsFilterer{contract: contract}, nil
}

// bindSankoPredicts binds a generic wrapper to an already deployed contract.
func bindSankoPredicts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SankoPredictsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SankoPredicts *SankoPredictsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SankoPredicts.Contract.SankoPredictsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SankoPredicts *SankoPredictsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SankoPredicts.Contract.SankoPredictsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SankoPredicts *SankoPredictsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SankoPredicts.Contract.SankoPredictsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SankoPredicts *SankoPredictsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SankoPredicts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SankoPredicts *SankoPredictsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SankoPredicts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SankoPredicts *SankoPredictsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SankoPredicts.Contract.contract.Transact(opts, method, params...)
}

// CreatorFeePercentile is a free data retrieval call binding the contract method 0xe3eb8a35.
//
// Solidity: function creatorFeePercentile() view returns(uint256)
func (_SankoPredicts *SankoPredictsCaller) CreatorFeePercentile(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "creatorFeePercentile")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorFeePercentile is a free data retrieval call binding the contract method 0xe3eb8a35.
//
// Solidity: function creatorFeePercentile() view returns(uint256)
func (_SankoPredicts *SankoPredictsSession) CreatorFeePercentile() (*big.Int, error) {
	return _SankoPredicts.Contract.CreatorFeePercentile(&_SankoPredicts.CallOpts)
}

// CreatorFeePercentile is a free data retrieval call binding the contract method 0xe3eb8a35.
//
// Solidity: function creatorFeePercentile() view returns(uint256)
func (_SankoPredicts *SankoPredictsCallerSession) CreatorFeePercentile() (*big.Int, error) {
	return _SankoPredicts.Contract.CreatorFeePercentile(&_SankoPredicts.CallOpts)
}

// DepositHistory is a free data retrieval call binding the contract method 0xeb57e649.
//
// Solidity: function depositHistory(uint256 , uint256 ) view returns(address user, uint256 amount, uint256 timestamp, uint8 pool)
func (_SankoPredicts *SankoPredictsCaller) DepositHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Pool      uint8
}, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "depositHistory", arg0, arg1)

	outstruct := new(struct {
		User      common.Address
		Amount    *big.Int
		Timestamp *big.Int
		Pool      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Pool = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// DepositHistory is a free data retrieval call binding the contract method 0xeb57e649.
//
// Solidity: function depositHistory(uint256 , uint256 ) view returns(address user, uint256 amount, uint256 timestamp, uint8 pool)
func (_SankoPredicts *SankoPredictsSession) DepositHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Pool      uint8
}, error) {
	return _SankoPredicts.Contract.DepositHistory(&_SankoPredicts.CallOpts, arg0, arg1)
}

// DepositHistory is a free data retrieval call binding the contract method 0xeb57e649.
//
// Solidity: function depositHistory(uint256 , uint256 ) view returns(address user, uint256 amount, uint256 timestamp, uint8 pool)
func (_SankoPredicts *SankoPredictsCallerSession) DepositHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Pool      uint8
}, error) {
	return _SankoPredicts.Contract.DepositHistory(&_SankoPredicts.CallOpts, arg0, arg1)
}

// DepositsA is a free data retrieval call binding the contract method 0x47b00525.
//
// Solidity: function depositsA(uint256 , address ) view returns(uint256)
func (_SankoPredicts *SankoPredictsCaller) DepositsA(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "depositsA", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsA is a free data retrieval call binding the contract method 0x47b00525.
//
// Solidity: function depositsA(uint256 , address ) view returns(uint256)
func (_SankoPredicts *SankoPredictsSession) DepositsA(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SankoPredicts.Contract.DepositsA(&_SankoPredicts.CallOpts, arg0, arg1)
}

// DepositsA is a free data retrieval call binding the contract method 0x47b00525.
//
// Solidity: function depositsA(uint256 , address ) view returns(uint256)
func (_SankoPredicts *SankoPredictsCallerSession) DepositsA(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SankoPredicts.Contract.DepositsA(&_SankoPredicts.CallOpts, arg0, arg1)
}

// DepositsB is a free data retrieval call binding the contract method 0xd1406cea.
//
// Solidity: function depositsB(uint256 , address ) view returns(uint256)
func (_SankoPredicts *SankoPredictsCaller) DepositsB(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "depositsB", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsB is a free data retrieval call binding the contract method 0xd1406cea.
//
// Solidity: function depositsB(uint256 , address ) view returns(uint256)
func (_SankoPredicts *SankoPredictsSession) DepositsB(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SankoPredicts.Contract.DepositsB(&_SankoPredicts.CallOpts, arg0, arg1)
}

// DepositsB is a free data retrieval call binding the contract method 0xd1406cea.
//
// Solidity: function depositsB(uint256 , address ) view returns(uint256)
func (_SankoPredicts *SankoPredictsCallerSession) DepositsB(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SankoPredicts.Contract.DepositsB(&_SankoPredicts.CallOpts, arg0, arg1)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_SankoPredicts *SankoPredictsCaller) FeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_SankoPredicts *SankoPredictsSession) FeePercentage() (*big.Int, error) {
	return _SankoPredicts.Contract.FeePercentage(&_SankoPredicts.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_SankoPredicts *SankoPredictsCallerSession) FeePercentage() (*big.Int, error) {
	return _SankoPredicts.Contract.FeePercentage(&_SankoPredicts.CallOpts)
}

// GameCount is a free data retrieval call binding the contract method 0x4d1975b4.
//
// Solidity: function gameCount() view returns(uint256)
func (_SankoPredicts *SankoPredictsCaller) GameCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "gameCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameCount is a free data retrieval call binding the contract method 0x4d1975b4.
//
// Solidity: function gameCount() view returns(uint256)
func (_SankoPredicts *SankoPredictsSession) GameCount() (*big.Int, error) {
	return _SankoPredicts.Contract.GameCount(&_SankoPredicts.CallOpts)
}

// GameCount is a free data retrieval call binding the contract method 0x4d1975b4.
//
// Solidity: function gameCount() view returns(uint256)
func (_SankoPredicts *SankoPredictsCallerSession) GameCount() (*big.Int, error) {
	return _SankoPredicts.Contract.GameCount(&_SankoPredicts.CallOpts)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(uint256 id, string gameName, string banner, uint8 gameType, string candidateA, string candidateB, uint256 initialPrice, uint256 finalPrice, uint256 expiryTime, uint256 lockTime, uint256 totalPoolA, uint256 totalPoolB, bool resolved, uint256 winnerPool, uint256 feePercent, string market, address creator, uint256 status)
func (_SankoPredicts *SankoPredictsCaller) Games(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	GameName     string
	Banner       string
	GameType     uint8
	CandidateA   string
	CandidateB   string
	InitialPrice *big.Int
	FinalPrice   *big.Int
	ExpiryTime   *big.Int
	LockTime     *big.Int
	TotalPoolA   *big.Int
	TotalPoolB   *big.Int
	Resolved     bool
	WinnerPool   *big.Int
	FeePercent   *big.Int
	Market       string
	Creator      common.Address
	Status       *big.Int
}, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		Id           *big.Int
		GameName     string
		Banner       string
		GameType     uint8
		CandidateA   string
		CandidateB   string
		InitialPrice *big.Int
		FinalPrice   *big.Int
		ExpiryTime   *big.Int
		LockTime     *big.Int
		TotalPoolA   *big.Int
		TotalPoolB   *big.Int
		Resolved     bool
		WinnerPool   *big.Int
		FeePercent   *big.Int
		Market       string
		Creator      common.Address
		Status       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GameName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Banner = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.GameType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.CandidateA = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CandidateB = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.InitialPrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.FinalPrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ExpiryTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LockTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TotalPoolA = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TotalPoolB = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[12], new(bool)).(*bool)
	outstruct.WinnerPool = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)
	outstruct.FeePercent = *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)
	outstruct.Market = *abi.ConvertType(out[15], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[16], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[17], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(uint256 id, string gameName, string banner, uint8 gameType, string candidateA, string candidateB, uint256 initialPrice, uint256 finalPrice, uint256 expiryTime, uint256 lockTime, uint256 totalPoolA, uint256 totalPoolB, bool resolved, uint256 winnerPool, uint256 feePercent, string market, address creator, uint256 status)
func (_SankoPredicts *SankoPredictsSession) Games(arg0 *big.Int) (struct {
	Id           *big.Int
	GameName     string
	Banner       string
	GameType     uint8
	CandidateA   string
	CandidateB   string
	InitialPrice *big.Int
	FinalPrice   *big.Int
	ExpiryTime   *big.Int
	LockTime     *big.Int
	TotalPoolA   *big.Int
	TotalPoolB   *big.Int
	Resolved     bool
	WinnerPool   *big.Int
	FeePercent   *big.Int
	Market       string
	Creator      common.Address
	Status       *big.Int
}, error) {
	return _SankoPredicts.Contract.Games(&_SankoPredicts.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(uint256 id, string gameName, string banner, uint8 gameType, string candidateA, string candidateB, uint256 initialPrice, uint256 finalPrice, uint256 expiryTime, uint256 lockTime, uint256 totalPoolA, uint256 totalPoolB, bool resolved, uint256 winnerPool, uint256 feePercent, string market, address creator, uint256 status)
func (_SankoPredicts *SankoPredictsCallerSession) Games(arg0 *big.Int) (struct {
	Id           *big.Int
	GameName     string
	Banner       string
	GameType     uint8
	CandidateA   string
	CandidateB   string
	InitialPrice *big.Int
	FinalPrice   *big.Int
	ExpiryTime   *big.Int
	LockTime     *big.Int
	TotalPoolA   *big.Int
	TotalPoolB   *big.Int
	Resolved     bool
	WinnerPool   *big.Int
	FeePercent   *big.Int
	Market       string
	Creator      common.Address
	Status       *big.Int
}, error) {
	return _SankoPredicts.Contract.Games(&_SankoPredicts.CallOpts, arg0)
}

// GetClaimableOrPendingGames is a free data retrieval call binding the contract method 0xb8c6f500.
//
// Solidity: function getClaimableOrPendingGames(address user) view returns(uint256[] claimableGameIds, uint256[] pendingGameIds)
func (_SankoPredicts *SankoPredictsCaller) GetClaimableOrPendingGames(opts *bind.CallOpts, user common.Address) (struct {
	ClaimableGameIds []*big.Int
	PendingGameIds   []*big.Int
}, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "getClaimableOrPendingGames", user)

	outstruct := new(struct {
		ClaimableGameIds []*big.Int
		PendingGameIds   []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClaimableGameIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.PendingGameIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetClaimableOrPendingGames is a free data retrieval call binding the contract method 0xb8c6f500.
//
// Solidity: function getClaimableOrPendingGames(address user) view returns(uint256[] claimableGameIds, uint256[] pendingGameIds)
func (_SankoPredicts *SankoPredictsSession) GetClaimableOrPendingGames(user common.Address) (struct {
	ClaimableGameIds []*big.Int
	PendingGameIds   []*big.Int
}, error) {
	return _SankoPredicts.Contract.GetClaimableOrPendingGames(&_SankoPredicts.CallOpts, user)
}

// GetClaimableOrPendingGames is a free data retrieval call binding the contract method 0xb8c6f500.
//
// Solidity: function getClaimableOrPendingGames(address user) view returns(uint256[] claimableGameIds, uint256[] pendingGameIds)
func (_SankoPredicts *SankoPredictsCallerSession) GetClaimableOrPendingGames(user common.Address) (struct {
	ClaimableGameIds []*big.Int
	PendingGameIds   []*big.Int
}, error) {
	return _SankoPredicts.Contract.GetClaimableOrPendingGames(&_SankoPredicts.CallOpts, user)
}

// GetDepositsForGame is a free data retrieval call binding the contract method 0xc722c558.
//
// Solidity: function getDepositsForGame(uint256 gameId) view returns((address,uint256,uint256,uint8)[])
func (_SankoPredicts *SankoPredictsCaller) GetDepositsForGame(opts *bind.CallOpts, gameId *big.Int) ([]SankoPredictsGameDeposit, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "getDepositsForGame", gameId)

	if err != nil {
		return *new([]SankoPredictsGameDeposit), err
	}

	out0 := *abi.ConvertType(out[0], new([]SankoPredictsGameDeposit)).(*[]SankoPredictsGameDeposit)

	return out0, err

}

// GetDepositsForGame is a free data retrieval call binding the contract method 0xc722c558.
//
// Solidity: function getDepositsForGame(uint256 gameId) view returns((address,uint256,uint256,uint8)[])
func (_SankoPredicts *SankoPredictsSession) GetDepositsForGame(gameId *big.Int) ([]SankoPredictsGameDeposit, error) {
	return _SankoPredicts.Contract.GetDepositsForGame(&_SankoPredicts.CallOpts, gameId)
}

// GetDepositsForGame is a free data retrieval call binding the contract method 0xc722c558.
//
// Solidity: function getDepositsForGame(uint256 gameId) view returns((address,uint256,uint256,uint8)[])
func (_SankoPredicts *SankoPredictsCallerSession) GetDepositsForGame(gameId *big.Int) ([]SankoPredictsGameDeposit, error) {
	return _SankoPredicts.Contract.GetDepositsForGame(&_SankoPredicts.CallOpts, gameId)
}

// GetGameDetails is a free data retrieval call binding the contract method 0x1b31abda.
//
// Solidity: function getGameDetails(uint256 gameId) view returns((uint256,string,string,uint8,string,string,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,address,uint256) game, (address,uint256,uint256,uint8)[] deposits)
func (_SankoPredicts *SankoPredictsCaller) GetGameDetails(opts *bind.CallOpts, gameId *big.Int) (struct {
	Game     SankoPredictsGame
	Deposits []SankoPredictsGameDeposit
}, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "getGameDetails", gameId)

	outstruct := new(struct {
		Game     SankoPredictsGame
		Deposits []SankoPredictsGameDeposit
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Game = *abi.ConvertType(out[0], new(SankoPredictsGame)).(*SankoPredictsGame)
	outstruct.Deposits = *abi.ConvertType(out[1], new([]SankoPredictsGameDeposit)).(*[]SankoPredictsGameDeposit)

	return *outstruct, err

}

// GetGameDetails is a free data retrieval call binding the contract method 0x1b31abda.
//
// Solidity: function getGameDetails(uint256 gameId) view returns((uint256,string,string,uint8,string,string,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,address,uint256) game, (address,uint256,uint256,uint8)[] deposits)
func (_SankoPredicts *SankoPredictsSession) GetGameDetails(gameId *big.Int) (struct {
	Game     SankoPredictsGame
	Deposits []SankoPredictsGameDeposit
}, error) {
	return _SankoPredicts.Contract.GetGameDetails(&_SankoPredicts.CallOpts, gameId)
}

// GetGameDetails is a free data retrieval call binding the contract method 0x1b31abda.
//
// Solidity: function getGameDetails(uint256 gameId) view returns((uint256,string,string,uint8,string,string,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,address,uint256) game, (address,uint256,uint256,uint8)[] deposits)
func (_SankoPredicts *SankoPredictsCallerSession) GetGameDetails(gameId *big.Int) (struct {
	Game     SankoPredictsGame
	Deposits []SankoPredictsGameDeposit
}, error) {
	return _SankoPredicts.Contract.GetGameDetails(&_SankoPredicts.CallOpts, gameId)
}

// GetGames is a free data retrieval call binding the contract method 0x33dd563f.
//
// Solidity: function getGames(bool includeResolved, uint256 status, address creator) view returns((uint256,string,string,uint8,string,string,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,address,uint256)[])
func (_SankoPredicts *SankoPredictsCaller) GetGames(opts *bind.CallOpts, includeResolved bool, status *big.Int, creator common.Address) ([]SankoPredictsGame, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "getGames", includeResolved, status, creator)

	if err != nil {
		return *new([]SankoPredictsGame), err
	}

	out0 := *abi.ConvertType(out[0], new([]SankoPredictsGame)).(*[]SankoPredictsGame)

	return out0, err

}

// GetGames is a free data retrieval call binding the contract method 0x33dd563f.
//
// Solidity: function getGames(bool includeResolved, uint256 status, address creator) view returns((uint256,string,string,uint8,string,string,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,address,uint256)[])
func (_SankoPredicts *SankoPredictsSession) GetGames(includeResolved bool, status *big.Int, creator common.Address) ([]SankoPredictsGame, error) {
	return _SankoPredicts.Contract.GetGames(&_SankoPredicts.CallOpts, includeResolved, status, creator)
}

// GetGames is a free data retrieval call binding the contract method 0x33dd563f.
//
// Solidity: function getGames(bool includeResolved, uint256 status, address creator) view returns((uint256,string,string,uint8,string,string,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,address,uint256)[])
func (_SankoPredicts *SankoPredictsCallerSession) GetGames(includeResolved bool, status *big.Int, creator common.Address) ([]SankoPredictsGame, error) {
	return _SankoPredicts.Contract.GetGames(&_SankoPredicts.CallOpts, includeResolved, status, creator)
}

// GetLeaderboard is a free data retrieval call binding the contract method 0x6d763a6e.
//
// Solidity: function getLeaderboard() view returns((address,uint256,uint256)[])
func (_SankoPredicts *SankoPredictsCaller) GetLeaderboard(opts *bind.CallOpts) ([]SankoPredictsLeaderboardEntry, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "getLeaderboard")

	if err != nil {
		return *new([]SankoPredictsLeaderboardEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]SankoPredictsLeaderboardEntry)).(*[]SankoPredictsLeaderboardEntry)

	return out0, err

}

// GetLeaderboard is a free data retrieval call binding the contract method 0x6d763a6e.
//
// Solidity: function getLeaderboard() view returns((address,uint256,uint256)[])
func (_SankoPredicts *SankoPredictsSession) GetLeaderboard() ([]SankoPredictsLeaderboardEntry, error) {
	return _SankoPredicts.Contract.GetLeaderboard(&_SankoPredicts.CallOpts)
}

// GetLeaderboard is a free data retrieval call binding the contract method 0x6d763a6e.
//
// Solidity: function getLeaderboard() view returns((address,uint256,uint256)[])
func (_SankoPredicts *SankoPredictsCallerSession) GetLeaderboard() ([]SankoPredictsLeaderboardEntry, error) {
	return _SankoPredicts.Contract.GetLeaderboard(&_SankoPredicts.CallOpts)
}

// Leaderboard is a free data retrieval call binding the contract method 0xd1d33d20.
//
// Solidity: function leaderboard(address ) view returns(address player, uint256 totalDeposits, uint256 totalWinnings)
func (_SankoPredicts *SankoPredictsCaller) Leaderboard(opts *bind.CallOpts, arg0 common.Address) (struct {
	Player        common.Address
	TotalDeposits *big.Int
	TotalWinnings *big.Int
}, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "leaderboard", arg0)

	outstruct := new(struct {
		Player        common.Address
		TotalDeposits *big.Int
		TotalWinnings *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Player = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalDeposits = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalWinnings = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Leaderboard is a free data retrieval call binding the contract method 0xd1d33d20.
//
// Solidity: function leaderboard(address ) view returns(address player, uint256 totalDeposits, uint256 totalWinnings)
func (_SankoPredicts *SankoPredictsSession) Leaderboard(arg0 common.Address) (struct {
	Player        common.Address
	TotalDeposits *big.Int
	TotalWinnings *big.Int
}, error) {
	return _SankoPredicts.Contract.Leaderboard(&_SankoPredicts.CallOpts, arg0)
}

// Leaderboard is a free data retrieval call binding the contract method 0xd1d33d20.
//
// Solidity: function leaderboard(address ) view returns(address player, uint256 totalDeposits, uint256 totalWinnings)
func (_SankoPredicts *SankoPredictsCallerSession) Leaderboard(arg0 common.Address) (struct {
	Player        common.Address
	TotalDeposits *big.Int
	TotalWinnings *big.Int
}, error) {
	return _SankoPredicts.Contract.Leaderboard(&_SankoPredicts.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SankoPredicts *SankoPredictsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SankoPredicts *SankoPredictsSession) Owner() (common.Address, error) {
	return _SankoPredicts.Contract.Owner(&_SankoPredicts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SankoPredicts *SankoPredictsCallerSession) Owner() (common.Address, error) {
	return _SankoPredicts.Contract.Owner(&_SankoPredicts.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_SankoPredicts *SankoPredictsCaller) Players(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "players", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_SankoPredicts *SankoPredictsSession) Players(arg0 *big.Int) (common.Address, error) {
	return _SankoPredicts.Contract.Players(&_SankoPredicts.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_SankoPredicts *SankoPredictsCallerSession) Players(arg0 *big.Int) (common.Address, error) {
	return _SankoPredicts.Contract.Players(&_SankoPredicts.CallOpts, arg0)
}

// SpToken is a free data retrieval call binding the contract method 0x8e148776.
//
// Solidity: function spToken() view returns(address)
func (_SankoPredicts *SankoPredictsCaller) SpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SankoPredicts.contract.Call(opts, &out, "spToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpToken is a free data retrieval call binding the contract method 0x8e148776.
//
// Solidity: function spToken() view returns(address)
func (_SankoPredicts *SankoPredictsSession) SpToken() (common.Address, error) {
	return _SankoPredicts.Contract.SpToken(&_SankoPredicts.CallOpts)
}

// SpToken is a free data retrieval call binding the contract method 0x8e148776.
//
// Solidity: function spToken() view returns(address)
func (_SankoPredicts *SankoPredictsCallerSession) SpToken() (common.Address, error) {
	return _SankoPredicts.Contract.SpToken(&_SankoPredicts.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 gameId) returns()
func (_SankoPredicts *SankoPredictsTransactor) Claim(opts *bind.TransactOpts, gameId *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "claim", gameId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 gameId) returns()
func (_SankoPredicts *SankoPredictsSession) Claim(gameId *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.Claim(&_SankoPredicts.TransactOpts, gameId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 gameId) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) Claim(gameId *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.Claim(&_SankoPredicts.TransactOpts, gameId)
}

// CreateCoinGame is a paid mutator transaction binding the contract method 0xc9a95819.
//
// Solidity: function createCoinGame(string _gameName, uint256 _expiryTime, uint256 _lockTime, string _market, uint256 _initialPrice, string _banner, address _creator) returns()
func (_SankoPredicts *SankoPredictsTransactor) CreateCoinGame(opts *bind.TransactOpts, _gameName string, _expiryTime *big.Int, _lockTime *big.Int, _market string, _initialPrice *big.Int, _banner string, _creator common.Address) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "createCoinGame", _gameName, _expiryTime, _lockTime, _market, _initialPrice, _banner, _creator)
}

// CreateCoinGame is a paid mutator transaction binding the contract method 0xc9a95819.
//
// Solidity: function createCoinGame(string _gameName, uint256 _expiryTime, uint256 _lockTime, string _market, uint256 _initialPrice, string _banner, address _creator) returns()
func (_SankoPredicts *SankoPredictsSession) CreateCoinGame(_gameName string, _expiryTime *big.Int, _lockTime *big.Int, _market string, _initialPrice *big.Int, _banner string, _creator common.Address) (*types.Transaction, error) {
	return _SankoPredicts.Contract.CreateCoinGame(&_SankoPredicts.TransactOpts, _gameName, _expiryTime, _lockTime, _market, _initialPrice, _banner, _creator)
}

// CreateCoinGame is a paid mutator transaction binding the contract method 0xc9a95819.
//
// Solidity: function createCoinGame(string _gameName, uint256 _expiryTime, uint256 _lockTime, string _market, uint256 _initialPrice, string _banner, address _creator) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) CreateCoinGame(_gameName string, _expiryTime *big.Int, _lockTime *big.Int, _market string, _initialPrice *big.Int, _banner string, _creator common.Address) (*types.Transaction, error) {
	return _SankoPredicts.Contract.CreateCoinGame(&_SankoPredicts.TransactOpts, _gameName, _expiryTime, _lockTime, _market, _initialPrice, _banner, _creator)
}

// CreateElectionGame is a paid mutator transaction binding the contract method 0x771558d9.
//
// Solidity: function createElectionGame(string _gameName, uint256 _expiryTime, uint256 _lockTime, string _market, string _candidateA, string _candidateB, string _banner, address _creator) returns()
func (_SankoPredicts *SankoPredictsTransactor) CreateElectionGame(opts *bind.TransactOpts, _gameName string, _expiryTime *big.Int, _lockTime *big.Int, _market string, _candidateA string, _candidateB string, _banner string, _creator common.Address) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "createElectionGame", _gameName, _expiryTime, _lockTime, _market, _candidateA, _candidateB, _banner, _creator)
}

// CreateElectionGame is a paid mutator transaction binding the contract method 0x771558d9.
//
// Solidity: function createElectionGame(string _gameName, uint256 _expiryTime, uint256 _lockTime, string _market, string _candidateA, string _candidateB, string _banner, address _creator) returns()
func (_SankoPredicts *SankoPredictsSession) CreateElectionGame(_gameName string, _expiryTime *big.Int, _lockTime *big.Int, _market string, _candidateA string, _candidateB string, _banner string, _creator common.Address) (*types.Transaction, error) {
	return _SankoPredicts.Contract.CreateElectionGame(&_SankoPredicts.TransactOpts, _gameName, _expiryTime, _lockTime, _market, _candidateA, _candidateB, _banner, _creator)
}

// CreateElectionGame is a paid mutator transaction binding the contract method 0x771558d9.
//
// Solidity: function createElectionGame(string _gameName, uint256 _expiryTime, uint256 _lockTime, string _market, string _candidateA, string _candidateB, string _banner, address _creator) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) CreateElectionGame(_gameName string, _expiryTime *big.Int, _lockTime *big.Int, _market string, _candidateA string, _candidateB string, _banner string, _creator common.Address) (*types.Transaction, error) {
	return _SankoPredicts.Contract.CreateElectionGame(&_SankoPredicts.TransactOpts, _gameName, _expiryTime, _lockTime, _market, _candidateA, _candidateB, _banner, _creator)
}

// Deposit is a paid mutator transaction binding the contract method 0x60798cab.
//
// Solidity: function deposit(uint256 gameId, uint8 option, uint256 amount) returns()
func (_SankoPredicts *SankoPredictsTransactor) Deposit(opts *bind.TransactOpts, gameId *big.Int, option uint8, amount *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "deposit", gameId, option, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x60798cab.
//
// Solidity: function deposit(uint256 gameId, uint8 option, uint256 amount) returns()
func (_SankoPredicts *SankoPredictsSession) Deposit(gameId *big.Int, option uint8, amount *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.Deposit(&_SankoPredicts.TransactOpts, gameId, option, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x60798cab.
//
// Solidity: function deposit(uint256 gameId, uint8 option, uint256 amount) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) Deposit(gameId *big.Int, option uint8, amount *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.Deposit(&_SankoPredicts.TransactOpts, gameId, option, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SankoPredicts *SankoPredictsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SankoPredicts *SankoPredictsSession) RenounceOwnership() (*types.Transaction, error) {
	return _SankoPredicts.Contract.RenounceOwnership(&_SankoPredicts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SankoPredicts *SankoPredictsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SankoPredicts.Contract.RenounceOwnership(&_SankoPredicts.TransactOpts)
}

// ResolveGame is a paid mutator transaction binding the contract method 0xc6e565d3.
//
// Solidity: function resolveGame(uint256 gameId, uint8 winner, uint256 _finalPrice) returns()
func (_SankoPredicts *SankoPredictsTransactor) ResolveGame(opts *bind.TransactOpts, gameId *big.Int, winner uint8, _finalPrice *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "resolveGame", gameId, winner, _finalPrice)
}

// ResolveGame is a paid mutator transaction binding the contract method 0xc6e565d3.
//
// Solidity: function resolveGame(uint256 gameId, uint8 winner, uint256 _finalPrice) returns()
func (_SankoPredicts *SankoPredictsSession) ResolveGame(gameId *big.Int, winner uint8, _finalPrice *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.ResolveGame(&_SankoPredicts.TransactOpts, gameId, winner, _finalPrice)
}

// ResolveGame is a paid mutator transaction binding the contract method 0xc6e565d3.
//
// Solidity: function resolveGame(uint256 gameId, uint8 winner, uint256 _finalPrice) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) ResolveGame(gameId *big.Int, winner uint8, _finalPrice *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.ResolveGame(&_SankoPredicts.TransactOpts, gameId, winner, _finalPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SankoPredicts *SankoPredictsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SankoPredicts *SankoPredictsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SankoPredicts.Contract.TransferOwnership(&_SankoPredicts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SankoPredicts.Contract.TransferOwnership(&_SankoPredicts.TransactOpts, newOwner)
}

// UpdateCreatorFeePercentage is a paid mutator transaction binding the contract method 0x8873a498.
//
// Solidity: function updateCreatorFeePercentage(uint256 _newPercentile) returns()
func (_SankoPredicts *SankoPredictsTransactor) UpdateCreatorFeePercentage(opts *bind.TransactOpts, _newPercentile *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "updateCreatorFeePercentage", _newPercentile)
}

// UpdateCreatorFeePercentage is a paid mutator transaction binding the contract method 0x8873a498.
//
// Solidity: function updateCreatorFeePercentage(uint256 _newPercentile) returns()
func (_SankoPredicts *SankoPredictsSession) UpdateCreatorFeePercentage(_newPercentile *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.UpdateCreatorFeePercentage(&_SankoPredicts.TransactOpts, _newPercentile)
}

// UpdateCreatorFeePercentage is a paid mutator transaction binding the contract method 0x8873a498.
//
// Solidity: function updateCreatorFeePercentage(uint256 _newPercentile) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) UpdateCreatorFeePercentage(_newPercentile *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.UpdateCreatorFeePercentage(&_SankoPredicts.TransactOpts, _newPercentile)
}

// UpdateFeePercentage is a paid mutator transaction binding the contract method 0x6cad3fb0.
//
// Solidity: function updateFeePercentage(uint256 newfeePercentage) returns()
func (_SankoPredicts *SankoPredictsTransactor) UpdateFeePercentage(opts *bind.TransactOpts, newfeePercentage *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "updateFeePercentage", newfeePercentage)
}

// UpdateFeePercentage is a paid mutator transaction binding the contract method 0x6cad3fb0.
//
// Solidity: function updateFeePercentage(uint256 newfeePercentage) returns()
func (_SankoPredicts *SankoPredictsSession) UpdateFeePercentage(newfeePercentage *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.UpdateFeePercentage(&_SankoPredicts.TransactOpts, newfeePercentage)
}

// UpdateFeePercentage is a paid mutator transaction binding the contract method 0x6cad3fb0.
//
// Solidity: function updateFeePercentage(uint256 newfeePercentage) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) UpdateFeePercentage(newfeePercentage *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.UpdateFeePercentage(&_SankoPredicts.TransactOpts, newfeePercentage)
}

// UpdateGameStatus is a paid mutator transaction binding the contract method 0xdb6ea62a.
//
// Solidity: function updateGameStatus(uint256 gameId, uint256 newStatus) returns()
func (_SankoPredicts *SankoPredictsTransactor) UpdateGameStatus(opts *bind.TransactOpts, gameId *big.Int, newStatus *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "updateGameStatus", gameId, newStatus)
}

// UpdateGameStatus is a paid mutator transaction binding the contract method 0xdb6ea62a.
//
// Solidity: function updateGameStatus(uint256 gameId, uint256 newStatus) returns()
func (_SankoPredicts *SankoPredictsSession) UpdateGameStatus(gameId *big.Int, newStatus *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.UpdateGameStatus(&_SankoPredicts.TransactOpts, gameId, newStatus)
}

// UpdateGameStatus is a paid mutator transaction binding the contract method 0xdb6ea62a.
//
// Solidity: function updateGameStatus(uint256 gameId, uint256 newStatus) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) UpdateGameStatus(gameId *big.Int, newStatus *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.UpdateGameStatus(&_SankoPredicts.TransactOpts, gameId, newStatus)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_SankoPredicts *SankoPredictsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_SankoPredicts *SankoPredictsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.Withdraw(&_SankoPredicts.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_SankoPredicts *SankoPredictsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SankoPredicts.Contract.Withdraw(&_SankoPredicts.TransactOpts, amount)
}

// SankoPredictsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SankoPredicts contract.
type SankoPredictsDepositIterator struct {
	Event *SankoPredictsDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsDeposit represents a Deposit event raised by the SankoPredicts contract.
type SankoPredictsDeposit struct {
	User   common.Address
	GameId *big.Int
	Amount *big.Int
	Option uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xa7db47d395930224de3f54139208b4a958fe3747c2ad1dee5a6624643a6def93.
//
// Solidity: event Deposit(address indexed user, uint256 gameId, uint256 amount, uint8 option)
func (_SankoPredicts *SankoPredictsFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*SankoPredictsDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &SankoPredictsDepositIterator{contract: _SankoPredicts.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xa7db47d395930224de3f54139208b4a958fe3747c2ad1dee5a6624643a6def93.
//
// Solidity: event Deposit(address indexed user, uint256 gameId, uint256 amount, uint8 option)
func (_SankoPredicts *SankoPredictsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SankoPredictsDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsDeposit)
				if err := _SankoPredicts.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xa7db47d395930224de3f54139208b4a958fe3747c2ad1dee5a6624643a6def93.
//
// Solidity: event Deposit(address indexed user, uint256 gameId, uint256 amount, uint8 option)
func (_SankoPredicts *SankoPredictsFilterer) ParseDeposit(log types.Log) (*SankoPredictsDeposit, error) {
	event := new(SankoPredictsDeposit)
	if err := _SankoPredicts.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SankoPredictsFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the SankoPredicts contract.
type SankoPredictsFeeUpdatedIterator struct {
	Event *SankoPredictsFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsFeeUpdated represents a FeeUpdated event raised by the SankoPredicts contract.
type SankoPredictsFeeUpdated struct {
	GameId        *big.Int
	NewFeePercent *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 gameId, uint256 newFeePercent)
func (_SankoPredicts *SankoPredictsFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*SankoPredictsFeeUpdatedIterator, error) {

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SankoPredictsFeeUpdatedIterator{contract: _SankoPredicts.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 gameId, uint256 newFeePercent)
func (_SankoPredicts *SankoPredictsFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *SankoPredictsFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsFeeUpdated)
				if err := _SankoPredicts.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeUpdated is a log parse operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 gameId, uint256 newFeePercent)
func (_SankoPredicts *SankoPredictsFilterer) ParseFeeUpdated(log types.Log) (*SankoPredictsFeeUpdated, error) {
	event := new(SankoPredictsFeeUpdated)
	if err := _SankoPredicts.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SankoPredictsFeeUpdated0Iterator is returned from FilterFeeUpdated0 and is used to iterate over the raw logs and unpacked data for FeeUpdated0 events raised by the SankoPredicts contract.
type SankoPredictsFeeUpdated0Iterator struct {
	Event *SankoPredictsFeeUpdated0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsFeeUpdated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsFeeUpdated0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsFeeUpdated0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsFeeUpdated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsFeeUpdated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsFeeUpdated0 represents a FeeUpdated0 event raised by the SankoPredicts contract.
type SankoPredictsFeeUpdated0 struct {
	NewfeePercentage *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated0 is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newfeePercentage)
func (_SankoPredicts *SankoPredictsFilterer) FilterFeeUpdated0(opts *bind.FilterOpts) (*SankoPredictsFeeUpdated0Iterator, error) {

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "FeeUpdated0")
	if err != nil {
		return nil, err
	}
	return &SankoPredictsFeeUpdated0Iterator{contract: _SankoPredicts.contract, event: "FeeUpdated0", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated0 is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newfeePercentage)
func (_SankoPredicts *SankoPredictsFilterer) WatchFeeUpdated0(opts *bind.WatchOpts, sink chan<- *SankoPredictsFeeUpdated0) (event.Subscription, error) {

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "FeeUpdated0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsFeeUpdated0)
				if err := _SankoPredicts.contract.UnpackLog(event, "FeeUpdated0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeUpdated0 is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newfeePercentage)
func (_SankoPredicts *SankoPredictsFilterer) ParseFeeUpdated0(log types.Log) (*SankoPredictsFeeUpdated0, error) {
	event := new(SankoPredictsFeeUpdated0)
	if err := _SankoPredicts.contract.UnpackLog(event, "FeeUpdated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SankoPredictsGameCreatedIterator is returned from FilterGameCreated and is used to iterate over the raw logs and unpacked data for GameCreated events raised by the SankoPredicts contract.
type SankoPredictsGameCreatedIterator struct {
	Event *SankoPredictsGameCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsGameCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsGameCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsGameCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsGameCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsGameCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsGameCreated represents a GameCreated event raised by the SankoPredicts contract.
type SankoPredictsGameCreated struct {
	GameId       *big.Int
	GameName     string
	GameType     uint8
	CandidateA   string
	CandidateB   string
	InitialPrice *big.Int
	ExpiryTime   *big.Int
	LockTime     *big.Int
	FeePercent   *big.Int
	Banner       string
	Creator      common.Address
	Status       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGameCreated is a free log retrieval operation binding the contract event 0xdc8a2367f4eb37400edfd399e2a1c519acc1e98fb473dc6cee99fa064d4d7d21.
//
// Solidity: event GameCreated(uint256 gameId, string gameName, uint8 gameType, string candidateA, string candidateB, uint256 initialPrice, uint256 expiryTime, uint256 lockTime, uint256 feePercent, string banner, address creator, uint256 status)
func (_SankoPredicts *SankoPredictsFilterer) FilterGameCreated(opts *bind.FilterOpts) (*SankoPredictsGameCreatedIterator, error) {

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "GameCreated")
	if err != nil {
		return nil, err
	}
	return &SankoPredictsGameCreatedIterator{contract: _SankoPredicts.contract, event: "GameCreated", logs: logs, sub: sub}, nil
}

// WatchGameCreated is a free log subscription operation binding the contract event 0xdc8a2367f4eb37400edfd399e2a1c519acc1e98fb473dc6cee99fa064d4d7d21.
//
// Solidity: event GameCreated(uint256 gameId, string gameName, uint8 gameType, string candidateA, string candidateB, uint256 initialPrice, uint256 expiryTime, uint256 lockTime, uint256 feePercent, string banner, address creator, uint256 status)
func (_SankoPredicts *SankoPredictsFilterer) WatchGameCreated(opts *bind.WatchOpts, sink chan<- *SankoPredictsGameCreated) (event.Subscription, error) {

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "GameCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsGameCreated)
				if err := _SankoPredicts.contract.UnpackLog(event, "GameCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGameCreated is a log parse operation binding the contract event 0xdc8a2367f4eb37400edfd399e2a1c519acc1e98fb473dc6cee99fa064d4d7d21.
//
// Solidity: event GameCreated(uint256 gameId, string gameName, uint8 gameType, string candidateA, string candidateB, uint256 initialPrice, uint256 expiryTime, uint256 lockTime, uint256 feePercent, string banner, address creator, uint256 status)
func (_SankoPredicts *SankoPredictsFilterer) ParseGameCreated(log types.Log) (*SankoPredictsGameCreated, error) {
	event := new(SankoPredictsGameCreated)
	if err := _SankoPredicts.contract.UnpackLog(event, "GameCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SankoPredictsGameResolvedIterator is returned from FilterGameResolved and is used to iterate over the raw logs and unpacked data for GameResolved events raised by the SankoPredicts contract.
type SankoPredictsGameResolvedIterator struct {
	Event *SankoPredictsGameResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsGameResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsGameResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsGameResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsGameResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsGameResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsGameResolved represents a GameResolved event raised by the SankoPredicts contract.
type SankoPredictsGameResolved struct {
	GameId *big.Int
	Winner uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGameResolved is a free log retrieval operation binding the contract event 0xc476b09c91681b1e03ac948089793c804a27569df87c784dec68c1e1bf244a4d.
//
// Solidity: event GameResolved(uint256 gameId, uint8 winner)
func (_SankoPredicts *SankoPredictsFilterer) FilterGameResolved(opts *bind.FilterOpts) (*SankoPredictsGameResolvedIterator, error) {

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "GameResolved")
	if err != nil {
		return nil, err
	}
	return &SankoPredictsGameResolvedIterator{contract: _SankoPredicts.contract, event: "GameResolved", logs: logs, sub: sub}, nil
}

// WatchGameResolved is a free log subscription operation binding the contract event 0xc476b09c91681b1e03ac948089793c804a27569df87c784dec68c1e1bf244a4d.
//
// Solidity: event GameResolved(uint256 gameId, uint8 winner)
func (_SankoPredicts *SankoPredictsFilterer) WatchGameResolved(opts *bind.WatchOpts, sink chan<- *SankoPredictsGameResolved) (event.Subscription, error) {

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "GameResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsGameResolved)
				if err := _SankoPredicts.contract.UnpackLog(event, "GameResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGameResolved is a log parse operation binding the contract event 0xc476b09c91681b1e03ac948089793c804a27569df87c784dec68c1e1bf244a4d.
//
// Solidity: event GameResolved(uint256 gameId, uint8 winner)
func (_SankoPredicts *SankoPredictsFilterer) ParseGameResolved(log types.Log) (*SankoPredictsGameResolved, error) {
	event := new(SankoPredictsGameResolved)
	if err := _SankoPredicts.contract.UnpackLog(event, "GameResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SankoPredictsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SankoPredicts contract.
type SankoPredictsOwnershipTransferredIterator struct {
	Event *SankoPredictsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsOwnershipTransferred represents a OwnershipTransferred event raised by the SankoPredicts contract.
type SankoPredictsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SankoPredicts *SankoPredictsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SankoPredictsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SankoPredictsOwnershipTransferredIterator{contract: _SankoPredicts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SankoPredicts *SankoPredictsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SankoPredictsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsOwnershipTransferred)
				if err := _SankoPredicts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SankoPredicts *SankoPredictsFilterer) ParseOwnershipTransferred(log types.Log) (*SankoPredictsOwnershipTransferred, error) {
	event := new(SankoPredictsOwnershipTransferred)
	if err := _SankoPredicts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SankoPredictsPayoutIterator is returned from FilterPayout and is used to iterate over the raw logs and unpacked data for Payout events raised by the SankoPredicts contract.
type SankoPredictsPayoutIterator struct {
	Event *SankoPredictsPayout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SankoPredictsPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SankoPredictsPayout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SankoPredictsPayout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SankoPredictsPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SankoPredictsPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SankoPredictsPayout represents a Payout event raised by the SankoPredicts contract.
type SankoPredictsPayout struct {
	User   common.Address
	GameId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPayout is a free log retrieval operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address indexed user, uint256 gameId, uint256 amount)
func (_SankoPredicts *SankoPredictsFilterer) FilterPayout(opts *bind.FilterOpts, user []common.Address) (*SankoPredictsPayoutIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SankoPredicts.contract.FilterLogs(opts, "Payout", userRule)
	if err != nil {
		return nil, err
	}
	return &SankoPredictsPayoutIterator{contract: _SankoPredicts.contract, event: "Payout", logs: logs, sub: sub}, nil
}

// WatchPayout is a free log subscription operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address indexed user, uint256 gameId, uint256 amount)
func (_SankoPredicts *SankoPredictsFilterer) WatchPayout(opts *bind.WatchOpts, sink chan<- *SankoPredictsPayout, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SankoPredicts.contract.WatchLogs(opts, "Payout", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SankoPredictsPayout)
				if err := _SankoPredicts.contract.UnpackLog(event, "Payout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayout is a log parse operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address indexed user, uint256 gameId, uint256 amount)
func (_SankoPredicts *SankoPredictsFilterer) ParsePayout(log types.Log) (*SankoPredictsPayout, error) {
	event := new(SankoPredictsPayout)
	if err := _SankoPredicts.contract.UnpackLog(event, "Payout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
