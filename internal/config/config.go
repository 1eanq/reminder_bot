package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Config struct {
	APIToken string `yaml:"api_token" env-required:"true"`
	DBPath   string `yaml:"db_path" env-required:"true"`
}

func NewConfig() *Config {
	configPath := FetchConfigPath()
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cfg Config

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func FetchConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		panic("config file path is empty")
	}
	return path
}
