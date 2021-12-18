package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Session struct {
	Key  string `yaml:"key"`
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
	Path string `yaml:"path"`
}

func NewSessionConfig(path string) Session {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}

	var session Session
	err = yaml.Unmarshal(config, &session)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}
	return session
}
