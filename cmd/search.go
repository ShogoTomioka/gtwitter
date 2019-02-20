package cmd

import (
	"fmt"

	"github.com/ShogoTomioka/gtwitter/lib"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func SearchTweet(c *cli.Context) error {
	var (
		config = oauth1.NewConfig(conf.Token.ConsumerKey, conf.Token.ConsumerSecret)
		token  = oauth1.NewToken(conf.Token.AccessToken, conf.Token.AccessSecret)

		httpClient = config.Client(oauth1.NoContext, token)
		client     = twitter.NewClient(httpClient)

		word  = c.String("word")
		count = c.Int("count")
	)

	if word == "" {
		color.Red("検索するワードを入力してね, -word or -w")
		return nil
	}

	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: word,
	})
	if err != nil {
		fmt.Errorf("Failed to find tweet!!")
	}

	tweets := search.Statuses
	for i, tweet := range tweets {
		if i == count {
			return nil
		}
		color.Yellow(lib.FormatCreatedAt(tweet.CreatedAt))
		color.Green("ユーザ: " + tweet.User.Name)
		fmt.Println(tweet.Text)
		fmt.Println("-------------------------------")
	}

	return nil
}
