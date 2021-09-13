package process

import (
	"context"
)

func (s *service) GetUsers(ctx context.Context) ([]User, error) {

	//arrayUser, err := s.Storage.GetUsers(ctx)

	return s.Storage.GetUsers(ctx)
}
