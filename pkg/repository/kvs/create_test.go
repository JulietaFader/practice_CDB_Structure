package kvs

import (
	"context"
	"errors"
	"testing"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/test"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/common/uuidv4"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"
	"github.com/stretchr/testify/assert"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func Test_Create_User_Success(t *testing.T) {
	ctx := context.Background()
	fr := fakeRepository{}
	r := repository{&fr, &fr}

	user := process.User{
		ID: uuidv4.NewUUIDv4().NewID(),
		FirstName: "Julieta",
		LastName: "Fader",
		Email: "julieta.fader@mercadolibre.com",
		Status: "active",
	}

	fr.SetFn = func(ctx context.Context, key string, value interface{}) error {
		return nil
	}

	fr.FindWhereFn = func(ctx context.Context, conditions *ds.QueryConditions, value interface{}) error {
		return nil
	}

	msg, err := r.CreateUser(ctx, &user)

	assert.Nil(t, err)
	assert.Equal(t, msg, "Created success")
}

func Test_Create_User_Already_Register(t *testing.T) {
	ctx := context.Background()
	fr := fakeRepository{}
	r := repository{&fr, &fr}

	user := process.User{
		ID: uuidv4.NewUUIDv4().NewID(),
		FirstName: "Julieta",
		LastName: "Fader",
		Email: "julieta.fader@mercadolibre.com",
		Status: "active",
	}

	fr.SetFn = func(ctx context.Context, key string, value interface{}) error {
		return nil
	}

	fr.FindWhereFn = func(ctx context.Context, conditions *ds.QueryConditions, value interface{}) error {
		test.UpdateParam(value, user)
		return nil
	}

	msg, err := r.CreateUser(ctx, &user)

	assert.Equal(t, err, errors.New("Error: The User is already exist in DB ..."))
	assert.Equal(t, msg, "")
}
