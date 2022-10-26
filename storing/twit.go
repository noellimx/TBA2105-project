package storing

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type DBCN_Twitt struct {
	db *sql.DB
}

func (dbcn *DBCN_Twitt) createTableTweets() {
	query := `CREATE TABLE tweets (
		"id_str" VARCHAR(50) NOT NULL PRIMARY KEY,
		"date_str" VARCHAR(12) NOT NULL,
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

func (dbcn *DBCN_Twitt) insertTweet(tweet *tweetDB) {

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

	dbcn.displayStudents()

	return dbcn
}

func (dbcn DBCN_Twitt) displayStudents() {
	row, err := dbcn.db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code, &name, &program)
		log.Println("Student: ", code, " ", name, " ", program)
	}
}
