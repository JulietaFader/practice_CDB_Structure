package kvs

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) GetByID(ctx context.Context, id string) (*process.User, error) {
	ok, user := r.chekIfExist(ctx, id)

	if ok {
		return user, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Error: User not found, wrong Id %s", id))
	}
}

func (r *repository) chekIfExist(ctx context.Context, id string) (bool, *process.User) {
	user := &process.User{}
	err := r.kvsRepo.Get(ctx, id, user)

	if err != nil {
		return  false, nil
	}

	if user.ID == id {
		return true, user
	}

	return false, nil
}