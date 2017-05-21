package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/vivalapanda/edenradiostatus/fetchhtml"
)

func main() {
}

func pollForDj() (isBot bool, djName string, err error) {
	djString, pollError := fetchhtml.PollUrlForID("http://edenofthewest.com/", "status-dj")
	if pollError != nil {
		return false, "", pollError
	}

	if djString == "Bot-sama" {
		return true, "", nil
	}

	return false, djName, nil
}

func sendTweet() (err error) {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, _, tweetErr := client.Statuses.Update("Just testing bot config", nil)
	if tweetErr != nil {
		return tweetErr
	}
	fmt.Printf("Tweet: %v\n", tweet)

	return
}
