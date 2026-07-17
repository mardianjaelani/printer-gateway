package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`

	Cors struct {
		AllowOrigins []string `json:"allow_origins"`
	} `json:"cors"`
}

func Load() (*Config, error) {

	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	err = json.NewDecoder(file).Decode(cfg)

	return cfg, err
}
