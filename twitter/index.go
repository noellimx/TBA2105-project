package twitter

import (
	"context"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type HelloService struct {
	c *twitter.Client
}

func (h *HelloService) helloFollowers() {

	println("[helloFollowers]")
	params := &twitter.FollowerListParams{}
	followers, _, err := h.c.Followers.List(params)

	if err != nil {
		println("Error retrieving followers")
		println(err.Error())
		return
	}

	println("Followers")
	println(followers.Users)
}

func (h *HelloService) helloSearchTweet() {

	println("[helloSearch]")
	results, _, err := h.c.Search.Tweets(&twitter.SearchTweetParams{Query: "hello"})

	if err != nil {
		println("Error retrieving search")
		println(err.Error())
		return
	}

	fmt.Printf("%+v", results)

}

func (h *HelloService) helloSearchUsers() {

	println("[helloSearchUsers]")
	results, _, err := h.c.Users.Search("hello", nil)

	if err != nil {
		println("Error retrieving search")
		println(err.Error())
		return
	}

	fmt.Printf("%+v", results)

}

func HelloPing(c *clientcredentials.Config) {

	httpClient := c.Client(context.TODO())
	client := twitter.NewClient(httpClient)

	h := &HelloService{
		c: client,
	}

	_ = h.helloFollowers
	_ = h.helloSearchTweet
	fn := h.helloSearchUsers

	fn()

}
