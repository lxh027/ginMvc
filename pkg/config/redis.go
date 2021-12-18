package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Redis struct {
	Env     string `yaml:"env"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Auth    string `yaml:"auth"`
	ConType string `yaml:"con_type"`

	MaxIdle   int           `yaml:"max_idle"`
	MaxActive int           `yaml:"max_active"`
	Timeout   time.Duration `yaml:"timeout"`
}

func NewRedisConfig(path string) Redis {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}

	var redis Redis
	err = yaml.Unmarshal(config, &redis)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}

	return redis
}
