package ds

import (
	"context"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

type DSRepository interface {
	GetUsers(ctx context.Context) ([]process.User, error)
	SearchByEmail(ctx context.Context, email string) (*process.User, error)
}

type repository struct {
	repo ds.Repository
}

func NewRepository(conf ds.Config) DSRepository {
	return &repository{
		ds.NewRepository(conf),
	}
}