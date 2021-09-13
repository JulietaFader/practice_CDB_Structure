package container

import (
	"github.com/mercadolibre/fury_asset-mgmt-core-cdb/cmd/app/config"
	"github.com/mercadolibre/fury_asset-mgmt-core-cdb/pkg/repository/ds"
	"github.com/mercadolibre/fury_asset-mgmt-core-cdb/pkg/repository/kvs"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

type storageRepo struct {
	kvs.KVSRepository
	ds.DsRepository
}

func Process (config config.Config) *process.Container {
	kvsRepo := kvs.NewRepository(config.Service.Kvs)
	dsRepo := ds.NewRepository(config.Service.Ds)

	storage := storageRepo{
		kvsRepo,
		dsRepo,
	}
	return &process.Container{
		Storage: storage,
	}
}