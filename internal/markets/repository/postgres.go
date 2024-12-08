package repository

import (
	"context"
	"database/sql"
	"math/big"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"

	"github.com/jeronimobarea/sp_predictions_bot/internal/markets"
)

const (
	tableName   = "markets"
	tableFields = "id, chain, creator, a_seed_amount, b_seed_amount, creation_date"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) PersistMarket(ctx context.Context, market markets.Market) error {
	q, args, err := sq.Insert(tableName).
		Columns("id", "chain", "creator", "a_seed_amount", "b_seed_amount").
		Values(
			market.ID,
			market.Chain,
			market.Creator,
			market.ASeedAmount,
			market.BSeedAmount,
		).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetMarketByID(ctx context.Context, marketID big.Int) (*markets.Market, error) {
	q, args, err := sq.Select(tableFields).
		From(tableName).
		Where(sq.Eq{
			"id": marketID,
		}).ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, q, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var market markets.Market
	err = row.Scan(
		&market.ID,
		&market.Chain,
		&market.Creator,

		&market.ASeedAmount,
		&market.BSeedAmount,

		&market.CreationDate,
		&market.UpdateDate,
	)
	if err != nil {
		return nil, err
	}
	return &market, nil
}

func (r *repository) UpdateMarket(ctx context.Context, market markets.Market) error {
	q, args, err := sq.Update(tableName).
		Set("a_seed_amount", market.ASeedAmount).
		Set("b_seed_amount", market.BSeedAmount).
		Set("update_date", time.Now()).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
