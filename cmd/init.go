package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var App cli.App

func init() {

	App.Commands = []cli.Command{
		{
			Name:  "tweet",
			Usage: "tweet with text",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args())
				return nil
			},
		},
		{
			Name:  "timeline",
			Usage: "show timeline",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "num,n",
					Value: "20",
					Usage: "set the limit of timelines",
				}},
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "trend",
			Usage: "get trends",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
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
					Name:  "consumersecret,cs",
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
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
	}

	// config := oauth1.NewConfig("NC73D8wtYilE8muXNocLzSSO2", "V8K7IyH024ggrliQOO4GDQu51Eg9oP2u91y9ZiitKjHx4b333l")
	// token := oauth1.NewToken("3214850788-4oCOxLrYB72Ngfq4sLXTVYnV4hyYsCXpm9izhmf", "8mk86yzdfXC1eWudqY0u7UFSq7QJ6Vr3OiwxMdjWOxtHl")
	// httpClient := config.Client(oauth1.NoContext, token)
}
