package common

import (
	"encoding/json"
	"os"
)

type Config struct {
	ApiKey string `json:"apiKey"`
	ApiUrl string `json:"apiUrl"`
}

func ReadConfig() *Config {
	// read file from config.json
	bytes, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	// unmarshal json to Config
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}
	return &config
}
