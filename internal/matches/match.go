package matches

import (
	"math/big"
	"time"
)

type (
	MatchCMD struct {
		ID        string
		Date      time.Time
		Name      string
		TeamA     string
		TeamB     string
		BannerURL string
	}

	Match struct {
		ID        string
		Date      time.Time
		Name      string
		TeamA     string
		TeamB     string
		BannerURL string

		SankoTx    *big.Int
		ArbitrumTx *big.Int
		EthereumTx *big.Int
		BaseTx     *big.Int
		ApeTx      *big.Int

		CreationDate time.Time
		UpdateDate   *time.Time
	}
)

func (m *Match) SetTx(chain string, tx big.Int) {}
