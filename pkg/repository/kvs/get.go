package kvs

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) GetByID(ctx context.Context, id string) (*process.User, error) {
	ok, user, _ := r.chekIfExist(ctx, id)

	if ok {
		return user, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Error: User not found, wrong Id %s", id))
	}
}

func (r *repository) chekIfExist(ctx context.Context, id string) (bool, *process.User, int) {
	users := &[]process.User{}
	err := r.kvsRepo.Get(ctx, KSVUserKeyName, users)

	if err != nil {
		return  false, nil, -1
	}

	for i, user := range *users {
		if user.ID == id {
			return true, &user, i
		}
	}

	return false, nil, -1
}