package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Secret  string `yaml:"secret"`
	Address string `yaml:"address"`
}

func ParseConfig(filename string) *Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("fail to read config file %s: %s", filename, err)
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("fail to unmarshal config: %s", err)
	}

	return config
}
