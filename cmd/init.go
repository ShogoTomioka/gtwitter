package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var App cli.App

func init() {

	App.Commands = []cli.Command{
		{
			Name:   "tweet",
			Usage:  "tweet with text",
			Action: Tweeting,
		},
		{
			Name:  "timeline",
			Usage: "show timeline",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "count,c",
					Value: "20",
					Usage: "set the limit of timelines, default is 20",
				}},
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "trend",
			Usage: "get trends",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "count,c",
					Value: "20",
					Usage: "number of getting trends, maximum is 50, default is 20",
				}},
			Action: ShowTrend,
		},
		{
			Name:  "config",
			Usage: "change configuration",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "consumer,c",
					Value: "",
					Usage: "set consumerKey",
				},
				cli.StringFlag{
					Name:  "consumerSecret,cs",
					Value: "",
					Usage: "set consumerSecretKey",
				},
				cli.StringFlag{
					Name:  "accsessToken,at",
					Value: "",
					Usage: "set accessToken",
				},
				cli.StringFlag{
					Name:  "accessSecret,as",
					Value: "",
					Usage: "set accessSecretKey",
				}},
			Action: ChangeConfig,
		},
	}
}
