package redismatching

import (
	"context"
	"fmt"
	"game-app/entity"
	"game-app/pkg/richerror"
	"github.com/redis/go-redis/v9"
	"time"
)

const WaitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = "redismatching.AddToWaitingList"

	_, err := d.adapter.Client().ZAdd(context.Background(),
		fmt.Sprintf("%s:%s", WaitingListPrefix, category),
		redis.Z{
			Score:  float64(time.Now().UnixMicro()),
			Member: fmt.Sprintf("%d", userID),
		}).Result()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil

}
