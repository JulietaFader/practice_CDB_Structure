package kvs

import (
	"context"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) GetUsers(ctx context.Context) ([]process.User, error) {
	users := &[]process.User{}

	err := r.dsRepo.FindWhere(ctx, &ds.QueryConditions{}, users)

	if err != nil {
		logger.Error(ctx, "ds_search_error", logger.Tag{"error": err})
		return  nil, err
	}

	return *users, nil
}