package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Log struct {
	Path string `yaml:"path"`
}

func NewLogConfig(path string) Log {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}

	var log Log
	err = yaml.Unmarshal(config, &log)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}
	return log
}
