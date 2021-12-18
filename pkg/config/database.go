package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Database struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	DbName      string `yaml:"db_name"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Charset     string `yaml:"charset"`
	ParseTime   string `yaml:"parse_time"`
	MaxIdleCons int    `yaml:"max_idle_cons"`
	MaxOpenCons int    `yaml:"max_open_cons"`
}

func NewDatabaseConfig(path string) Database {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}
	var database Database
	err = yaml.Unmarshal(config, &database)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}

	return database
}
