package matches

import "context"

type Repository interface {
	PersistMatch(ctx context.Context, match Match) error
	UpdateMatch(ctx context.Context, match Match) error
	GetMatchByID(ctx context.Context, matchID string) (*Match, error)
}
