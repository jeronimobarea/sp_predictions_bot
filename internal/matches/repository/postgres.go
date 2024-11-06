package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/jeronimobarea/sp_predictions_bot/internal/matches"
)

const (
	tableName   = "matches"
	tableFields = "id, date, name, team_a, team_b, banner_url"
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
	q := `INSERT INTO ` + tableName + `(id, date, name, team_a, team_b, banner_url)` +
		`VALUES ($1, $2, $3, $4, $5, $6)`

	args := []interface{}{
		match.ID,
		match.Date,
		match.Name,
		match.TeamA,
		match.TeamB,
		match.BannerURL,
	}

	_, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetMatchByID(ctx context.Context, matchID string) (*matches.Match, error) {
	q := `SELECT ` + tableFields + ` FROM ` + tableName

	row := r.db.QueryRowContext(ctx, q)
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
	); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return &match, nil
}
