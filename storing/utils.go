package storing

import (
	"fmt"
	"os"

	"github.com/noellimx/TBA2105-project/utils"
)

func overwriteFilePath(filename string) {

	os.Remove(filename) // I delete the file to avoid duplicated records.

	fmt.Println("Creating db...")
	file, err := os.Create(filename) // Create SQLite file
	if err != nil {
		utils.VFatal(err.Error())
	}
	fmt.Printf("Database [%s] created", filename)
	file.Close()
}
