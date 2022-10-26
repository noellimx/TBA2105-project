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
		"date_str" VARCHAR(` + fmt.Sprintf("%d", dateStrLength) + `) NOT NULL,
		"yyyy" VARCHAR(4) NOT NULL,
		"mm" CHAR(2) NOT NULL,
		"dd" CHAR(2) NOT NULL,
		"hh" CHAR(2) NOT NULL,
		"text" TEXT NOT NULL
	  );`

	log.Println("Create Tweets table...")
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Tweets table created")

}

func (dbcn *DBCN_Twitt) insertTweet(tweet *typings.TweetDB) {

	/*


		"id_str" VARCHAR(50) NOT NULL PRIMARY KEY,
				"date_str" VARCHAR(12) NOT NULL,
				"yyyy" VARCHAR(4) NOT NULL,
				"mm" CHAR(2) NOT NULL,
				"dd" CHAR(2) NOT NULL,
				"hh" CHAR(2) NOT NULL,
				"text" TEXT NOT NULL
	*/
	log.Println("Inserting tweets record ...")

	query := `INSERT INTO tweets VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := dbcn.db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(tweet.IdStr, tweet.DateStr, tweet.Yyyy, tweet.Mm, tweet.Dd, tweet.Hh, tweet.Text)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (dbcn *DBCN_Twitt) addPennyToStudent(code string) {

	// query := `UPDATE student SET credit = credit + 1 WHERE code = ?` // SQL Statement for Create Table

	// statement, err := dbcn.db.Prepare(query) // Prepare statement.
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// statement.Exec(code)

}

func NewDBCN_Twitt(dbFileName string) *DBCN_Twitt {
	overwriteFilePath(dbFileName)

	sqliteDatabase, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", dbFileName))

	return &DBCN_Twitt{
		db: sqliteDatabase,
	}
}
func InitTwitDB() *DBCN_Twitt {
	// SQLite is a file based database.

	dbcn := NewDBCN_Twitt(TwitDbFileName)
	dbcn.createTableTweets() // Create Database Tables

	return dbcn
}
