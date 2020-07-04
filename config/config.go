package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type (
	Config struct {
		Bot      Bot
		Server   Server
		Database Database
	}

	Bot struct {
		Id         string
		TopGGToken string
		DBLToken   string
	}

	Server struct {
		Bind string
	}

	Database struct {
		Uri string
	}
)

var (
	Conf Config
)

func readFile(name string) string {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func LoadConfig() {
	raw := readFile("config.toml")
	_, err := toml.Decode(raw, &Conf)
	if err != nil {
		panic(err)
	}
}
