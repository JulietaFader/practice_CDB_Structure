package process

import (
	"context"
)

/*1- Cambiar la capa http por la estructura que usams en core-cdb
2- Cambiar el repositorio de persistencia para que use el kvs de core-lib
Como el service necesita acceder a datos externos de diferentes repos,
hace el pasamano a través del container. Ahi el funcionamiento del container*/

type Service interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user *User) (string, error)
	UpdateUser(ctx context.Context, id string, user *User) (*User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}
type service struct {
	*Container
}
func NewService(cont *Container) Service {
	return &service{
		Container: cont,
	}
}

/*



}*/
