package config

import (
	"encoding/json"
	"os"
	"path/filepath"
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

	// Ambil lokasi file .exe
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	// Folder tempat .exe berada
	exeDir := filepath.Dir(exePath)

	// Path config.json
	configPath := filepath.Join(exeDir, "config.json")

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
