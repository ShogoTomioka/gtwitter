package cmd

import (
	"fmt"

	"github.com/ShogoTomioka/gtwitter/lib"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func Tweeting(c *cli.Context) error {

	var (
		config = oauth1.NewConfig(conf.Token.ConsumerKey, conf.Token.ConsumerSecret)
		token  = oauth1.NewToken(conf.Token.AccessToken, conf.Token.AccessSecret)

		httpClient = config.Client(oauth1.NoContext, token)
		client     = twitter.NewClient(httpClient)

		text = c.Args().Get(0)
	)
	if text == "" {
		return nil
	}

	// Send a Tweet
	tweet, _, err := client.Statuses.Update(text, nil)
	color.Yellow(lib.FormatCreatedAt(tweet.CreatedAt))
	color.Green((tweet.User.Name))
	fmt.Println("テキスト: ", tweet.Text)

	return err
}
