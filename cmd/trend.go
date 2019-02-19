package cmd

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/urfave/cli"
)

func ShowTrend(c *cli.Context) error {
	var (
		config = oauth1.NewConfig(conf.Token.ConsumerKey, conf.Token.ConsumerSecret)
		token  = oauth1.NewToken(conf.Token.AccessToken, conf.Token.AccessSecret)

		httpClient = config.Client(oauth1.NoContext, token)
		client     = twitter.NewClient(httpClient)
	)

	number, err := strconv.Atoi(c.String("number"))
	if err != nil {
		panic(fmt.Errorf("Incorrect input number !!"))
	}

	param := &twitter.TrendsPlaceParams{}

	// 東京のWOEIDをセット
	woeid := int64(1110809)
	trends, _, _ := client.Trends.Place(woeid, param)
	for _, w := range trends {
		for n, t := range w.Trends {
			if n == number {
				return nil
			}
			fmt.Println(t.Name, "\nURL:", t.URL)
		}
	}

	return nil
}
