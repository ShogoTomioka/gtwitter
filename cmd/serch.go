package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/urfave/cli"
)

func SerchTweet(c *cli.Context) error {
	var (
		config = oauth1.NewConfig(conf.Token.ConsumerKey, conf.Token.ConsumerSecret)
		token  = oauth1.NewToken(conf.Token.AccessToken, conf.Token.AccessSecret)

		httpClient = config.Client(oauth1.NoContext, token)
		client     = twitter.NewClient(httpClient)
	)

	word := c.String("word")
	count := c.Int("count")
	if word == "" {
		panic(fmt.Errorf("Input word to find tweets!"))
	}

	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: word,
	})
	if err != nil {
		panic(fmt.Errorf("Failed to find tweet!!"))
	}

	tweets := search.Statuses
	for i, tweet := range tweets {
		if i == count {
			return nil
		}
		fmt.Println("ユーザ: ", tweet.User.Name)
		fmt.Println(tweet.Text)
	}

	return nil
}
