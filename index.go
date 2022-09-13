package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/noellimx/TBA2105-project.git/config"
)

func basicFatal() {
	log.Fatalf("Error ")
}

var CONFIG_PATH string = "./config.json"

var globalConfig = config.ReadConfig(CONFIG_PATH)

var haveClient = false

type clientT struct {
	c *http.Client

	globalConfig *config.GlobalConfig
}

func NewClientT(globalConfig *config.GlobalConfig) (*clientT, error) {
	if haveClient {

		return nil, errors.New("math: square root of negative number")
	}

	haveClient = true
	return &clientT{c: &http.Client{}, globalConfig: globalConfig}, nil
}

func (cT *clientT) getExample() {

	fmt.Println("[cT.getExample]")

	resp, err := cT.c.Get("http://example.com")

	if err != nil {
		basicFatal()
	}

	defer resp.Body.Close()

	fmt.Printf("Status: %d \n", resp.StatusCode)

}

const myUsername string = "noellimx"

func (ct *clientT) twitterExampleGetUserMeV2() {
	fmt.Println("[cT.twitterExample]")

	username := myUsername
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s", username), nil)

	q := req.URL.Query()
	q.Add("api_key", "key_from_environment_or_flag")
	q.Add("another_thing", "foo & bar")

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))
	println(req.URL.RawQuery)
	resp, err := ct.c.Do(req)

	if err != nil {
		basicFatal()
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	fmt.Printf("Status: %d \n", statusCode)
}

func (ct *clientT) twitterExampleRecentSearchV2(query string) {

	fmt.Println("[cT.twitterExampleRecentSearch]")
	url := "https://api.twitter.com/2/tweets/search/recent"

	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	println(req.URL.RawQuery)
	resp, err := ct.c.Do(req)

	if err != nil {
		basicFatal()
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	fmt.Printf("Status: %d \n", statusCode)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s", body)
}

func (ct *clientT) twitterExampleFullArchiveSearch(query string) {

	fmt.Println("[cT.twitterExampleFullArchiveSearch]")
	url := "https://api.twitter.com/2/tweets/search/all"

	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	println(req.URL.RawQuery)
	resp, err := ct.c.Do(req)

	if err != nil {
		basicFatal()
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	fmt.Printf("Status: %d \n", statusCode)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s", body)
}

func main() {

	fmt.Printf("Global Config: %+v \n", globalConfig)

	cT, err := NewClientT(globalConfig)

	if err != nil {
		log.Fatalf(err.Error())
	}

	cT.getExample()
	cT.twitterExampleGetUserMeV2()

	cT.twitterExampleRecentSearchV2("hello")

	cT.twitterExampleFullArchiveSearch("hello")
}
