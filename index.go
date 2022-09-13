package main

import (
	"fmt"

	"github.com/noellimx/TBA2105-project.git/config"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

func main() {

	fmt.Printf("Global Config: %+v", globalConfig)

}
