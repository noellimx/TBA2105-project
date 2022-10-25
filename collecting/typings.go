package collecting

type FullResponse struct {
	Results []struct {
		CreatedAt            string      `json:"created_at"`
		ID                   int64       `json:"id"`
		IDStr                string      `json:"id_str"`
		Text                 string      `json:"text"`
		Source               string      `json:"source"`
		Truncated            bool        `json:"truncated"`
		InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
		User                 struct {
			ID                             int64         `json:"id"`
			IDStr                          string        `json:"id_str"`
			Name                           string        `json:"name"`
			ScreenName                     string        `json:"screen_name"`
			Location                       string        `json:"location"`
			URL                            interface{}   `json:"url"`
			Description                    string        `json:"description"`
			TranslatorType                 string        `json:"translator_type"`
			Protected                      bool          `json:"protected"`
			Verified                       bool          `json:"verified"`
			FollowersCount                 int           `json:"followers_count"`
			FriendsCount                   int           `json:"friends_count"`
			ListedCount                    int           `json:"listed_count"`
			FavouritesCount                int           `json:"favourites_count"`
			StatusesCount                  int           `json:"statuses_count"`
			CreatedAt                      string        `json:"created_at"`
			UtcOffset                      interface{}   `json:"utc_offset"`
			TimeZone                       interface{}   `json:"time_zone"`
			GeoEnabled                     bool          `json:"geo_enabled"`
			Lang                           interface{}   `json:"lang"`
			ContributorsEnabled            bool          `json:"contributors_enabled"`
			IsTranslator                   bool          `json:"is_translator"`
			ProfileBackgroundColor         string        `json:"profile_background_color"`
			ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string        `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool          `json:"profile_background_tile"`
			ProfileLinkColor               string        `json:"profile_link_color"`
			ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string        `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
			ProfileImageURL                string        `json:"profile_image_url"`
			ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
			ProfileBannerURL               string        `json:"profile_banner_url"`
			DefaultProfile                 bool          `json:"default_profile"`
			DefaultProfileImage            bool          `json:"default_profile_image"`
			Following                      interface{}   `json:"following"`
			FollowRequestSent              interface{}   `json:"follow_request_sent"`
			Notifications                  interface{}   `json:"notifications"`
			WithheldInCountries            []interface{} `json:"withheld_in_countries"`
		} `json:"user"`
		Geo             interface{} `json:"geo"`
		Coordinates     interface{} `json:"coordinates"`
		Place           interface{} `json:"place"`
		Contributors    interface{} `json:"contributors"`
		RetweetedStatus struct {
			CreatedAt            string      `json:"created_at"`
			ID                   int64       `json:"id"`
			IDStr                string      `json:"id_str"`
			Text                 string      `json:"text"`
			Source               string      `json:"source"`
			Truncated            bool        `json:"truncated"`
			InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
			InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
			InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
			InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
			User                 struct {
				ID                             int           `json:"id"`
				IDStr                          string        `json:"id_str"`
				Name                           string        `json:"name"`
				ScreenName                     string        `json:"screen_name"`
				Location                       string        `json:"location"`
				URL                            string        `json:"url"`
				Description                    string        `json:"description"`
				TranslatorType                 string        `json:"translator_type"`
				Protected                      bool          `json:"protected"`
				Verified                       bool          `json:"verified"`
				FollowersCount                 int           `json:"followers_count"`
				FriendsCount                   int           `json:"friends_count"`
				ListedCount                    int           `json:"listed_count"`
				FavouritesCount                int           `json:"favourites_count"`
				StatusesCount                  int           `json:"statuses_count"`
				CreatedAt                      string        `json:"created_at"`
				UtcOffset                      interface{}   `json:"utc_offset"`
				TimeZone                       interface{}   `json:"time_zone"`
				GeoEnabled                     bool          `json:"geo_enabled"`
				Lang                           interface{}   `json:"lang"`
				ContributorsEnabled            bool          `json:"contributors_enabled"`
				IsTranslator                   bool          `json:"is_translator"`
				ProfileBackgroundColor         string        `json:"profile_background_color"`
				ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
				ProfileBackgroundImageURLHTTPS string        `json:"profile_background_image_url_https"`
				ProfileBackgroundTile          bool          `json:"profile_background_tile"`
				ProfileLinkColor               string        `json:"profile_link_color"`
				ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
				ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
				ProfileTextColor               string        `json:"profile_text_color"`
				ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
				ProfileImageURL                string        `json:"profile_image_url"`
				ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
				ProfileBannerURL               string        `json:"profile_banner_url"`
				DefaultProfile                 bool          `json:"default_profile"`
				DefaultProfileImage            bool          `json:"default_profile_image"`
				Following                      interface{}   `json:"following"`
				FollowRequestSent              interface{}   `json:"follow_request_sent"`
				Notifications                  interface{}   `json:"notifications"`
				WithheldInCountries            []interface{} `json:"withheld_in_countries"`
			} `json:"user"`
			Geo           interface{} `json:"geo"`
			Coordinates   interface{} `json:"coordinates"`
			Place         interface{} `json:"place"`
			Contributors  interface{} `json:"contributors"`
			IsQuoteStatus bool        `json:"is_quote_status"`
			ExtendedTweet struct {
				FullText         string `json:"full_text"`
				DisplayTextRange []int  `json:"display_text_range"`
				Entities         struct {
					Hashtags     []interface{} `json:"hashtags"`
					Urls         []interface{} `json:"urls"`
					UserMentions []interface{} `json:"user_mentions"`
					Symbols      []interface{} `json:"symbols"`
				} `json:"entities"`
			} `json:"extended_tweet"`
			QuoteCount    int `json:"quote_count"`
			ReplyCount    int `json:"reply_count"`
			RetweetCount  int `json:"retweet_count"`
			FavoriteCount int `json:"favorite_count"`
			Entities      struct {
				Hashtags []interface{} `json:"hashtags"`
				Urls     []struct {
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					Indices     []int  `json:"indices"`
				} `json:"urls"`
				UserMentions []interface{} `json:"user_mentions"`
				Symbols      []interface{} `json:"symbols"`
			} `json:"entities"`
			Favorited   bool `json:"favorited"`
			Retweeted   bool `json:"retweeted"`
			EditHistory struct {
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
		IsQuoteStatus bool `json:"is_quote_status"`
		QuoteCount    int  `json:"quote_count"`
		ReplyCount    int  `json:"reply_count"`
		RetweetCount  int  `json:"retweet_count"`
		FavoriteCount int  `json:"favorite_count"`
		Entities      struct {
			Hashtags     []interface{} `json:"hashtags"`
			Urls         []interface{} `json:"urls"`
			UserMentions []struct {
				ScreenName string `json:"screen_name"`
				Name       string `json:"name"`
				ID         int    `json:"id"`
				IDStr      string `json:"id_str"`
				Indices    []int  `json:"indices"`
			} `json:"user_mentions"`
			Symbols []interface{} `json:"symbols"`
		} `json:"entities"`
		Favorited   bool `json:"favorited"`
		Retweeted   bool `json:"retweeted"`
		EditHistory struct {
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
		PossiblySensitive bool `json:"possibly_sensitive,omitempty"`
		ExtendedTweet     struct {
			FullText         string `json:"full_text"`
			DisplayTextRange []int  `json:"display_text_range"`
			Entities         struct {
				Hashtags []struct {
					Text    string `json:"text"`
					Indices []int  `json:"indices"`
				} `json:"hashtags"`
				Urls         []interface{} `json:"urls"`
				UserMentions []interface{} `json:"user_mentions"`
				Symbols      []interface{} `json:"symbols"`
			} `json:"entities"`
		} `json:"extended_tweet,omitempty"`
		DisplayTextRange  []int  `json:"display_text_range,omitempty"`
		QuotedStatusID    int64  `json:"quoted_status_id,omitempty"`
		QuotedStatusIDStr string `json:"quoted_status_id_str,omitempty"`
		QuotedStatus      struct {
			CreatedAt            string      `json:"created_at"`
			ID                   int64       `json:"id"`
			IDStr                string      `json:"id_str"`
			Text                 string      `json:"text"`
			DisplayTextRange     []int       `json:"display_text_range"`
			Source               string      `json:"source"`
			Truncated            bool        `json:"truncated"`
			InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
			InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
			InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
			InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
			User                 struct {
				ID                             int           `json:"id"`
				IDStr                          string        `json:"id_str"`
				Name                           string        `json:"name"`
				ScreenName                     string        `json:"screen_name"`
				Location                       string        `json:"location"`
				URL                            string        `json:"url"`
				Description                    string        `json:"description"`
				TranslatorType                 string        `json:"translator_type"`
				Protected                      bool          `json:"protected"`
				Verified                       bool          `json:"verified"`
				FollowersCount                 int           `json:"followers_count"`
				FriendsCount                   int           `json:"friends_count"`
				ListedCount                    int           `json:"listed_count"`
				FavouritesCount                int           `json:"favourites_count"`
				StatusesCount                  int           `json:"statuses_count"`
				CreatedAt                      string        `json:"created_at"`
				UtcOffset                      interface{}   `json:"utc_offset"`
				TimeZone                       interface{}   `json:"time_zone"`
				GeoEnabled                     bool          `json:"geo_enabled"`
				Lang                           interface{}   `json:"lang"`
				ContributorsEnabled            bool          `json:"contributors_enabled"`
				IsTranslator                   bool          `json:"is_translator"`
				ProfileBackgroundColor         string        `json:"profile_background_color"`
				ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
				ProfileBackgroundImageURLHTTPS string        `json:"profile_background_image_url_https"`
				ProfileBackgroundTile          bool          `json:"profile_background_tile"`
				ProfileLinkColor               string        `json:"profile_link_color"`
				ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
				ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
				ProfileTextColor               string        `json:"profile_text_color"`
				ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
				ProfileImageURL                string        `json:"profile_image_url"`
				ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
				ProfileBannerURL               string        `json:"profile_banner_url"`
				DefaultProfile                 bool          `json:"default_profile"`
				DefaultProfileImage            bool          `json:"default_profile_image"`
				Following                      interface{}   `json:"following"`
				FollowRequestSent              interface{}   `json:"follow_request_sent"`
				Notifications                  interface{}   `json:"notifications"`
				WithheldInCountries            []interface{} `json:"withheld_in_countries"`
			} `json:"user"`
			Geo           interface{} `json:"geo"`
			Coordinates   interface{} `json:"coordinates"`
			Place         interface{} `json:"place"`
			Contributors  interface{} `json:"contributors"`
			IsQuoteStatus bool        `json:"is_quote_status"`
			ExtendedTweet struct {
				FullText         string `json:"full_text"`
				DisplayTextRange []int  `json:"display_text_range"`
				Entities         struct {
					Hashtags []interface{} `json:"hashtags"`
					Urls     []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
					UserMentions []interface{} `json:"user_mentions"`
					Symbols      []interface{} `json:"symbols"`
					Media        []struct {
						ID            int64  `json:"id"`
						IDStr         string `json:"id_str"`
						Indices       []int  `json:"indices"`
						MediaURL      string `json:"media_url"`
						MediaURLHTTPS string `json:"media_url_https"`
						URL           string `json:"url"`
						DisplayURL    string `json:"display_url"`
						ExpandedURL   string `json:"expanded_url"`
						Type          string `json:"type"`
						Sizes         struct {
							Large struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"large"`
							Thumb struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"thumb"`
							Medium struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"medium"`
							Small struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"small"`
						} `json:"sizes"`
					} `json:"media"`
				} `json:"entities"`
				ExtendedEntities struct {
					Media []struct {
						ID            int64  `json:"id"`
						IDStr         string `json:"id_str"`
						Indices       []int  `json:"indices"`
						MediaURL      string `json:"media_url"`
						MediaURLHTTPS string `json:"media_url_https"`
						URL           string `json:"url"`
						DisplayURL    string `json:"display_url"`
						ExpandedURL   string `json:"expanded_url"`
						Type          string `json:"type"`
						Sizes         struct {
							Large struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"large"`
							Thumb struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"thumb"`
							Medium struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"medium"`
							Small struct {
								W      int    `json:"w"`
								H      int    `json:"h"`
								Resize string `json:"resize"`
							} `json:"small"`
						} `json:"sizes"`
					} `json:"media"`
				} `json:"extended_entities"`
			} `json:"extended_tweet"`
			QuoteCount    int `json:"quote_count"`
			ReplyCount    int `json:"reply_count"`
			RetweetCount  int `json:"retweet_count"`
			FavoriteCount int `json:"favorite_count"`
			Entities      struct {
				Hashtags []interface{} `json:"hashtags"`
				Urls     []struct {
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					Indices     []int  `json:"indices"`
				} `json:"urls"`
				UserMentions []interface{} `json:"user_mentions"`
				Symbols      []interface{} `json:"symbols"`
			} `json:"entities"`
			Favorited         bool `json:"favorited"`
			Retweeted         bool `json:"retweeted"`
			PossiblySensitive bool `json:"possibly_sensitive"`
			EditHistory       struct {
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
		ExtendedEntities struct {
			Media []struct {
				ID            int64  `json:"id"`
				IDStr         string `json:"id_str"`
				Indices       []int  `json:"indices"`
				MediaURL      string `json:"media_url"`
				MediaURLHTTPS string `json:"media_url_https"`
				URL           string `json:"url"`
				DisplayURL    string `json:"display_url"`
				ExpandedURL   string `json:"expanded_url"`
				Type          string `json:"type"`
				Sizes         struct {
					Thumb struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"thumb"`
					Small struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"small"`
					Medium struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"medium"`
					Large struct {
						W      int    `json:"w"`
						H      int    `json:"h"`
						Resize string `json:"resize"`
					} `json:"large"`
				} `json:"sizes"`
			} `json:"media"`
		} `json:"extended_entities,omitempty"`
	} `json:"results"`
	Next              string `json:"next"`
	RequestParameters struct {
		MaxResults int    `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	} `json:"requestParameters"`
}

type responseResults struct {
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`

	User struct {
		ID         int64  `json:"id"`
		ScreenName string `json:"screen_name"`
	} `json:"user"`
	RetweetedStatus struct {
		ExtendedTweet struct {
			FullText string `json:"full_text"`
		} `json:"-"`
	} `json:"-"`
}
type SelectedMarshalledResponse struct {
	Results           []responseResults `json:"results"`
	Next              string            `json:"next"`
	RequestParameters struct {
		MaxResults int    `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	} `json:"requestParameters"`
}
