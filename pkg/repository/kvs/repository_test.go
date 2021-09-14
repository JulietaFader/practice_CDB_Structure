package kvs

import (
	"context"
	"testing"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"
	"github.com/stretchr/testify/assert"
)

func TestAdaptorRepository_NewRepository(t *testing.T) {

	r := NewRepository(
		kvs.Config{
			Name: "",
		},
		ds.Config{
			Name: "",
		},
	)

	assert.NotNil(t, r)

}

type fakeRepository struct {
	GetFn func(ctx context.Context, key string, value interface{}) error

	SetFn func(ctx context.Context, key string, value interface{}) error

	SetWithTTLFn func(ctx context.Context, key string, value interface{}, ttl int64) error

	SetWithDefaultTTLFn func(ctx context.Context, key string, value interface{}) error

	DeleteFn func(ctx context.Context, key string) error

	FlushCacheFn func(ctx context.Context) error

	FindWhereFn func(ctx context.Context, conditions *ds.QueryConditions, value interface{}) error
}

func (f fakeRepository) Set(ctx context.Context, key string, value interface{}) error {
	return f.SetFn(ctx, key, value)
}

func (f fakeRepository) SetWithTTL(ctx context.Context, key string, value interface{}, ttl int64) error {
	return f.SetWithTTLFn(ctx, key, value, ttl)
}

func (f fakeRepository) SetWithDefaultTTL(ctx context.Context, key string, value interface{}) error {
	return f.SetWithDefaultTTLFn(ctx, key, value)
}

func (f fakeRepository) Delete(ctx context.Context, key string) error {
	return f.DeleteFn(ctx, key)
}

func (f fakeRepository) FlushCache(ctx context.Context) error {
	return f.FlushCacheFn(ctx)
}

func (f fakeRepository) Get(ctx context.Context, key string, value interface{}) error {
	return f.GetFn(ctx, key, value)
}

func (f fakeRepository) FindWhere(ctx context.Context, conditions *ds.QueryConditions, value interface{}) error {
	return f.FindWhereFn(ctx, conditions, value)
}
