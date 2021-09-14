package kvs

import (
	"context"
	"testing"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/common/uuidv4"
	"github.com/stretchr/testify/assert"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func Test_Create_User_Success(t *testing.T) {
	ctx := context.Background()
	fr := fakeRepository{}
	r := repository{&fr}

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

	msg, err := r.CreateUser(ctx, &user)

	assert.Nil(t, err)
	assert.Equal(t, msg, "Created success")
}
