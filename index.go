package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/utils"
	"github.com/noellimx/TBA2105-project/wrangling"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

var YYYYMMDDFrom string = "20221001"
var YYYYMMDDTo string = "20221025"

var query string = "jb checkpoint OR jb causeway OR jb customs OR woodlands checkpoint OR woodlands causeway OR woodlands customs OR johor checkpoint OR johor causeway OR johor customs point_radius:[103.7692886848949 1.4526057415829072 25mi]"

type OptsExtract struct {
	RequestCount int
}

func extractProject(mode extractMode, opts *OptsExtract) {
	fmt.Println("[Extract]")
	fmt.Printf("Global Config: %+v \n", globalConfig)
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
		fmt.Println("[Extract] No recognised instruction for extraction.")
		return
	}
	cT.GetAndStore(query, YYYYMMDDFrom, YYYYMMDDTo, devEnv)

}



var hrsInDay int = 24

func (t *TTime) JumpHour() {

	nextH := (t.Hh + 1)
	if nextH == hrsInDay {
		nextH = 0
		t.jumpDay()
	}
	t.Hh = nextH
}

var daysInYr int = 31

func (t *TTime) jumpDay() {

	nextDay := (t.Dd + 1)
	if nextDay > daysInYr {
		nextDay = 1
		t.jumpMonth()
	}
	t.Dd = nextDay
}

var mmInYr int = 12

func (t *TTime) jumpMonth() {

	nextMth := (t.Mm + 1)
	if nextMth > mmInYr {
		nextMth = 1
		t.jumpYear()
	}
	t.Mm = nextMth
}

func (t *TTime) jumpYear() {
	t.Yyyy += 1
}

func (t *TTime) AsString() string {
	return t.yString() + t.mString() + t.dString() + t.hString()
}

func (t *TTime) yString() string {
	return fmt.Sprintf("%04d", t.Yyyy)
}

func (t *TTime) mString() string {
	return fmt.Sprintf("%02d", t.Mm)
}

func (t *TTime) dString() string {
	return fmt.Sprintf("%02d", t.Dd)
}

func (t *TTime) hString() string {
	return fmt.Sprintf("%02d", t.Hh)
}

func processProject(fn string) {

	dbcn := storing.NewDBCN_Twitt(fn, false)

	dbcn.CreateTableWords()

	tt := NewTTime(2022, 10, 1, 0)

	hours := 24
	days := 30
	for i := 0; i < hours*days; i++ {

		yyyymmddhh := tt.AsString()
		ptexts := dbcn.GetTweetsInTheHour(yyyymmddhh)

		for _, ptext := range *ptexts {

			lemmasT := wrangling.LemmatizeText(ptext.Text)
			lemmas := lemmasT.Lemmas

			dbcn.AddWordCounts(yyyymmddhh, lemmas, ptext.RetweetOrFavCount)
		}
		tt.JumpHour()
	}

}

type extractMode int

const (
	extFIRST        extractMode = 1
	extTWO          extractMode = 2
	extALL_Premium  extractMode = 3
	extSOME_Premium extractMode = 4
)

func initLog() {

	file, err := os.OpenFile("log-", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("Hello world!")
}
func main() {

	var cmd string

	args := os.Args

	args_l := len(args)

	if args_l == 1 {

		fmt.Println("No command specified")
		return
	} else {
		cmd = os.Args[1]
	}

	switch cmd {
	case "extract-first":
		extractProject(extFIRST, nil)
	case "extract-two":
		extractProject(extTWO, nil)
	case "extract-prem-some":

		if args_l < 2 {
			fmt.Println("[Process extract-some-prem] Please specify how many request to send.")
			return
		}
		requestCount, err := strconv.Atoi(args[2])

		if err != nil {
			fmt.Println("[Process extract-some-prem] Invalid request count specified.")

			return

		}

		extractProject(extSOME_Premium, &OptsExtract{RequestCount: requestCount})
	case "process":
		if args_l < 2 {
			fmt.Println("[Process] Please specify existing database")
			return
		}
		filename := args[2]
		processProject(filename)
	default:
		fmt.Println("command unrecognized")
	}
}
