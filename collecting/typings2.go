package collecting

type FullResponse struct {
	Results []struct {
		CreatedAt            string       `json:"created_at"`
		ID                   int64        `json:"id"`
		IDStr                string       `json:"id_str"`
		Text                 string       `json:"text"`
		Source               string       `json:"source"`
		Truncated            bool         `json:"truncated"`
		InReplyToStatusID    interface{}  `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{}  `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{}  `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{}  `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  interface{}  `json:"in_reply_to_screen_name"`
		User                 userResponse `json:"user"`
		Geo                  interface{}  `json:"geo"`
		Coordinates          interface{}  `json:"coordinates"`
		Place                interface{}  `json:"place"`
		Contributors         interface{}  `json:"contributors"`
		RetweetedStatus      struct {
			CreatedAt            string                `json:"created_at"`
			ID                   int64                 `json:"id"`
			IDStr                string                `json:"id_str"`
			Text                 string                `json:"text"`
			Source               string                `json:"source"`
			Truncated            bool                  `json:"truncated"`
			InReplyToStatusID    interface{}           `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{}           `json:"in_reply_to_status_id_str"`
			InReplyToUserID      interface{}           `json:"in_reply_to_user_id"`
			InReplyToUserIDStr   interface{}           `json:"in_reply_to_user_id_str"`
			InReplyToScreenName  interface{}           `json:"in_reply_to_screen_name"`
			User                 userResponse          `json:"user"`
			Geo                  interface{}           `json:"geo"`
			Coordinates          interface{}           `json:"coordinates"`
			Place                interface{}           `json:"place"`
			Contributors         interface{}           `json:"contributors"`
			IsQuoteStatus        bool                  `json:"is_quote_status"`
			ExtendedTweet        extendedTweetResponse `json:"extended_tweet"`
			QuoteCount           int                   `json:"quote_count"`
			ReplyCount           int                   `json:"reply_count"`
			RetweetCount         int                   `json:"retweet_count"`
			FavoriteCount        int                   `json:"favorite_count"`
			Entities             entitiesResponse      `json:"entities"`
			Favorited            bool                  `json:"favorited"`
			Retweeted            bool                  `json:"retweeted"`
			EditHistory          struct {
				InitialTweetID string   `json:"initial_tweet_id"`
				EditTweetIds   []string `json:"edit_tweet_ids"`
			} `json:"edit_history"`
			EditControls struct {
				EditableUntilMs int64 `json:"editable_until_ms"`
				EditsRemaining  int   `json:"edits_remaining"`
			} `json:"edit_controls"`
			Editable    bool   `json:"editable"`
			FilterLevel string `json:"filter_level"`
			Lang        string `json:"lang"`
		} `json:"retweeted_status,omitempty"`
		IsQuoteStatus bool             `json:"is_quote_status"`
		QuoteCount    int              `json:"quote_count"`
		ReplyCount    int              `json:"reply_count"`
		RetweetCount  int              `json:"retweet_count"`
		FavoriteCount int              `json:"favorite_count"`
		Entities      entitiesResponse `json:"entities"`
		Favorited     bool             `json:"favorited"`
		Retweeted     bool             `json:"retweeted"`
		EditHistory   struct {
			InitialTweetID string   `json:"initial_tweet_id"`
			EditTweetIds   []string `json:"edit_tweet_ids"`
		} `json:"edit_history"`
		EditControls struct {
			EditableUntilMs int64 `json:"editable_until_ms"`
			EditsRemaining  int   `json:"edits_remaining"`
		} `json:"edit_controls"`
		Editable      bool   `json:"editable"`
		FilterLevel   string `json:"filter_level"`
		Lang          string `json:"lang"`
		MatchingRules []struct {
			Tag interface{} `json:"tag"`
		} `json:"matching_rules"`
		PossiblySensitive bool                  `json:"possibly_sensitive,omitempty"`
		ExtendedTweet     extendedTweetResponse `json:"extended_tweet,omitempty"`
		DisplayTextRange  []int                 `json:"display_text_range,omitempty"`
		QuotedStatusID    int64                 `json:"quoted_status_id,omitempty"`
		QuotedStatusIDStr string                `json:"quoted_status_id_str,omitempty"`
		QuotedStatus      struct {
			CreatedAt            string                `json:"created_at"`
			ID                   int64                 `json:"id"`
			IDStr                string                `json:"id_str"`
			Text                 string                `json:"text"`
			DisplayTextRange     []int                 `json:"display_text_range"`
			Source               string                `json:"source"`
			Truncated            bool                  `json:"truncated"`
			InReplyToStatusID    interface{}           `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{}           `json:"in_reply_to_status_id_str"`
			InReplyToUserID      interface{}           `json:"in_reply_to_user_id"`
			InReplyToUserIDStr   interface{}           `json:"in_reply_to_user_id_str"`
			InReplyToScreenName  interface{}           `json:"in_reply_to_screen_name"`
			User                 userResponse          `json:"user"`
			Geo                  interface{}           `json:"geo"`
			Coordinates          interface{}           `json:"coordinates"`
			Place                interface{}           `json:"place"`
			Contributors         interface{}           `json:"contributors"`
			IsQuoteStatus        bool                  `json:"is_quote_status"`
			ExtendedTweet        extendedTweetResponse `json:"extended_tweet"`
			QuoteCount           int                   `json:"quote_count"`
			ReplyCount           int                   `json:"reply_count"`
			RetweetCount         int                   `json:"retweet_count"`
			FavoriteCount        int                   `json:"favorite_count"`
			Entities             entitiesResponse      `json:"entities"`
			Favorited            bool                  `json:"favorited"`
			Retweeted            bool                  `json:"retweeted"`
			PossiblySensitive    bool                  `json:"possibly_sensitive"`
			EditHistory          struct {
				InitialTweetID string   `json:"initial_tweet_id"`
				EditTweetIds   []string `json:"edit_tweet_ids"`
			} `json:"edit_history"`
			EditControls struct {
				EditableUntilMs int64 `json:"editable_until_ms"`
				EditsRemaining  int   `json:"edits_remaining"`
			} `json:"edit_controls"`
			Editable    bool   `json:"editable"`
			FilterLevel string `json:"filter_level"`
			Lang        string `json:"lang"`
		} `json:"quoted_status,omitempty"`
		QuotedStatusPermalink struct {
			URL      string `json:"url"`
			Expanded string `json:"expanded"`
			Display  string `json:"display"`
		} `json:"quoted_status_permalink,omitempty"`
		ExtendedEntities extendedEntitiesResponse `json:"extended_entities,omitempty"`
	} `json:"results"`
	Next              string `json:"next"`
	RequestParameters struct {
		MaxResults int    `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	} `json:"requestParameters"`
}
