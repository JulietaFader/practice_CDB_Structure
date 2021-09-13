package kvs

import (
	"context"
	"errors"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
)

func (r *repository) CreateUser(ctx context.Context, user *process.User) (string, error) {

	if r.chekIfAlreadyExist(ctx, user) {
		return "", errors.New("Error: The User is already exist in DB ...")
	}

	err := r.kvsRepo.Set(ctx, user.ID, user)

	if err != nil {
		return "", errors.New("Error: Trying to create a new User ...")
	}

	return "Created success", nil
}

func (r *repository) chekIfAlreadyExist(ctx context.Context, user *process.User) bool {
	_user, err := r.SearchByEmail(ctx, user.Email)

	if err != nil {
		return false
	}

	if _user.FirstName == user.FirstName && _user.LastName == user.LastName && _user.Status == user.Status {
		return true

	}

	return false
}
