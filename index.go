package main

import (
	"fmt"
	"log"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/config"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

func main() {

	fmt.Printf("Global Config: %+v \n", globalConfig)

	cT, err := collecting.GetGlobalClientT(globalConfig)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// cT.getPREMIUMFullArchiveForTheSampleDayLocationSG()
	cT.GetNonPREMIUM30DaysForTheSampleDayLocationSG_Once()
}
