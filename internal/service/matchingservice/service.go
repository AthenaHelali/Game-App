package matchingservice

import (
	"context"
	entity2 "game-app/internal/entity"
	param2 "game-app/internal/param"
	"game-app/internal/pkg/richerror"
	"time"
)

type Repo interface {
	AddToWaitingList(userID uint, category entity2.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity2.Category) ([]entity2.WaitingMember, error)
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

type Service struct {
	repo   Repo
	config Config
}

func New(config Config, repo Repo) Service {
	return Service{config: config, repo: repo}
}

func (s Service) AddToWaitingList(req param2.AddToWaitingListRequest) (param2.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitingList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return param2.AddToWaitingListResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return param2.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}

func (s Service) MatchWaitingUsers(_ param2.MatchWaitingUsersRequest) (param2.MatchWaitingUsersResponse, error) {

	return param2.MatchWaitingUsersResponse{}, nil
}
