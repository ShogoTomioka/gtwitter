package cmd

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

	// Send a Tweet
	_, _, err := client.Statuses.Update(text, nil)

	return err
}
