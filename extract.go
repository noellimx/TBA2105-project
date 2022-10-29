package main

import (
	"log"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/utils"
)

type OptsExtract struct {
	RequestCount int
}

// for timeStart.Month() == 10 {
// 	fmt.Printf("[SampleTwitDate] %s\n", timeStart)
// 	timeStart = timeStart.AddDate(0, 0, 1)
// }

func extractProject(mode extractMode, opts *OptsExtract) {
	log.Println("[Extract] --------------------------------------------------------------------------------------------------------------------------------")
	log.Printf("Global Config: %+v \n", globalConfig)
	cT, err := collecting.GetGlobalClientT(globalConfig)
	if err != nil {
		utils.VFatal(err.Error())
	}

	var devEnv *collecting.DevEnv = nil

	var overwrite bool = false

	switch mode {
	case extFIRST:
		devEnv = collecting.NonPremium30Day
		devEnv.RequestCount = 1
	case extTWO:
		devEnv = collecting.NonPremium30Day
		devEnv.RequestCount = 2
	case extALL_Premium:
		devEnv = collecting.PremiumFullArchive
		devEnv.RequestCount = -1
		overwrite = false
	case extSOME_Premium:
		devEnv = collecting.PremiumFullArchive
		if opts != nil {
			devEnv.RequestCount = opts.RequestCount
		}
		overwrite = false
	default:
		log.Println("[Extract] No recognised instruction for extraction.")
		return
	}

	cT.Dbcn = storing.InitTwitDB(overwrite, globalConfig.DB.DBFileName)
	cT.GetAndStore(query, YYYYMMDDFrom, YYYYMMDDTo, devEnv)

}
