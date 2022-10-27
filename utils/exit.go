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
	fmt.Println(msg)
	log.Fatalf(msg)
}
