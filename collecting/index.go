package collecting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/typings"
	"github.com/noellimx/TBA2105-project/utils"
)

type DevEnv struct {
	RequestCount         int
	MaxResultsPerRequest int
	Env                  string
	Endpoint             string
}

var NonPremium30Day *DevEnv = &DevEnv{
	MaxResultsPerRequest: 100,
	Env:                  "env2",
	Endpoint:             "30day",
}

var PremiumFullArchive *DevEnv = &DevEnv{
	MaxResultsPerRequest: 500,
	Env:                  "env1",
	Endpoint:             "fullarchive",
}
var httpMethods = &struct {
	post string
	get  string
}{post: "POST", get: "GEt"}

type ClientTWit struct {
	c            *http.Client
	Dbcn         *storing.DBCN_Twitt
	globalConfig *config.GlobalConfig
}

var haveClient = false

var globalClient *ClientTWit

func GetGlobalClientT(globalConfig *config.GlobalConfig) (*ClientTWit, error) {
	if haveClient {

		return nil, errors.New("math: square root of negative number")
	}

	haveClient = true
	globalClient = &ClientTWit{c: &http.Client{}, globalConfig: globalConfig}
	return globalClient, nil
}

// next is the current
func (ct *ClientTWit) twitterSearch1_1(query string, yyyymmdd_s string, yyyymmdd_e string, next string, devEnv *DevEnv) (nextnext string, tweets []*typings.TweetDB) {

	// maxResults int, env string, mode string
	maxResults := devEnv.MaxResultsPerRequest
	env := devEnv.Env
	endpoint := devEnv.Endpoint

	fn_name := "[cT.twitterSearch1_1]"
	log.Println(fn_name)

	// 1. Forming Post Body Map
	hhmmStart := "0000"
	hhmmEnd := "2359"

	postBodyMap := make(map[string]string)

	postBodyMap["query"] = query
	postBodyMap["fromDate"] = fmt.Sprintf("%s%s", yyyymmdd_s, hhmmStart)
	postBodyMap["toDate"] = fmt.Sprintf("%s%s", yyyymmdd_e, hhmmEnd)
	postBodyMap["maxResults"] = fmt.Sprintf("%d", maxResults)

	if next != "" {
		postBodyMap["next"] = next
	}
	postBody, _ := json.Marshal(postBodyMap)
	responseBody := bytes.NewBuffer(postBody)

	// 2. Form HTTPS Request
	url := fmt.Sprintf("https://api.twitter.com/1.1/tweets/search/%s/%s.json", endpoint, env)

	log.Println(url)
	req, _ := http.NewRequest(httpMethods.post, url, responseBody)

	q := req.URL.Query()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.globalConfig.Twitter.Bearer))
	req.Header.Add("Content-Type", "application/json")

	log.Printf("[cT.twitterSearch1_1] Query: %s \n", query)

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	log.Printf("[cT.twitterSearch1_1] Raw Query : %s \n", req.URL.RawQuery)

	// 3. Execute Request

	log.Printf("[cT.twitterSearch1_1] Attempt to do client req. ptr to req: %p \n", &req)

	resp, err := ct.c.Do(req)
	if err != nil {

		log.Println("[cT.twitterSearch1_1] Do client request error")
		utils.VFatal(err.Error())
	}
	defer resp.Body.Close()

	// 4. Read
	statusCode := resp.StatusCode

	log.Printf("[%s] Status: %d \n", fn_name, statusCode)

	body, _ := io.ReadAll(resp.Body)
	if statusCode != 200 {
		utils.VFatal(string(body))
	}
	writeDataJSONPath := fmt.Sprintf("twitterSearch1_1-%s-%s-%s.json", postBodyMap["fromDate"], postBodyMap["toDate"], next)
	f, err := os.Create(writeDataJSONPath)

	if err != nil {
		utils.VFatal(err.Error())
	}

	// 5. Process

	bodyJSON := &SelectedMarshalledResponse{}
	json.Unmarshal(body, bodyJSON)
	bodyJSON.RequestParameters.Query = postBodyMap["query"]
	log.Printf("Needful Data JSON-Writing to: %s\n", writeDataJSONPath)

	var tweetDBs []*typings.TweetDB
	for idx, result := range bodyJSON.Results {
		if result.ExtendedTweet.FullText != "" {
			bodyJSON.Results[idx].Text = result.ExtendedTweet.FullText
			twDB := storing.ResulttoTweetDB(bodyJSON.Results[idx])
			tweetDBs = append(tweetDBs, twDB)
		}
	}
	data, _ := json.Marshal(bodyJSON)
	f.Write(data)
	return bodyJSON.Next, tweetDBs
}

func (cT *ClientTWit) GetAndStore(query string, yyyymmddFrom string, yyyymmddTo string, devEnv *DevEnv) {
	log.Printf("[GetAndStore] ---------------- \n")

	var total int = devEnv.RequestCount

	var requestCount int = devEnv.RequestCount

	var infinite bool = false
	next := ""

	if requestCount == -1 {
		infinite = true
		log.Printf("[GetAndStore] [%d](Infinite request cycles) \n", requestCount)
	}

	log.Printf("[GetAndStore] Environment: %s Request Count: %d Endpoint: %s \n", devEnv.Env, devEnv.RequestCount, devEnv.Endpoint)

	for infinite || requestCount > 0 {
		log.Printf("[GetAndStore] RemainingRequest[%d/%d] Searching in 2 secs... \n", requestCount, total)
		time.Sleep(2 * time.Second)
		next_, tweetDBs := cT.twitterSearch1_1(query, yyyymmddFrom, yyyymmddTo, next, devEnv)
		next = next_
		cT.Dbcn.InsertTweets(tweetDBs)

		log.Printf("next: [%s]\n", next_)

		if next_ == "" {
			break
		}

		devEnv.RequestCount--
	}
}

func (cT *ClientTWit) demos() {

}
func (cT *ClientTWit) Demos() {
	cT.demos()
}
