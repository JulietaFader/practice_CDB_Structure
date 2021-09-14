package locker

import (
	"context"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/lock"
)

type Repository interface {
	LockByUser(ctx context.Context, userID uint64) (context.Context, error)
	Unlock(ctx context.Context) error
	KeepAlive(ctx context.Context) (context.Context, error)
}

type repository struct {
	lock.Repository
}

func NewRepository(config lock.Config) Repository {
	return &repository{Repository: lock.NewRepository(config)}
}
