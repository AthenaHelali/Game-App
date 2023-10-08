package matchingservice

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
	"game-app/pkg/timestamp"
	"time"
)

type Repo interface {
	AddToWaitingList(userID uint, category entity.Category) error
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

func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitingList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return param.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}

func (s Service) MatchWaitingUsers(req param.MatchWaitingUsersRequest) (param.MatchWaitingUsersResponse, error) {
	fmt.Println("matching...", timestamp.Now())
	return param.MatchWaitingUsersResponse{}, nil
}
