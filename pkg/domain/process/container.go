package process

import (
	"context"
)

type Container struct {
	Storage Storage
	Lock Locker
}

type Storage interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) (string, error)
	UpdateUser(ctx context.Context, id string, user *User) (*User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}

type Locker interface {
	LockByUser(ctx context.Context, userID uint64) (context.Context, error)
	Unlock(ctx context.Context) error
	KeepAlive(ctx context.Context) (context.Context, error)
}