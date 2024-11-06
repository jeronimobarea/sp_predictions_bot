package matches

import "context"

type Service interface {
	CreateMatch(ctx context.Context, cmd MatchCMD) error
	GetMatchByID(ctx context.Context, matchID string) (*Match, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (svc *service) CreateMatch(ctx context.Context, cmd MatchCMD) error {
	return svc.repo.PersistMatch(ctx, Match{
		ID:        cmd.ID,
		Date:      cmd.Date,
		Name:      cmd.Name,
		TeamA:     cmd.TeamA,
		TeamB:     cmd.TeamB,
		BannerURL: cmd.BannerURL,
	})
}

func (svc *service) GetMatchByID(ctx context.Context, matchID string) (*Match, error) {
	return svc.repo.GetMatchByID(ctx, matchID)
}
