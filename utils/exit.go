package utils

import (
	"fmt"
	"log"
)

func BasicFatal() {
	log.Fatalf("Error ")
}
func VFatal(msg string) {
	fmt.Println("ABORTED")
	log.Fatalf(msg)
}
