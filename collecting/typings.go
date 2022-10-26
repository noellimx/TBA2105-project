package collecting

import "github.com/noellimx/TBA2105-project/typings"

type userMentionsResponse struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ID         int    `json:"id"`
	IDStr      string `json:"id_str"`
	Indices    []int  `json:"indices"`
}

type extendedEntitiesResponse struct {
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
}
type userResponse struct {
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
}

type extendedTweetResponse struct {
	FullText         string `json:"full_text"`
	DisplayTextRange []int  `json:"display_text_range"`
	Entities         struct {
		Hashtags     []interface{} `json:"hashtags"`
		Urls         []interface{} `json:"urls"`
		UserMentions []interface{} `json:"user_mentions"`
		Symbols      []interface{} `json:"symbols"`
	} `json:"entities"`
}

type entitiesResponse struct {
	Hashtags []interface{} `json:"hashtags"`
	Urls     []struct {
		URL         string `json:"url"`
		ExpandedURL string `json:"expanded_url"`
		DisplayURL  string `json:"display_url"`
		Indices     []int  `json:"indices"`
	} `json:"urls"`
	UserMentions []userMentionsResponse `json:"user_mentions"`
	Symbols      []interface{}          `json:"symbols"`
}

type SelectedMarshalledResponse struct {
	Results           []*typings.ResponseResults `json:"results"`
	Next              string                     `json:"next"`
	RequestParameters struct {
		MaxResults int    `json:"maxResults"`
		FromDate   string `json:"fromDate"`
		ToDate     string `json:"toDate"`
	} `json:"requestParameters"`
}
