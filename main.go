package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/utils"
	"github.com/noellimx/TBA2105-project/wrangling"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

var YYYYMMDDFrom string = "20210101"
var YYYYMMDDTo string = "20211231"

var query1 string = "jb checkpoint OR jb causeway OR jb customs OR woodlands checkpoint OR woodlands causeway OR woodlands customs OR johor checkpoint OR johor causeway OR johor customs point_radius:[103.7692886848949 1.4526057415829072 25mi]"
var queryWithoutGeo string = "jb checkpoint OR jb causeway OR jb customs OR woodlands checkpoint OR woodlands causeway OR woodlands customs OR johor checkpoint OR johor causeway OR johor customs"

var queryWithGeoNoSearchValShort = "point_radius:[103.7692886848949 1.4526057415829072 3mi]"

var _ string = "point_radius:[103.7692886848949 1.4526057415829072 25mi]"
var query string = queryWithoutGeo

func processProject(fn string) {
	fmt.Printf("[processProject] \n")
	dbcn := storing.NewDBCN_Twitt(fn, false)

	dbcn.CreateTableWords()

	dbcn.CleanTableWords()

	ttt := time.Date(2021, 1, 1, 0, 0, 0, 0, storing.LOCSGTIME)

	now := time.Now()

	_, _, _, _, yyyyddmmhhnow := utils.GoTimeToYYYYMMDDHH(&now)

	log.Printf("[processProject] %s now-> %s", now, yyyyddmmhhnow)

	for {
		_, _, _, _, yyyymmddhh := utils.GoTimeToYYYYMMDDHH(&ttt)

		if yyyyddmmhhnow == yyyymmddhh {
			log.Printf("[processProject] %s", yyyyddmmhhnow)
			break
		}
		log.Printf("[processProject] Processing for: %s", yyyyddmmhhnow)

		ptexts := dbcn.GetTweetsInTheHour(yyyymmddhh)

		for _, ptext := range *ptexts {

			lemmasT := wrangling.LemmatizeText(ptext.Text)
			lemmas := lemmasT.Lemmas

			dbcn.AddWordCounts(yyyymmddhh, lemmas, ptext.RetweetOrFavCount)
		}

		ttt = ttt.Add(time.Hour)
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

	file, err := os.OpenFile(fmt.Sprintf("./data/log-%s.txt", postpent), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("Hello world!")
}

func logPad() {
	log.Println()
	log.Println()
	log.Println()

}
func main() {

	var cmd string

	args := os.Args

	args_l := len(args)

	for i, ar := range args {

		fmt.Printf("args [%d] %s\n", i, ar)

	}

	if args_l == 1 {
		log.Println("No command specified")
		return
	} else {
		cmd = args[1]
	}

	initLog(cmd)

	logPad()
	log.Println("---------main---------")

	log.Printf("command: %s\n", cmd)

	switch cmd {
	case "extract-first":
		extractProject(extFIRST, nil)
	case "extract-two":
		extractProject(extTWO, nil)
	case "extract-prem-some":
		defaultRequestCount := 1
		var requestCount int = defaultRequestCount
		if args_l == 2 {
			log.Printf("[main:extract-some-prem] Request Count Defaulted to %d. \n", requestCount)
		} else {
			requestCount_, err := strconv.Atoi(args[2])
			requestCount = requestCount_
			if err != nil {
				log.Println("[main:extract-some-prem] Invalid request count specified.")
				return
			}
		}
		extractProject(extSOME_Premium, &OptsExtract{RequestCount: requestCount})

	case "extract-prem-all":
		extractProject(extALL_Premium, nil)
	case "process":
		if args_l < 3 {
			log.Println("[main:process] Please specify existing database")
			break
		}
		log.Printf("%t %d %d", (args_l < 2), args_l, 2)
		dbfilename := args[2]
		processProject(dbfilename)
	case "sampletime":
		storing.SampleTwitDateToTimeDate()
	default:
		log.Printf("command [%s] unrecognized", cmd)
	}

	fmt.Println("end.main.end")
	log.Println("----end.main.end---")
	log.Println()
}
