package collecting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/noellimx/TBA2105-project/config"
	"github.com/noellimx/TBA2105-project/storing"
	"github.com/noellimx/TBA2105-project/typings"
	"github.com/noellimx/TBA2105-project/utils"
)

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

func (ct *ClientT) twitterSearch1_1(query string, s_yyyymmdd string, e_yyyymmdd string, next string, maxResults int, env string) (string, []*typings.TweetDB) {

	fn_name := "[cT.twitterSearch1_1]"
	fmt.Println(fn_name)

	// 1. Forming Post Body Map
	hhmmStart := "0000"
	hhmmEnd := "2359"

	postBodyMap := make(map[string]string)

	postBodyMap["query"] = query
	postBodyMap["fromDate"] = fmt.Sprintf("%s%s", s_yyyymmdd, hhmmStart)
	postBodyMap["toDate"] = fmt.Sprintf("%s%s", e_yyyymmdd, hhmmEnd)
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

	fmt.Printf("[cT.twitterSearch1_1] Query: %s \n", query)

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.RawQuery)

	// 3. Execute Request

	fmt.Printf("[cT.twitterSearch1_1] Attempt to do client req: %p \n", &req)

	resp, err := ct.c.Do(req)
	fmt.Printf("remaining %s\n", resp.Header[http.CanonicalHeaderKey("x-rate-limit-reset")])
	fmt.Printf("remaining %s\n", resp.Header[http.CanonicalHeaderKey("x-rate-limit-remaining")])
	if err != nil {

		fmt.Println("[cT.twitterSearch1_1] Do client request error")
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
	writeBodyToPath := fmt.Sprintf("twitterSearch1_1-%s-%s-%s.json", postBodyMap["fromDate"], postBodyMap["toDate"], next)
	f, err := os.Create(writeBodyToPath)

	if err != nil {
		utils.VFatal(err.Error())
	}

	// 5. Process

	bodyJSON := &SelectedMarshalledResponse{}
	json.Unmarshal(body, bodyJSON)
	bodyJSON.RequestParameters.Query = postBodyMap["query"]
	fmt.Printf("Writing to : %s", writeBodyToPath)

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

var nonPremiumEnv string = "env2"

func (cT *ClientT) GetAndStoreNonPREMIUM30DaysForCustomDateLocationSG_FirstResult(query string, yyyymmddFrom string, yyyymmddTo string, dbcn *storing.DBCN_Twitt) {

	_, tweetDBs := cT.twitterSearch1_1(query, yyyymmddFrom, yyyymmddTo, "", 100, nonPremiumEnv)

	dbcn.InsertTweets(tweetDBs)

}

func (cT *ClientT) GetAndStoreNonPREMIUM30DaysForCustomDateLocationSG_TwoResult(query string, yyyymmddFrom string, yyyymmddTo string, dbcn *storing.DBCN_Twitt) {

	next, tweetDBs := cT.twitterSearch1_1(query, yyyymmddFrom, yyyymmddTo, "", 100, nonPremiumEnv)

	dbcn.InsertTweets(tweetDBs)

	if next != "" {
		_, tweetDBs := cT.twitterSearch1_1(query, yyyymmddFrom, yyyymmddTo, "", 100, nonPremiumEnv)
		dbcn.InsertTweets(tweetDBs)
	}
}

func (cT *ClientT) GetAndStoreNonPREMIUM30DaysForCustomDateLocationSG_AllResult(query string, yyyymmddFrom string, yyyymmddTo string) {
	next := ""
	for {
		fmt.Printf("[GetNonPREMIUM30DaysForCustomDateLocationSG_AllResult] Searching in 2 secs... \n")
		time.Sleep(2 * time.Second)
		next, _ = cT.twitterSearch1_1(query, yyyymmddFrom, yyyymmddTo, next, 100, nonPremiumEnv)

		fmt.Printf("next: [%s]\n", next)

		if next == "" {
			break
		}
		print("looping \n")
	}
}

func (cT *ClientT) demos() {

}
func (cT *ClientT) Demos() {
	cT.demos()
}
