package markets

import (
	"context"
	"math/big"
)

type Repository interface {
	PersistMarket(ctx context.Context, market Market) error
	UpdateMarket(ctx context.Context, market Market) error
	GetMarketByID(ctx context.Context, marketID big.Int) (*Market, error)
}
