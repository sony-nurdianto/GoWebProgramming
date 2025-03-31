package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Configuration struct {
	Address      string
	ReadTimeout  int8
	WriteTimeout int16
	Static       string
	DatabaseUrl  string
}

func LoadConfig() (*Configuration, error) {
	configPath, err := filepath.Abs("config.json")
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
