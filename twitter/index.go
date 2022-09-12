package twitter

import (
	"context"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

func HelloPing(c *clientcredentials.Config) {

	httpClient := c.Client(context.TODO())

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
