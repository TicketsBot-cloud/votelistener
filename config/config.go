package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type(
	Config struct {
		Bot Bot
		Server Server
		Database Database
	}

	Bot struct {
		Id string
		Token string
	}

	Server struct {
		Bind string
	}

	Database struct {
		Host string
		Port uint16
		Username string
		Password string
		Database string
	}
)

var(
	Conf Config
)

func readFile(name string) string {
	contents, err := ioutil.ReadFile(name); if err != nil {
		panic(err)
	}

	return string(contents)
}

func LoadConfig() {
	raw := readFile("config.toml")
	_, err := toml.Decode(raw, &Conf); if err != nil {
		panic(err)
	}
}
