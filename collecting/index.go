package collecting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/utils"
)

type requestParameters struct {
}

type response200FullArchiveSearch struct {
	Next              string            `json:"next"`
	RequestParameters requestParameters `json:"requestParameters"`
}

var httpMethods = &struct {
	post string
	get  string
}{post: "POST", get: "GEt"}


type ClientT struct {
	c *http.Client

	globalConfig *config.GlobalConfig
}

var haveClient = false

var globalClient *ClientT

func GetGlobalClientT(globalConfig *config.GlobalConfig) (*ClientT, error) {
	if haveClient {

		return nil, errors.New("math: square root of negative number")
	}

	haveClient = true
	globalClient = &ClientT{c: &http.Client{}, globalConfig: globalConfig}
	return globalClient, nil
}

func (ct *ClientT) TwitterPremiumFullArchiveSearchV1(query string, yy string, mm string, dd string, next string, maxResults int) string {

	fn_name := "[cT.twitterExampleFullArchiveSearchV1]"
	fmt.Println(fn_name)

	// 1. Forming Post Body Map
	hhmmStart := "0000"
	hhmmEnd := "2359"

	maxResults = 500

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

	fmt.Println(url)
	req, _ := http.NewRequest(httpMethods.post, url, responseBody)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))
	req.Header.Add("Content-Type", "application/json")

	fmt.Printf("Query: %s ", query)

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	println(req.URL.RawQuery)

	// 3. Execute Request
	resp, err := ct.c.Do(req)

	if err != nil {
		utils.VFatal(err.Error())
	}
	defer resp.Body.Close()

	// 4. Read
	statusCode := resp.StatusCode

	fmt.Printf("[%s] Status: %d \n", fn_name, statusCode)

	body, _ := io.ReadAll(resp.Body)
	if statusCode != 200 {
		utils.VFatal(string(body))
	}
	writeBodyToPath := fmt.Sprintf("twitterExampleFullArchiveSearchV1-%s-%s-%s-%s-%s.json", postBodyMap["query"], postBodyMap["maxResults"], postBodyMap["fromDate"], postBodyMap["toDate"], next)
	f, err := os.Create(writeBodyToPath)

	if err != nil {
		utils.VFatal(err.Error())
	}

	// 5. Process

	bodyJSON := &response200FullArchiveSearch{}
	json.Unmarshal(body, bodyJSON)
	fmt.Printf("Writing to : %s", writeBodyToPath)
	f.Write(body)
	return bodyJSON.Next
}

func (ct *ClientT) twitterExample7DaysSearchV1Day(query string, yy string, mm string, dd string, next string, maxResults int, env string) string {

	fn_name := "[cT.twitterExample7DaysSearchV1]"
	fmt.Println(fn_name)

	// 1. Forming Post Body Map
	hhmmStart := "0000"
	hhmmEnd := "2359"

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
	url := fmt.Sprintf("https://api.twitter.com/1.1/tweets/search/30day/%s.json", env)

	fmt.Println(url)
	req, _ := http.NewRequest(httpMethods.post, url, responseBody)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))
	req.Header.Add("Content-Type", "application/json")

	fmt.Printf("Query: %s ", query)

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	println(req.URL.RawQuery)

	// 3. Execute Request
	resp, err := ct.c.Do(req)

	if err != nil {
		utils.VFatal(err.Error())
	}
	defer resp.Body.Close()

	// 4. Read
	statusCode := resp.StatusCode

	fmt.Printf("[%s] Status: %d \n", fn_name, statusCode)

	body, _ := io.ReadAll(resp.Body)
	if statusCode != 200 {
		utils.VFatal(string(body))
	}
	writeBodyToPath := fmt.Sprintf("twitterExampleFullArchiveSearchV1-%s-%s-%s-%s-%s.json", postBodyMap["query"], postBodyMap["maxResults"], postBodyMap["fromDate"], postBodyMap["toDate"], next)
	f, err := os.Create(writeBodyToPath)

	if err != nil {
		utils.VFatal(err.Error())
	}

	// 5. Process

	bodyJSON := &response200FullArchiveSearch{}
	json.Unmarshal(body, bodyJSON)
	fmt.Printf("Writing to : %s", writeBodyToPath)
	f.Write(body)
	return bodyJSON.Next
}

