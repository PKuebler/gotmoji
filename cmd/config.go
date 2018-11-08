package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Gitmojis []Gitmoji
}

type Gitmoji struct {
	Emoji       string
	Entity      string
	Code        string
	Description string
	Name        string
}

func LoadConfig() Config {
	var config Config

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
