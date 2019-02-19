package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

type Config struct {
	Token TokenConfig
}

type TokenConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

var conf Config

func init() {
	_, err := toml.DecodeFile("./config/token.toml", &conf)
	if err != nil {
		fmt.Println(err)
	}
}

func ChangeConfig(c *cli.Context) error {
	var (
		consumerKey    = c.String("consumer")
		consumerSecret = c.String("consumerSecret")
		accessToken    = c.String("accsessToken")
		accessSecret   = c.String("accessSecret")
	)

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic(fmt.Errorf("Config file does not have  enough authentication keys or secrets !!"))
	}

	conf.Token.ConsumerKey = consumerKey
	conf.Token.ConsumerSecret = consumerSecret
	conf.Token.AccessToken = accessToken
	conf.Token.AccessSecret = accessSecret

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf); err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("./config/token.toml")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(([]byte)(buf.String()))
	if err != nil {
		return err
	}
	return nil
}
