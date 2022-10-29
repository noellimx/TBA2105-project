package typings

type ResponseResults struct {
	CreatedAt string `json:"created_at"`
	IdStr     string `json:"id_str"`
	RetweetCount  int              `json:"retweet_count"`
	FavoriteCount int              `json:"favorite_count"`
	Text          string `json:"text"`
	Truncated     bool   `json:"truncated"`
	ExtendedTweet struct {
		FullText string `json:"full_text"`
	} `json:"extended_tweet,omitempty"`

	User struct {
		ID         int64  `json:"id"`
		ScreenName string `json:"screen_name"`
	} `json:"user"`
	RetweetedStatus struct {
		ExtendedTweet struct {
			FullText string `json:"full_text"`
		} `json:"extended_tweet"`
	} `json:"retweeted_status"`
}
