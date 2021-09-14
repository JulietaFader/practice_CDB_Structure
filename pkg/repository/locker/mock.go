package locker

import "github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/lock"

func NewMock(err error, ttl int) Repository {
	return &repository{Repository: lock.NewMock(err, ttl)}
}
