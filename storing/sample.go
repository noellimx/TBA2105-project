package storing

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var dbFileName string = "tweets.db"

type DBCN struct {
	db *sql.DB
}

func (dbcn *DBCN) insertStudent(code string, name string, program string) {

	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := dbcn.db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func (dbcn *DBCN) createTableStudent() {

	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")

}

func main() {
	os.Remove(dbFileName) // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating db...")
	file, err := os.Create(dbFileName) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Database [%s] created", dbFileName)
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", dbFileName))
	defer sqliteDatabase.Close() // Defer Closing the database

	dbcn := &DBCN{

		db: sqliteDatabase,
	}

	dbcn.createTableStudent() // Create Database Tables

	// INSERT RECORDS
	dbcn.insertStudent("0001", "Liana Kim", "Bachelor")
	dbcn.insertStudent("0002", "Glen Rangel", "Bachelor")

	// DISPLAY INSERTED RECORDS
	displayStudents(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

func (dbcn DBCN) displayStudents() {
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
