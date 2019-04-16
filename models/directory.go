package models

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

type Directory struct {
	Gitmojis []Gitmoji `json:"gitmojis"`
}

func SelectOption(gitmoji Gitmoji) string {
	return fmt.Sprintf("%s %s (%s)", gitmoji.Emoji, aurora.Bold(gitmoji.Code), gitmoji.Description)
}

func (d *Directory) SelectOptions() []string {
	options := []string{}

	for _, gitmoji := range d.Gitmojis {
		options = append(options, SelectOption(gitmoji))
	}

	return options
}

func (d *Directory) SelectedOption(selected string) *Gitmoji {
	for _, gitmoji := range d.Gitmojis {
		if selected == SelectOption(gitmoji) {
			return &gitmoji
		}
	}

	return nil
}
