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
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "timeline",
			Usage: "show timeline",
			// 引数に表示するタイムラインの数を指定できるようにする
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "trend",
			Usage: "get trends",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	// config := oauth1.NewConfig("NC73D8wtYilE8muXNocLzSSO2", "V8K7IyH024ggrliQOO4GDQu51Eg9oP2u91y9ZiitKjHx4b333l")
	// token := oauth1.NewToken("3214850788-4oCOxLrYB72Ngfq4sLXTVYnV4hyYsCXpm9izhmf", "8mk86yzdfXC1eWudqY0u7UFSq7QJ6Vr3OiwxMdjWOxtHl")
	// httpClient := config.Client(oauth1.NoContext, token)
}
