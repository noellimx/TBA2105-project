package storing

import (
	"fmt"
	"log"
	"os"
)

func overwriteFilePath(filename string) {

	os.Remove(filename) // I delete the file to avoid duplicated records.

	fmt.Println("Creating db...")
	file, err := os.Create(filename) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Database [%s] created", filename)
	file.Close()
}
