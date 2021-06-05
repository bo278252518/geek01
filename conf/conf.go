package conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/pkg/errors"
)

type Config struct {
	Dsn string `yaml:"dsn"`
}

func NewConfig(configFile string) (*Config, error) {
	fileSource := file.NewSource(
		file.WithPath(configFile),
	)
	conf, _ := config.NewConfig()

	if err := conf.Load(fileSource); err != nil {
		return nil, errors.Wrapf(err, "conf.Load path[%s] failed.", configFile)
	}

	var c Config
	if err := conf.Scan(&c); err != nil {
		return nil, errors.Wrap(err, "conf.Scan failed.")
	}
	return &c, nil
}
