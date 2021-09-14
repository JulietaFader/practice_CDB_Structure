package config

import (
	"sync"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/configmanager"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/ds"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/lock"
)

type Config struct {
	configmanager.BaseConfig
	Service struct {
		Lock                lock.Config       `yaml:"lock"`
		Kvs                 kvs.Config        `yaml:"kvs"`
		Ds                  ds.Config         `yaml:"ds"`
	} `yaml:"services"`
}

var (
	ymlConf Config
	once    sync.Once
)

func Get() Config {
	once.Do(func() {
		ymlConf = Config{}
		configmanager.ReadFromYml(&ymlConf)
	})
	return ymlConf
}
