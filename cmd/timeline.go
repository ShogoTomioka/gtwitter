package cmd

import (
	"fmt"
	"strconv"

	"github.com/ShogoTomioka/gtwitter/lib"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func GetTimeline(c *cli.Context) error {

	var (
		config = oauth1.NewConfig(conf.Token.ConsumerKey, conf.Token.ConsumerSecret)
		token  = oauth1.NewToken(conf.Token.AccessToken, conf.Token.AccessSecret)

		httpClient = config.Client(oauth1.NoContext, token)
		client     = twitter.NewClient(httpClient)
	)

	count, err := strconv.Atoi(c.String("count"))
	if err != nil {
		fmt.Errorf("Incorrect input number !!")
	}

	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: count,
	})
	if err != nil {
		fmt.Errorf("Failed to get timelines!!")
	}

	for _, tweet := range tweets {
		color.Yellow(lib.FormatCreatedAt(tweet.CreatedAt))
		color.Green(("ユーザ:" + tweet.User.Name))
		fmt.Println("テキスト: ", tweet.Text)
		fmt.Println("-------------------------------------------------")
	}

	return err
}
