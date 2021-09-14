package process

import ("context")

func (s service) UpdateUser(ctx context.Context, id string, user *User) (*User, error) {

	return s.Storage.UpdateUser(ctx, id, user)
}
