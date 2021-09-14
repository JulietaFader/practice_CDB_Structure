package kvs

import (
	"context"
	"errors"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) CreateUser(ctx context.Context, user *process.User) (string, error) {
	var users []process.User
	_ = r.kvsRepo.Get(ctx, KSVUserKeyName, &users)

	users = append(users, *user)

	err := r.kvsRepo.Set(ctx, KSVUserKeyName, users)

	if err != nil {
		return "", errors.New("Error: Trying to create a new User ...")
	}

	return "Created success", nil
}