package kvs

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) DeleteUser(ctx context.Context, id string) (string, error) {
	var users []process.User
	ok, _, i := r.chekIfExist(ctx, id)

	if !ok {
		return "", errors.New(fmt.Sprintf("Error: User not found, wrong Id %s", id))
	}

	err := r.kvsRepo.Get(ctx, KSVUserKeyName, &users)

	users = append(users[:i], users[i+1:]...)

	err = r.kvsRepo.Set(ctx, KSVUserKeyName, users)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: Trying to delete user"))
	}

	return "Delete success!", nil
}