package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"

	"github.com/jeronimobarea/sp_predictions_bot/internal/matches"
)

const (
	tableName   = "matches"
	tableFields = "id, date, name, team_a, team_b, banner_url, sanko_tx, arbitrum_tx, ethereum_tx, base_tx, ape_tx, creation_date"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) PersistMatch(ctx context.Context, match matches.Match) error {
	q, args, err := sq.Insert(tableName).
		Columns(tableFields).
		Values(
			match.ID,
			match.Date,
			match.Name,
			match.TeamA,
			match.TeamB,
			match.BannerURL,
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

func (r *repository) GetMatchByID(ctx context.Context, matchID string) (*matches.Match, error) {
	q, args, err := sq.Select(tableFields).
		From(tableName).
		Where(sq.Eq{
			"id": matchID,
		}).ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, q, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var match matches.Match
	if err := row.Scan(
		&match.ID,
		&match.Date,
		&match.Name,
		&match.TeamA,
		&match.TeamB,
		&match.BannerURL,

		&match.SankoTx,
		&match.ArbitrumTx,
		&match.EthereumTx,
		&match.BaseTx,
		&match.ApeTx,

		&match.CreationDate,
		&match.UpdateDate,
	); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return &match, nil
}

func (r *repository) UpdateMatch(ctx context.Context, match matches.Match) error {
	q, args, err := sq.Update(tableName).
		Set("sanko_tx", match.SankoTx).
		Set("arbitrum_tx", match.ArbitrumTx).
		Set("ethereum_tx", match.EthereumTx).
		Set("base_tx", match.BaseTx).
		Set("ape_tx", match.ApeTx).
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