func (ct *ClientT) twitterExample7DaysSearchVDayCustom(query string, s_date string, e_date string, next string, maxResults int, env string) string {

	fn_name := "[cT.twitterExample7DaysSearchVDayCustom]"
	fmt.Println(fn_name)

	// 1. Forming Post Body Map
	hhmmStart := "0000"
	hhmmEnd := "2359"

	postBodyMap := make(map[string]string)

	postBodyMap["query"] = query
	postBodyMap["fromDate"] = fmt.Sprintf("%s%s", s_date, hhmmStart)
	postBodyMap["toDate"] = fmt.Sprintf("%s%s", e_date, hhmmEnd)
	postBodyMap["maxResults"] = fmt.Sprintf("%d", maxResults)

	if next != "" {
		postBodyMap["next"] = next
	}
	postBody, _ := json.Marshal(postBodyMap)
	responseBody := bytes.NewBuffer(postBody)

	// 2. Form HTTPS Request
	url := fmt.Sprintf("https://api.twitter.com/1.1/tweets/search/30day/%s.json", env)

	fmt.Println(url)
	req, _ := http.NewRequest(httpMethods.post, url, responseBody)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))
	req.Header.Add("Content-Type", "application/json")

	fmt.Printf("Query: %s ", query)

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	println(req.URL.RawQuery)

	// 3. Execute Request
	resp, err := ct.c.Do(req)

	if err != nil {
		utils.VFatal(err.Error())
	}
	defer resp.Body.Close()

	// 4. Read
	statusCode := resp.StatusCode

	fmt.Printf("[%s] Status: %d \n", fn_name, statusCode)

	body, _ := io.ReadAll(resp.Body)
	if statusCode != 200 {
		utils.VFatal(string(body))
	}
	writeBodyToPath := fmt.Sprintf("twitterExampleFullArchiveSearchV1-%s-%s-%s-%s-%s.json", postBodyMap["query"], postBodyMap["maxResults"], postBodyMap["fromDate"], postBodyMap["toDate"], next)
	f, err := os.Create(writeBodyToPath)

	if err != nil {
		utils.VFatal(err.Error())
	}

	// 5. Process

	bodyJSON := &response200FullArchiveSearch{}
	json.Unmarshal(body, bodyJSON)
	fmt.Printf("Writing to : %s", writeBodyToPath)
	f.Write(body)
	return bodyJSON.Next
}
func (cT *ClientT) getFullArchiveForTheSampleDay() {
	next := ""
	for {
		next = cT.TwitterPremiumFullArchiveSearchV1("hello", "2021", "12", "01", next, 100)
		if next == "" {
			break
		}
		print("looping \n")
	}

}

func (cT *ClientT) getPREMIUMFullArchiveForTheSampleDayLocationSG() {
	next := ""
	for {
		next = cT.TwitterPremiumFullArchiveSearchV1("Hello", "2021", "12", "01", next, 100)

		fmt.Printf("next: [%s]\n", next)

		if next == "" {
			break
		}

		print("looping \n")
	}

}

func (cT *ClientT) getNonPREMIUM30DaysForTheSampleDayLocationSG() {

	next := ""
	for {
		next = cT.twitterExample7DaysSearchV1Day("traffic geocode:1.4521061839361646,103.76931474572983,5mi", "2022", "09", "25", next, 100, "env2")

		fmt.Printf("next: [%s]\n", next)

		if next == "" || true {
			break
		}
		print("looping \n")
	}

}

func (cT *ClientT) GetNonPREMIUM30DaysForTheSampleDayLocationSG_Once() {
	// query := "traffic geocode:1.4521061839361646,103.76931474572983,5mi"
	query := "jb customs OR woodlands checkpoint OR johor causeway point_radius:[103.7692886848949 1.4526057415829072 25mi]"
	// query := "causeway traffic jam point_radius:[103.7692886848949 1.4526057415829072 25mi]"

	// query := "jb customs OR woodlands checkpoint OR johor causeway"
	// cT.twitterExample7DaysSearchV1Day(query, "2022", "09", "25", "", 100, "env2")

	cT.twitterExample7DaysSearchVDayCustom(query, "20220925", "20221023", "", 100, "env2")
}

func (cT *ClientT) demos() {

}
func (cT *ClientT) Demos() {
	cT.demos()
}
