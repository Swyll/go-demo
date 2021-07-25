package ini

import (
	mconfig "go-demo/config"
	"go-demo/pkg/config"
)

type Ini struct{}

func NewConfig() mconfig.IniParser {
	return Ini{}
}

func (i Ini) Load(path string, conf interface{}) error {
	return config.Load(path, conf)
}

func (i Ini) Save(path string, conf interface{}) error {
	return config.Save(path, conf)
}
