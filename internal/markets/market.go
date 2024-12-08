package markets

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type (
	MarketCMD struct {
		GameName   string
		CandidateA string
		CandidateB string
		ExpiryTime time.Time
		LockTime   time.Time
		ImageURL   string
	}

	Market struct {
		ID      big.Int
		Chain   string
		Creator common.Address

		ASeedAmount *big.Int
		BSeedAmount *big.Int

		CreationDate time.Time
		UpdateDate   *time.Time
	}
)

func (m Market) isSeeded(option uint8) bool {
	return option == MarketASeedingOption && m.ASeedAmount != nil || option == MarketBSeedingOption && m.BSeedAmount != nil
}

func (m *Market) setAmount(option uint8, amount *big.Int) {
	if option == MarketASeedingOption {
		m.ASeedAmount = amount
	}
	if option == MarketBSeedingOption {
		m.BSeedAmount = amount
	}
}
