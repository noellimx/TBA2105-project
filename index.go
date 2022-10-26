package main

import (
	"fmt"
	"log"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/wrangling"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

func main() {

	fmt.Printf("Global Config: %+v \n", globalConfig)

	cT, err := collecting.GetGlobalClientT(globalConfig)

	if err != nil {
		log.Fatalf(err.Error())

	}

	query := "jb customs OR woodlands checkpoint OR johor causeway OR causeway point_radius:[103.7692886848949 1.4526057415829072 12mi]"
	// cT.getPREMIUMFullArchiveForTheSampleDayLocationSG()

	once := 1

	switch once {
	case 1:
		cT.GetNonPREMIUM30DaysForCustomDateLocationSG_FirstResult(query, "20220925", "20221023")
	case 2:
		cT.GetNonPREMIUM30DaysForCustomDateLocationSG_AllResult(query, "20220925", "20221023")
	}

	storing.InitTwitDB()

	storing.SampleTwitDateToTimeDate()
	return

	storing.SampleDBRun()
	wrangling.LemmaJargonSample()

}
