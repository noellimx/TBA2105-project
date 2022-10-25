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

	// query := "traffic geocode:1.4521061839361646,103.76931474572983,5mi"

	// query := "causeway traffic jam point_radius:[103.7692886848949 1.4526057415829072 25mi]"

	// query := "jb customs OR woodlands checkpoint OR johor causeway"
	// cT.twitterExample7DaysSearchV1Day(query, "2022", "09", "25", "", 100, "env2")

	query := "jb customs OR woodlands checkpoint OR johor causeway point_radius:[103.7692886848949 1.4526057415829072 12mi]"
	// cT.getPREMIUMFullArchiveForTheSampleDayLocationSG()

	once := 2

	switch once {
	case 1:
		cT.GetNonPREMIUM30DaysForCustomDateLocationSG_FirstResult(query, "20220925", "20221023")
	case 2:
		cT.GetNonPREMIUM30DaysForCustomDateLocationSG_AllResult(query, "20220925", "20221023")
	}

}
