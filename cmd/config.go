package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

//UserConfig 設定ファイルのユーザ部分
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
	conf.Token.ConsumerKey = c.String("consumer")
	conf.Token.ConsumerSecret = c.String("consumerSecret")
	conf.Token.AccessToken = c.String("accsessToken")
	conf.Token.AccessSecret = c.String("accessSecret")

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
