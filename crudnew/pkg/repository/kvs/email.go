package kvs

import (
	"context"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) SearchByEmail(ctx context.Context, email string) (*process.User, error) {
	user := &process.User{}
	condition := map[string]interface{}{"email": email}
	limit := int32(1)

	err := r.dsRepo.FindWhere(ctx, &ds.QueryConditions{
		EQ: condition,
		Limit: &limit,
	}, user)

	if err != nil {
		logger.Error(ctx, "ds_search_error", logger.Tag{"error": err})
		return nil, err
	}

	return user, nil
}
