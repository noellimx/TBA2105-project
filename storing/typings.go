package storing

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/itchyny/timefmt-go"
	"github.com/noellimx/TBA2105-project/typings"
	"github.com/noellimx/TBA2105-project/utils"
)

var dateStrLength int = 10

// %a %b %d %H:%M:%S +0000 %Y

var STD_TWITTER_UTC_STRING_PARSE string = "%a %b %d %H:%M:%S +0000 %Y"

func DateTwitterToDateDB(dateTwitter string) *time.Time {

	strings.Split(dateTwitter, " ")
	t, err := timefmt.Parse(dateTwitter, STD_TWITTER_UTC_STRING_PARSE)
	if err != nil {
		utils.VFatal(err.Error())
	}

	return &t

}

func validateTweetDB(idStr, dateStr, yyyy, mm, dd, hh, text string) bool {

	if len(dateStr) != dateStrLength {
		log.Printf("[validateTweetDB] length of dateStr not %d\n", dateStrLength)
		return false
	}

	return true
}

var LOCSGTIME, _ = time.LoadLocation("Singapore")

func ResulttoTweetDBAndTimeConversion(c *typings.ResponseResults) *typings.TweetDB {

	idStr := c.IdStr
	createdAtLocal_Time := DateTwitterToDateDB(c.CreatedAt).In(LOCSGTIME)
	c.CreatedAt = timefmt.Format(createdAtLocal_Time.In(LOCSGTIME), STD_TWITTER_UTC_STRING_PARSE)
	log.Printf("[ResulttoTweetDB] Time: %s ID: %s \n", createdAtLocal_Time, c.IdStr)

	yyyy, mm, dd, hh, dateStr := utils.GoTimeToYYYYMMDD(&createdAtLocal_Time)

	text := c.Text

	retweetOrFavCount := c.FavoriteCount + c.RetweetCount
	return newTweetDB(idStr, dateStr, yyyy, mm, dd, hh, text, retweetOrFavCount)
}
func newTweetDB(idStr string, dateStr string, yyyy string, mm string, dd string, hh string, text string, rtFC int) *typings.TweetDB {

	ok := validateTweetDB(idStr, dateStr, yyyy, mm, dd, hh, text)

	if !ok {
		return nil
	}
	return &typings.TweetDB{
		IdStr:             idStr,
		Yyyymmddhh:        dateStr,
		Yyyy:              yyyy,
		Mm:                mm,
		Dd:                dd,
		Hh:                hh,
		Text:              text,
		RetweetOrFavCount: rtFC,
	}
}

func SampleTwitDateToTimeDate() {
	twitdate := "Sun Oct 23 11:53:11 +0000 2022"
	some := DateTwitterToDateDB(twitdate)
	log.Printf("\n[SampleTwitDate] %s \n", twitdate)
	log.Printf("\n[SampleTwitDate] %d %d %d %d %d %d\n", some.Year(), int(some.Month()), some.Day(), some.Hour(), some.Minute(), some.Second())
	log.Printf("ori: %s -> local: %s \n", some, some.Local())

	locsg, _ := time.LoadLocation("Singapore")
	log.Printf("ori: %s -> local Local()s: %s \n", some, some.Local())
	log.Printf("ori: %s -> local In: %s \n", some, timefmt.Format(some.In(locsg), STD_TWITTER_UTC_STRING_PARSE))

	timeStart := time.Date(2022, 10, 1, 0, 0, 0, 0, LOCSGTIME)
	fmt.Println("[SampleTwitDate] CONVERT")

	for timeStart.Month() == 10 {
		fmt.Printf("[SampleTwitDate] %s\n", timeStart)
		timeStart = timeStart.AddDate(0, 0, 1)
	}

}
