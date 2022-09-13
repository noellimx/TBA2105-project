package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type GlobalConfig struct {
	Twitter struct {
		ClientKey      string `json:"client_key"`
		Bearer         string `json:"bearer"`
		ClientSecret   string `json:"client_secret"`
		DevEnvironment string `json:"dev_environment"`
	} `json:"twitter"`
}

func ReadConfig(path string) *GlobalConfig {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Error Reading Config from path. " + err.Error())
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	globalConfig := GlobalConfig{}
	json.Unmarshal(byteValue, &globalConfig)

	return &globalConfig
}
