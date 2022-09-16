package utils

import "log"

func BasicFatal() {
	log.Fatalf("Error ")
}
func VFatal(msg string) {
	println("ABORTED")
	log.Fatalf(msg)
}
