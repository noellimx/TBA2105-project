package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/noellimx/TBA2105-project.git/config"
)

var httpMethods = &struct {
	post string
	get  string
}{post: "POST", get: "GEt"}

func basicFatal() {
	log.Fatalf("Error ")
}

func vFatal(msg string) {
	log.Fatalf(msg)

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
		vFatal(err.Error())
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
	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("%s \n", body)
}

type requestParameters struct {
}

type response200FullArchiveSearch struct {
	Next              string            `json:"next"`
	RequestParameters requestParameters `json:"requestParameters"`
}

func (ct *clientT) twitterExampleFullArchiveSearchV1(query string, yy string, mm string, dd string, next string, maxResults int) string {

	fn_name := "[cT.twitterExampleFullArchiveSearchV1]"
	fmt.Println(fn_name)

	// 1. Forming Post Body Map
	hhmmStart := "0000"
	hhmmEnd := "2359"

	if 10 < maxResults || maxResults < 100 {
		maxResults = 100

		fmt.Printf("invalid maxResults. Defaulted to %d\n", maxResults)
	}

	postBodyMap := make(map[string]string)

	postBodyMap["query"] = query
	postBodyMap["fromDate"] = fmt.Sprintf("%s%s%s%s", yy, mm, dd, hhmmStart)
	postBodyMap["toDate"] = fmt.Sprintf("%s%s%s%s", yy, mm, dd, hhmmEnd)
	postBodyMap["maxResults"] = fmt.Sprintf("%d", maxResults)

	if next != "" {
		postBodyMap["next"] = next
	}
	postBody, _ := json.Marshal(postBodyMap)
	responseBody := bytes.NewBuffer(postBody)

	// 2. Form HTTPS Request
	url := fmt.Sprintf("https://api.twitter.com/1.1/tweets/search/fullarchive/%s.json", ct.globalConfig.Twitter.DevEnvironment)

	req, _ := http.NewRequest(httpMethods.post, url, responseBody)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))
	req.Header.Add("Content-Type", "application/json")

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	println(req.URL.RawQuery)

	// 3. Execute Request
	resp, err := ct.c.Do(req)

	if err != nil {
		basicFatal()
	}
	defer resp.Body.Close()

	// 4. Read
	statusCode := resp.StatusCode

	fmt.Printf("[%s] Status: %d \n", fn_name, statusCode)

	body, _ := io.ReadAll(resp.Body)

	writeBodyToPath := fmt.Sprintf("twitterExampleFullArchiveSearchV1-%s-%s-%s-%s-%s.json", postBodyMap["query"], postBodyMap["maxResults"], postBodyMap["fromDate"], postBodyMap["toDate"], next)
	f, err := os.Create(writeBodyToPath)

	if err != nil {
		vFatal(err.Error())
	}

	// 5. Process

	bodyJSON := &response200FullArchiveSearch{}
	json.Unmarshal(body, bodyJSON)

	f.Write(body)
	return bodyJSON.Next
}

func demos(cT *clientT) {
	cT.getExample()
	cT.twitterExampleGetUserMeV2()
	cT.twitterExampleRecentSearchV2("hello")

}
func Demos(cT *clientT) {
	demos(cT)
}

func getFullArchiveForTheSampleDay(cT *clientT) {
	next := ""
	for {
		next = cT.twitterExampleFullArchiveSearchV1("hello", "2012", "12", "01", next, 100)
		if next == "" {
			break
		}
		print("looping \n")
	}
}

func main() {

	fmt.Printf("Global Config: %+v \n", globalConfig)

	cT, err := NewClientT(globalConfig)

	if err != nil {
		log.Fatalf(err.Error())
	}

	getFullArchiveForTheSampleDay(cT)

}
