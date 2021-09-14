package kvs

import (
	"context"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

type KVSRepository interface {
	GetUsers(ctx context.Context) ([]process.User, error)
	GetByID(ctx context.Context, id string) (*process.User, error)
	CreateUser(ctx context.Context, user *process.User) (string, error)
	UpdateUser(ctx context.Context, id string, user *process.User) (*process.User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	SearchByEmail(ctx context.Context, email string) (*process.User, error)
}

type repository struct {
	kvsRepo kvs.Repository
	dsRepo ds.Repository
}

func NewRepository(kvsConf kvs.Config, dsConf ds.Config) KVSRepository {
	return &repository{
		kvs.NewRepository(kvsConf),
		ds.NewRepository(dsConf),
	}
}







//var index = 0
