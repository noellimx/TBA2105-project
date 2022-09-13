package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noellimx/TBA2105-project.git/config"
)

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

type clientT struct {
	c *http.Client
}

func NewClientT() (*clientT, error) {

	return &clientT{c: &http.Client{}}, nil
}

func (cT *clientT) getExample() {

	fmt.Println("[cT.getExample]")

	resp, err := cT.c.Get("http://example.com")

	if err != nil {
		log.Fatalf("Error ")
	}

	defer resp.Body.Close()

	fmt.Printf("Status: %d \n", resp.StatusCode)

}

func main() {

	fmt.Printf("Global Config: %+v \n", globalConfig)

	cT, _ := NewClientT()

	cT.getExample()

}
