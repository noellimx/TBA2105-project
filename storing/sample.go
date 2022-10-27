package storing

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var sampleDbFileName string = "sample.db"

type DBCN_Sample struct {
	db *sql.DB
}

func (dbcn *DBCN_Sample) insertStudent(code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := dbcn.db.Prepare(insertStudentSQL) // Prepare statement.

	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)

	statement.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (dbcn *DBCN_Sample) addPennyToStudent(code string) {
	query := `UPDATE student SET credit = credit + 1 WHERE code = ?` // SQL Statement for Create Table

	statement, err := dbcn.db.Prepare(query) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}

	statement.Exec(code)

}

func (dbcn *DBCN_Sample) createTableStudent() {

	createStudentTableSQL := `CREATE TABLE student (
		"code" char(4) NOT NULL PRIMARY KEY,
		"name" TEXT,
		"program" TEXT,
		"credit" int DEFAULT 0
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := dbcn.db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")

}

func newDBCN(dbFileName string) *DBCN_Sample {

	os.Remove(dbFileName) // I delete the file to avoid duplicated records.

	log.Println("Creating db...")
	file, err := os.Create(dbFileName) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Database [%s] created", dbFileName)
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", dbFileName))
	defer sqliteDatabase.Close() // Defer Closing the database

	return &DBCN_Sample{

		db: sqliteDatabase,
	}
}
func SampleDBRun() {
	// SQLite is a file based database.

	dbcn := newDBCN(sampleDbFileName)
	dbcn.createTableStudent() // Create Database Tables

	// INSERT RECORDS
	dbcn.insertStudent("0001", "Liana Kim", "Bachelor")
	dbcn.insertStudent("0002", "Glen Rangel", "Bachelor")
	dbcn.addPennyToStudent("0002")

	// DISPLAY INSERTED RECORDS
	dbcn.displayStudents()
}

func (dbcn *DBCN_Sample) displayStudents() {
	rows, err := dbcn.db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		rows.Scan(&id, &code, &name, &program)
		log.Println("Student: ", code, " ", name, " ", program)
	}
}
