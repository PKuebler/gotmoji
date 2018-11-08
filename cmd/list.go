package cmd

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

func List() {
	config := LoadConfig()

	for _, emoji := range config.Gitmojis {
		fmt.Println(emoji.Emoji, "-", Magenta(emoji.Code), "-", emoji.Description)
	}
}
