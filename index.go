package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/noellimx/TBA2105-project.git/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type GlobalConfig struct {
	Twitter struct {
		ClientId     string "json:`client_id`"
		ClientSecret string "json:`client_secret`"
	} "json:`twitter`"
}

func readConfig(path string) *GlobalConfig {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Error Reading Config from path. " + err.Error())
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	globalConfig := GlobalConfig{}
	json.Unmarshal(byteValue, &globalConfig)

	return &globalConfig
}

func main() {
	print("Hello world---- \n")

	config := &clientcredentials.Config{
		TokenURL: "https://api.twitter.com/oauth2/token",
	}

	twitter.Some()

	print("Exit")
}
