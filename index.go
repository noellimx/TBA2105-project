package main

import (
	"log"

	"github.com/noellimx/TBA2105-project.git/config"
	"github.com/noellimx/TBA2105-project.git/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

var CONFIG_PATH string = "./config.json"

func main() {
	print("Hello world---- \n")

	globalConfig := config.ReadConfig(CONFIG_PATH)

	config := &clientcredentials.Config{
		ClientID:     globalConfig.Twitter.ClientKey,
		ClientSecret: globalConfig.Twitter.ClientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	if config.ClientID == "" || config.ClientSecret == "" {

		log.Fatal("Client credentials not supplied")
	}
	twitter.HelloPing(config)

	print("Exit")
}
