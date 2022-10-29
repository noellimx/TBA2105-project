package utils

import (
	"log"
)

func BasicFatal() {
	log.Fatalf("Error ")
}
func VFatal(msg string) {
	log.Println("ABORTED")
	log.Println(msg)
	log.Fatalf(msg)
}
