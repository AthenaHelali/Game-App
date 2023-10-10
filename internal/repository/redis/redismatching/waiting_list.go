package redismatching

import (
	"context"
	"fmt"
	entity2 "game-app/internal/entity"
	"game-app/internal/pkg/richerror"
	"game-app/internal/pkg/timestamp"
	"github.com/redis/go-redis/v9"
)

const WaitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, category entity2.Category) error {
	const op = "redismatching.AddToWaitingList"

	_, err := d.adapter.Client().ZAdd(context.Background(),
		fmt.Sprintf("%s:%s", WaitingListPrefix, category),
		redis.Z{
			Score:  float64(timestamp.Now()),
			Member: fmt.Sprintf("%d", userID),
		}).Result()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil

}

func (d DB) GetWaitingListByCategory(ctx context.Context, category entity2.Category) ([]entity2.WaitingMember, error) {
	const op = "redismatching.GetWaitingListByCategory"

	//d.adapter.Client().ZRangeWithScores()
	return nil, nil

}
