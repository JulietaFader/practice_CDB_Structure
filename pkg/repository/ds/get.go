package ds

import (
	"context"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) GetUsers(ctx context.Context) ([]process.User, error) {
	domUsers := &[]process.User{}

	err := r.repo.FindWhere(ctx, &ds.QueryConditions{}, domUsers)

	if err != nil {
		logger.Error(ctx, "ds_search_error", logger.Tag{"error": err})
		return  nil, err
	}

	return *domUsers, nil
}