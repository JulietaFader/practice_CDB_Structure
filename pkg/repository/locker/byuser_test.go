package locker

import (
	"context"
	"testing"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/lock"
	"github.com/stretchr/testify/assert"
)

type fakeLocker struct {
	lock.Repository
	LockFn func(ctx context.Context, resource string) (context.Context, error)
}

func (f fakeLocker) Lock(ctx context.Context, resource string) (context.Context, error) {
	return f.LockFn(ctx, resource)
}

func Test_repository_LockByUser(t *testing.T) {
	fake := fakeLocker{}
	r := repository{&fake}

	ctxTest := context.Background()
	fake.LockFn = func(ctx context.Context, resource string) (context.Context, error) {
		assert.Equal(t, "usr-123456", resource)
		assert.Equal(t, ctxTest, ctx)
		return ctx, nil
	}

	ctx, err := r.LockByUser(ctxTest, 123456)
	assert.Nil(t, err)
	assert.Equal(t, ctxTest, ctx)
}
