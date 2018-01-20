package main

import (
	"github.com/ChimeraCoder/anaconda"
)

//SendTweet sends a tweet. Don't forget your setup keys!
func SendTweet(message string) {

	anaconda.SetConsumerKey("your-consumer-key")
	anaconda.SetConsumerSecret("your-consumer-secret")
	api := anaconda.NewTwitterApi("your-access-token",
		"your-access-token-secret")

	api.PostTweet(message, nil)
}
