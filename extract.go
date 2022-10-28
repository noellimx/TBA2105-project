package main

import (
	"log"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/utils"
)

func extractProject(mode extractMode, opts *OptsExtract) {
	log.Println("[Extract] --------------------------------------------------------------------------------------------------------------------------------")
	log.Printf("Global Config: %+v \n", globalConfig)
	cT, err := collecting.GetGlobalClientT(globalConfig)
	if err != nil {
		utils.VFatal(err.Error())
	}
	cT.Dbcn = storing.InitTwitDB(true)

	var devEnv *collecting.DevEnv = nil

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
	case extSOME_Premium:
		devEnv = collecting.PremiumFullArchive
		if opts != nil {
			devEnv.RequestCount = opts.RequestCount
		}
	default:
		log.Println("[Extract] No recognised instruction for extraction.")
		return
	}
	cT.GetAndStore(query, YYYYMMDDFrom, YYYYMMDDTo, devEnv)

}
