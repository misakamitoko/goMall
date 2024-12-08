package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	EtcdConfig EtcdConfig `yaml:"Etcd"`
	Port       int        `yaml:"Port"`
	Routes     []Route    `yaml:"Routes"`
}

type EtcdConfig struct {
	Hosts       []string `yaml:"Hosts"`
	DialTimeout uint32   `yaml:"DialTimeout"`
}

type Route struct {
	ServiceName string   `yaml:"ServiceName"`
	Path        []string `yaml:"Path"`
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
