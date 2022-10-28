package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/noellimx/TBA2105-project/utils"
)

type GlobalConfig struct {
	Twitter struct {
		ClientKey      string `json:"client_key"`
		Bearer         string `json:"bearer"`
		ClientSecret   string `json:"client_secret"`
		DevEnvironment string `json:"dev_environment"`
	} `json:"twitter"`

	Now time.Time

	DB struct {
		DBFileName string `json:"db_file_name"`
	} `json:"db"`
}

func ReadConfig(path string) *GlobalConfig {
	jsonFile, err := os.Open(path)
	if err != nil {
		utils.VFatal("Error Reading Config from path. " + err.Error())
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		utils.VFatal(err.Error())
	}
	globalConfig := GlobalConfig{}
	json.Unmarshal(byteValue, &globalConfig)

	var Now = time.Now()

	globalConfig.Now = Now

	var _ string = fmt.Sprintf("twitter-%d%02d%02d%02d%02d.db", int(Now.Month()), Now.Day(), Now.Hour(), Now.Minute(), Now.Second())

	return &globalConfig
}
