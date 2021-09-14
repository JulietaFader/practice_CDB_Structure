package main

import (
	"github.com/mercadolibre/crudnew/cmd/config"
	"github.com/mercadolibre/crudnew/pkg/domain/process"
	"github.com/mercadolibre/crudnew/pkg/repository/kvs"
	"github.com/mercadolibre/crudnew/pkg/repository/locker"
	"github.com/mercadolibre/crudnew/pkg/rest"
	restKVS "github.com/mercadolibre/crudnew/pkg/rest/reader"
)

func main() {
	conf := config.Get()


	//Mock que simula la db de kvs??
	var lockerRepository process.Locker
	lockerRepository = locker.NewMock(nil, 3000)

	userRepo := kvs.NewRepository(conf.Service.Kvs)

	userCont := process.Container{
		Storage: userRepo,
		Lock: lockerRepository,
	}
	userService := process.NewService(&userCont)
	userHandler := restKVS.NewHandler(userService)

	if err := rest.API(userHandler).Run(); err != nil {
		panic(err.Error())
	}
}
