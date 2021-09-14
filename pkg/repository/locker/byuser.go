package locker

import (
	"context"
	"fmt"
)

const (
	userLockKeyFormat = "usr-%d"
)

func (r repository) LockByUser(ctx context.Context, userID uint64) (context.Context, error) {
	return r.Repository.Lock(ctx, fmt.Sprintf(userLockKeyFormat, userID))
}
