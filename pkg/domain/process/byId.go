package process

import (
	"context"
)

func (s *service) GetByID(ctx context.Context, id string) (*User, error) {

	return s.Storage.GetByID(ctx, id)
}
