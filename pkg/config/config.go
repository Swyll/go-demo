package config

import (
	"go-demo/pkg/utils"
	"io/ioutil"

	"gopkg.in/ini.v1"
)

// CreateFileIfNotExist 如不存在则创建文件
func CreateFileIfNotExist(path string) error {
	ok, err := utils.PathExists(path)
	if ok {
		return err
	}
	return ioutil.WriteFile(path, []byte(""), 0666)

}

// Load 加载配置
func Load(path string, conf interface{}) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}

	return cfg.MapTo(conf)
}

// Save 配置
func Save(path string, conf interface{}) error {
	if err := CreateFileIfNotExist(path); err != nil {
		return err
	}
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}

	err = cfg.ReflectFrom(conf)
	if err != nil {
		return err
	}

	return cfg.SaveTo(path)
}
