package redispresence

import (
	"context"
	"game-app/pkg/richerror"
	"time"
)

func (d DB) Upsert(ctx context.Context, key string, timeStamp int64, expTime time.Duration) error {
	const op = "redispresence.Upsert"
	_, err := d.adapter.Client().Set(ctx, key, timeStamp, expTime).Result()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
