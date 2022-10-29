package storing

import (
	"log"
	"os"

	"github.com/noellimx/TBA2105-project/utils"
)

func overwriteFilePath(filename string) {

	os.Remove(filename) // I delete the file to avoid duplicated records.

	log.Println("[overwriteFilePath] Creating file... This is an overwrite operation.")
	file, err := os.Create(filename) // Create SQLite file
	if err != nil {
		utils.VFatal(err.Error())
	}
	log.Printf("Database [%s] created", filename)
	file.Close()
}
