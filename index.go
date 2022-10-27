package main

import (
	"fmt"
	"os"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/utils"
	"github.com/noellimx/TBA2105-project/wrangling"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

var fromYYYYMMDD string = "20221022"

var query string = "jb checkpoint OR jb causeway OR jb customs OR woodlands checkpoint OR woodlands causeway OR woodlands customs OR johor checkpoint OR johor causeway OR johor customs point_radius:[103.7692886848949 1.4526057415829072 12mi]"

func extractProject(mode extractMode) {
	fmt.Println("[Extract]")
	fmt.Printf("Global Config: %+v \n", globalConfig)
	cT, err := collecting.GetGlobalClientT(globalConfig)
	if err != nil {

		utils.VFatal(err.Error())
	}
	dbcn := storing.InitTwitDB(true)

	switch mode {
	case extFIRST:
		cT.GetAndStoreNonPREMIUM30DaysForCustomDateLocationSG_FirstResult(query, fromYYYYMMDD, "20221023", dbcn)
	case extALL:
		cT.GetAndStoreNonPREMIUM30DaysForCustomDateLocationSG_AllResult(query, fromYYYYMMDD, "20221023")
	}

}

type tTime struct {
	Yyyy int
	Dd   int
	Mm   int
	Hh   int
}

func newTTime(y, m, d, h int) *tTime {
	return &tTime{
		Yyyy: y,
		Dd:   d,
		Mm:   m,
		Hh:   h,
	}
}

var hrsInDay int = 24

func (t *tTime) JumpHour() {

	nextH := (t.Hh + 1)
	if nextH == hrsInDay {
		nextH = 0
		t.jumpDay()
	}
	t.Hh = nextH
}

var daysInYr int = 31

func (t *tTime) jumpDay() {

	nextDay := (t.Dd + 1)
	if nextDay > daysInYr {
		nextDay = 1
		t.jumpMonth()
	}
	t.Dd = nextDay
}

var mmInYr int = 12

func (t *tTime) jumpMonth() {

	nextMth := (t.Mm + 1)
	if nextMth > mmInYr {
		nextMth = 1
		t.jumpYear()
	}
	t.Mm = nextMth
}

func (t *tTime) jumpYear() {
	t.Yyyy += 1
}

func (t *tTime) AsString() string {
	return t.yString() + t.mString() + t.dString() + t.hString()
}

func (t *tTime) yString() string {
	return fmt.Sprintf("%04d", t.Yyyy)
}

func (t *tTime) mString() string {
	return fmt.Sprintf("%02d", t.Mm)
}

func (t *tTime) dString() string {
	return fmt.Sprintf("%02d", t.Dd)
}

func (t *tTime) hString() string {
	return fmt.Sprintf("%02d", t.Hh)
}

func processProject(fn string) {

	dbcn := storing.NewDBCN_Twitt(fn, false)

	dbcn.CreateTableWords()

	tt := newTTime(2022, 10, 16, 0)

	for i := 0; i < 24*23; i++ {

		yyyymmddhh := tt.AsString()
		texts := dbcn.GetTweetsInTheHour(yyyymmddhh)

		for _, text := range *texts {

			lemmasT := wrangling.LemmatizeText(text)
			lemmas := lemmasT.Lemmas

			dbcn.AddWordCounts(yyyymmddhh, lemmas)
		}
		tt.JumpHour()
	}

}

type extractMode int

const (
	extFIRST extractMode = 1
	extTWO   extractMode = 2
	extALL   extractMode = 3
)

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
		extractProject(extFIRST)
	case "extract-two":
		extractProject(extTWO)
	case "process":
		if args_l < 2 {
			fmt.Println("[Process] Please specify existing database")
		}
		filename := args[2]
		processProject(filename)
	default:
		fmt.Println("command unrecognized")
	}
}
