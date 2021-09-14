package kvs

import (
	"context"
	"errors"
	"fmt"
)

func (r *repository) DeleteUser(ctx context.Context, id string) (string, error) {
	ok, _ := r.chekIfExist(ctx, id)

	if !ok {
		return "", errors.New(fmt.Sprintf("Error: User not found, wrong Id %s", id))
	}

	err := r.kvsRepo.Delete(ctx, id)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: Trying to delete user"))
	}

	return "Delete success!", nil
}

/*
func (r *repository) DeleteUser(ctx context.Context, id string) (string, error) {

	for index, u := range r.repo {

		if u.ID == id {

			r.repo = append(r.repo[:index], r.repo[index+1:]...)

			return "Delete success!", nil

		}

	}
	return "", errors.New(fmt.Sprintf("Error: User not found, wrong Id %s", id))
}
*/