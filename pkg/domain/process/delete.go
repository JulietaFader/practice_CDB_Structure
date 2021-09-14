package process

import (
	"context"
)

func (s service) DeleteUser(ctx context.Context, id string) (string, error) {

	return s.Storage.DeleteUser(ctx, id)
}
