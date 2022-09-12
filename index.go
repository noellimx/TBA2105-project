package main

import (
	"fmt"

	"github.com/noellimx/TBA2105-project.git/config"
	"github.com/noellimx/TBA2105-project.git/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	print("Hello world---- \n")

	configPath := "./config.json"

	globalConfig := config.ReadConfig(configPath)

	config := &clientcredentials.Config{
		ClientID:     globalConfig.Twitter.ClientKey,
		ClientSecret: globalConfig.Twitter.ClientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	fmt.Printf("%v", config)

	twitter.HelloPing(config)

	print("Exit")
}
