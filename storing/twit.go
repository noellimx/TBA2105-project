package storing

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/noellimx/TBA2105-project/typings"
	"github.com/noellimx/TBA2105-project/utils"
)

type DBCN_Twitt struct {
	db *sql.DB
}

func (dbcn *DBCN_Twitt) createTableTweet() {
	query := `CREATE TABLE tweets (
		"id_str" VARCHAR(50) NOT NULL PRIMARY KEY,
		"yyyymmddhh" VARCHAR(` + fmt.Sprintf("%d", dateStrLength) + `) NOT NULL,
		"yyyy" VARCHAR(4) NOT NULL,
		"mm" CHAR(2) NOT NULL,
		"dd" CHAR(2) NOT NULL,
		"hh" CHAR(2) NOT NULL,
		"text" TEXT NOT NULL,
		"retweet_or_fav_count" INT NOT NULL
	  );`

	log.Println("Creating Tweets table...")
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		utils.VFatal(err.Error())
	}
	statement.Exec()
	statement.Close()
	log.Println("Tweets table created")

}

// working software over comprehensive documentation
// individuals and interactions over processes and tools
// collaboration over contract negotiation
// responding to change over following a plan

func (dbcn *DBCN_Twitt) AddWordCounts(yyyymmddhh string, lemmas []string, retweetOrFavCount int) {
	for _, l := range lemmas {
		dbcn.AddWordCount(yyyymmddhh, l, retweetOrFavCount+1)
	}
}

func (dbcn *DBCN_Twitt) AddWordCount(yyyymmddhh string, lemma string, retweetOrFavCount int) {

	row := dbcn.db.QueryRow("SELECT count FROM words WHERE yyyymmddhh=? AND word=?;", yyyymmddhh, lemma)

	var existingCount string

	err := row.Scan(&existingCount)

	if err != nil {
		log.Printf("[AddWordCount] Not found %s [%s]\n", yyyymmddhh, lemma)
		dbcn.InsertWordCount(yyyymmddhh, lemma)
	}

	addQuery := `UPDATE words SET count = count + ? WHERE yyyymmddhh=? AND word=?;`

	statement, err := dbcn.db.Prepare(addQuery) // Prepare statement.

	if err != nil {
		utils.VFatal(err.Error())
	}

	countToAdd := 1 + retweetOrFavCount

	_, err = statement.Exec(countToAdd, yyyymmddhh, lemma)
	if err != nil {
		utils.VFatal(err.Error())
	}
	defer statement.Close()

}

func (dbcn *DBCN_Twitt) InsertWordCount(yyyymmddhh string, lemma string) {
	log.Printf("[InsertWordCount]\n")
	query := `INSERT INTO words(yyyymmddhh,word) VALUES (?, ?)`
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		utils.VFatal(err.Error())
	}
	_, err = statement.Exec(yyyymmddhh, lemma)
	if err != nil {
		utils.VFatal(err.Error())
	}
	statement.Close()
}

type PText struct {
	Text              string
	RetweetOrFavCount int
}

func (dbcn *DBCN_Twitt) GetTweetsInTheHour(yyyymmddhh string) *[]*PText {
	rows, _ := dbcn.db.Query("SELECT text, retweet_or_fav_count FROM tweets WHERE yyyymmddhh=?;", yyyymmddhh)
	defer rows.Close()

	var Ptexts []*PText
	for rows.Next() { // Iterate and fetch the records from result cursor
		var text string

		var rtFC int
		rows.Scan(&text, &rtFC)

		log.Printf("%s Text: %s Count: %d\n", yyyymmddhh, text, rtFC)

		ptext := &PText{
			Text:              text,
			RetweetOrFavCount: rtFC,
		}
		Ptexts = append(Ptexts, ptext)
	}

	return &Ptexts
}

func (dbcn *DBCN_Twitt) CreateTableWords() {
	query := `CREATE TABLE IF NOT EXISTS words (
		"yyyymmddhh" VARCHAR(` + fmt.Sprintf("%d", dateStrLength) + `) NOT NULL,
		"word" TEXT NOT NULL,
		"count" INT NOT NULL DEFAULT 0
	  );`

	log.Println("Creating Words table...")
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		utils.VFatal(err.Error())
	}
	statement.Exec()
	statement.Close()
	log.Println("Words table created if not exist")
}

func (dbcn *DBCN_Twitt) InsertTweets(tweets []*typings.TweetDB) {

	log.Printf("[InsertTweets] Inserting.... See below for logs.")

	var indexes_log strings.Builder

	for i, t := range tweets {
		indexes_log.Write([]byte(fmt.Sprintf("%d ", i)))
		dbcn.insertTweet(t)
	}

	log.Printf("[InsertTweets] %s", indexes_log.String())

	log.Println("[InsertTweets] End")
}

func (dbcn *DBCN_Twitt) insertTweet(tweet *typings.TweetDB) {

	if tweet == nil {
		log.Printf("\n [Insert Tweet] nil tweet. Returning with no-op. \n")
		return
	}

	query := `INSERT INTO tweets VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		utils.VFatal(err.Error())
	}

	log.Printf("[ResulttoTweetDB] Time: %s ID: %s \n", tweet.Yyyymmddhh, tweet.IdStr)

	_, err = statement.Exec(tweet.IdStr, tweet.Yyyymmddhh, tweet.Yyyy, tweet.Mm, tweet.Dd, tweet.Hh, tweet.Text, tweet.RetweetOrFavCount)
	if err != nil {
		utils.VFatal(err.Error())
	}
}

func NewDBCN_Twitt(dbFileName string, overwrite bool) *DBCN_Twitt {

	if overwrite {
		overwriteFilePath(dbFileName)
	}

	sqliteDatabase, _ := sql.Open(DbDriver, fmt.Sprintf("./%s", dbFileName))

	return &DBCN_Twitt{
		db: sqliteDatabase,
	}
}
func InitTwitDB(overwrite bool, dbName string) *DBCN_Twitt {
	// SQLite is a file based database.

	dbcn := NewDBCN_Twitt(dbName, true)
	dbcn.createTableTweet() // Create Database Tables

	return dbcn
}
