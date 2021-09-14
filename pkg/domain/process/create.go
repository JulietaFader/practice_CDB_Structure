package process

import ("context")

func (s *service) CreateUser(ctx context.Context, user *User) (string, error) {

	return s.Storage.CreateUser(ctx, user)

}
