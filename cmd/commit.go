package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/pkuebler/gotmoji/gitmoji"
	"github.com/pkuebler/gotmoji/models"
)

var useUTF8 bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&useUTF8, "utf8", "u", false, "utf8 emoji (default :code:)")
	rootCmd.AddCommand(&cobra.Command{
		Use:   "commit",
		Short: "commit",
		Long:  `commit`,
		Run:   commitCmd,
	})
}

func commitCmd(cmd *cobra.Command, args []string) {
	directory := gitmoji.FetchEmojis()

	if directory == nil {
		fmt.Println("Can't fetch the emoji directory.")
		return
	}

	gitmojiOptions := directory.SelectOptions()

	qs := []*survey.Question{
		{
			Name: "gitmoji",
			Prompt: &survey.Select{
				Message: "Choose a gitmoji",
				Options: gitmojiOptions,
				FilterFn: func(filter string, options []string) (filtered []string) {
					result := survey.DefaultFilterFn(filter, options)
					for _, v := range result {
						if len(v) >= 5 {
							filtered = append(filtered, v)
						}
					}
					return
				},
			},
		},
		{
			Name: "title",
			Prompt: &survey.Input{
				Message: "Enter the commit title",
			},
		},
		{
			Name: "message",
			Prompt: &survey.Multiline{
				Message: "Enter the commit message",
			},
		},
		{
			Name: "issue",
			Prompt: &survey.Input{
				Message: "Issue / PR reference",
			},
		},
	}

	// the answers will be written to this struct
	answers := models.Commit{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	answers.Gitmoji = directory.SelectedOption(answers.Selected)

	gitmoji.Commit(answers, useUTF8)
}
