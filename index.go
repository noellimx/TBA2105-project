package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/noellimx/TBA2105-project/collecting"
	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/typings"
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



func processProject(fn string) {

	dbcn := storing.NewDBCN_Twitt(fn, false)

	dbcn.CreateTableWords()

	tt := typings.NewTTime(2022, 10, 1, 0)

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

func initLog(postpent string) {

	file, err := os.OpenFile(fmt.Sprintf("log-%s.txt", postpent), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	log.Println("---------main---------")
	if args_l == 1 {

		log.Println("No command specified")
		return
	} else {
		cmd = os.Args[1]
	}

	fmt.Println()
	initLog(cmd)

	switch cmd {
	case "extract-first":
		extractProject(extFIRST, nil)
	case "extract-two":
		extractProject(extTWO, nil)
	case "extract-prem-some":

		if args_l < 2 {
			log.Println("[Process extract-some-prem] Please specify how many request to send.")
			return
		}
		requestCount, err := strconv.Atoi(args[2])

		if err != nil {
			log.Println("[Process extract-some-prem] Invalid request count specified.")

			return

		}

		extractProject(extSOME_Premium, &OptsExtract{RequestCount: requestCount})
	case "process":
		if args_l < 2 {
			log.Println("[Process] Please specify existing database")
			return
		}
		filename := args[2]
		processProject(filename)
	default:
		log.Println("command unrecognized")
	}
}
