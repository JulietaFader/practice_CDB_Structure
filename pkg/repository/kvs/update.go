package kvs

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) UpdateUser(ctx context.Context, id string, user *process.User) (*process.User, error) {
	ok, _, _ := r.chekIfExist(ctx, id)

	if !ok {
		return nil, errors.New(fmt.Sprintf("Error: User not found, wrong Id %s", id))
	}

	err := r.kvsRepo.Set(ctx, user.ID, user)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: Trying to update user"))
	}

	return user, nil
}