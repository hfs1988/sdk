package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type App struct {
	Port int `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`
}

func GetConfig(filename string) Config {
	var conf Config
	// read yaml file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal yaml file & assign to var Config
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
