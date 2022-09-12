package twitter

import (
	"context"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

func Some() {

	config := &clientcredentials.Config{
		TokenURL: "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(context.TODO())

	client := twitter.NewClient(httpClient)

	params := &twitter.FollowerListParams{}
	followers, _, err := client.Followers.List(params)

	if err != nil {
		println("Error retrieving followers")
		println(err.Error())
		return
	}

	println("Followers")
	println(followers.Users)

}
