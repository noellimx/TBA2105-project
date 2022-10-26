package storing

/*
"id_str" VARCHAR(50) NOT NULL PRIMARY KEY,

	"date_str" VARCHAR(12) NOT NULL,
	"yyyy" VARCHAR(4) NOT NULL,
	"mm" CHAR(2) NOT NULL,
	"dd" CHAR(2) NOT NULL,
	"hh" CHAR(2) NOT NULL,
	"text" TEXT NOT NULL
*/
type tweetDB struct {
	IdStr   string
	DateStr string
	Yyyy    string
	Mm      string
	Dd      string
	Hh      string
	Text    string
}

func newTweetDB(IdStr string, DateStr string, Yyyy string, Mm string, Dd string, Hh string, Text string) *tweetDB {
	return &tweetDB{
		IdStr:   IdStr,
		DateStr: DateStr,
		Yyyy:    Yyyy,
		Mm:      Mm,
		Dd:      Dd,
		Hh:      Hh,
		Text:    Text,
	}
}
