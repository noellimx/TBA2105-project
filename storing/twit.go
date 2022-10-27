package storing

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/noellimx/TBA2105-project/typings"
)

type DBCN_Twitt struct {
	db *sql.DB
}

func (dbcn *DBCN_Twitt) createTableTweets() {
	query := `CREATE TABLE tweets (
		"id_str" VARCHAR(50) NOT NULL PRIMARY KEY,
		"yyyymmddhh" VARCHAR(` + fmt.Sprintf("%d", dateStrLength) + `) NOT NULL,
		"yyyy" VARCHAR(4) NOT NULL,
		"mm" CHAR(2) NOT NULL,
		"dd" CHAR(2) NOT NULL,
		"hh" CHAR(2) NOT NULL,
		"text" TEXT NOT NULL
	  );`

	log.Println("Creating Tweets table...")
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	statement.Close()
	log.Println("Tweets table created")

}

// working software over comprehensive documentation
// individuals and interactions over processes and tools
// collaboration over contract negotiation
// responding to change over following a plan

func (dbcn *DBCN_Twitt) AddWordCounts(yyyymmddhh string, lemmas []string) {
	for _, l := range lemmas {
		dbcn.AddWordCount(yyyymmddhh, l)
	}
}

func (dbcn *DBCN_Twitt) AddWordCount(yyyymmddhh string, lemma string) {

	row := dbcn.db.QueryRow("SELECT count FROM words WHERE yyyymmddhh=? AND word=?;", yyyymmddhh, lemma)

	var count string

	err := row.Scan(&count)

	if err != nil {
		fmt.Printf("[AddWordCount] Not found %s [%s]\n", yyyymmddhh, lemma)
		dbcn.InsertWordCount(yyyymmddhh, lemma)
	}

	addQuery := `UPDATE words SET count = count + 1 WHERE yyyymmddhh=? AND word=?;`

	statement, err := dbcn.db.Prepare(addQuery) // Prepare statement.

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(yyyymmddhh, lemma)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer statement.Close()

}

func (dbcn *DBCN_Twitt) InsertWordCount(yyyymmddhh string, lemma string) {
	fmt.Printf("[InsertWordCount]\n")
	query := `INSERT INTO words(yyyymmddhh,word) VALUES (?, ?)`
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(yyyymmddhh, lemma)
	if err != nil {
		log.Fatalln(err.Error())
	}
	statement.Close()
}

func (dbcn *DBCN_Twitt) GetTweetsInTheHour(yyyymmddhh string) *[]string {
	rows, _ := dbcn.db.Query("SELECT text FROM tweets WHERE yyyymmddhh=?;", yyyymmddhh)
	defer rows.Close()

	var texts []string
	for rows.Next() { // Iterate and fetch the records from result cursor
		var text string
		rows.Scan(&text)
		log.Printf("%s Text: %s \n", yyyymmddhh, text)
		texts = append(texts, text)
	}

	return &texts
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
		log.Fatal(err.Error())
	}
	statement.Exec()
	statement.Close()
	log.Println("Words table created if not exist")
}

func (dbcn *DBCN_Twitt) InsertTweets(tweets []*typings.TweetDB) {
	for i, t := range tweets {
		fmt.Printf("[InsertTweets] @index %d \n", i)
		dbcn.InsertTweet(t)
	}
	fmt.Println("[InsertTweets] End\n")
}

func (dbcn *DBCN_Twitt) InsertTweet(tweet *typings.TweetDB) {

	if tweet == nil {
		fmt.Printf("[Insert Tweet] nil tweet. Returning with no-op. \n")
		return
	}
	/*
		"id_str" VARCHAR(50) NOT NULL PRIMARY KEY,
				"yyyymmddhh" VARCHAR(12) NOT NULL,
				"yyyy" VARCHAR(4) NOT NULL,
				"mm" CHAR(2) NOT NULL,
				"dd" CHAR(2) NOT NULL,
				"hh" CHAR(2) NOT NULL,
				"text" TEXT NOT NULL
	*/
	log.Printf("[Insert Tweet] Inserting tweet record ... \n")

	query := `INSERT INTO tweets VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(tweet.IdStr, tweet.Yyyymmddhh, tweet.Yyyy, tweet.Mm, tweet.Dd, tweet.Hh, tweet.Text)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func NewDBCN_Twitt(dbFileName string, overwrite bool) *DBCN_Twitt {

	if overwrite {
		overwriteFilePath(dbFileName)
	}

	sqliteDatabase, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", dbFileName))

	return &DBCN_Twitt{
		db: sqliteDatabase,
	}
}
func InitTwitDB(overwrite bool) *DBCN_Twitt {
	// SQLite is a file based database.

	dbcn := NewDBCN_Twitt(TwitDbFileName, true)
	dbcn.createTableTweets() // Create Database Tables

	return dbcn
}
