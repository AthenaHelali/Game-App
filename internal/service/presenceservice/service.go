package presenceservice

import (
	"context"
	"fmt"
	"game-app/internal/param"
	"game-app/internal/pkg/richerror"
	"time"
)

type Config struct {
	PresenceExpirationTime time.Duration `koanf:"expiration_time"`
	PresencePrefix         string        `koanf:"prefix"`
}

type Repo interface {
	Upsert(ctx context.Context, key string, timeStamp int64, expTime time.Duration) error
}

type Service struct {
	config Config
	repo   Repo
}

func New(config Config, repo Repo) Service {
	return Service{
		config: config,
		repo:   repo,
	}
}

func (s Service) Upsert(ctx context.Context, req param.UpsertPresenceRequest) (param.UpsertPresenceResponse, error) {
	const op = "presenceservice.Upsert"

	err := s.repo.Upsert(ctx,
		fmt.Sprintf("%s:%d", s.config.PresencePrefix, req.UserID),
		req.Timestamp, s.config.PresenceExpirationTime)
	if err != nil {
		return param.UpsertPresenceResponse{}, richerror.New(op).WithError(err)
	}

	return param.UpsertPresenceResponse{}, nil
}
