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

func DateTwitterToDateDB(dateTwitter string) *time.Time {

	strings.Split(dateTwitter, " ")
	t, err := timefmt.Parse(dateTwitter, "%a %b %d %H:%M:%S +0000 %Y")

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
func ResulttoTweetDB(c *typings.ResponseResults) *typings.TweetDB {

	idStr := c.IdStr
	t := DateTwitterToDateDB(c.CreatedAt)

	log.Printf("[ResulttoTweetDB] Time: %s ID: %s \n", t, c.IdStr)

	yyyy := fmt.Sprintf("%04d", t.Year())
	mm := fmt.Sprintf("%02d", int(t.Month()))
	dd := fmt.Sprintf("%02d", t.Day())
	hh := fmt.Sprintf("%02d", t.Hour())

	text := c.Text

	dateStr := fmt.Sprintf("%s%s%s%s", yyyy, mm, dd, hh)

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
	log.Printf("\n[SampleTwitDate] %d %d %d %d %d %d", some.Year(), int(some.Month()), some.Day(), some.Hour(), some.Minute(), some.Second())
}
