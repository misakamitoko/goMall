package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	EtcdConfig EtcdConfig `yaml:"Etcd"`
}

type EtcdConfig struct {
	Hosts       []string `yaml:"Hosts"`
	DialTimeout uint32   `yaml:"DialTimeout"`
}

func LoadConfig(config *Config, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatal(err)
	}
}
