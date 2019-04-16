package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/pkuebler/gotmoji/gitmoji"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Print all emojis with description",
		Long:  `Print all gitmojis with description and shortcode`,
		Run:   listCmd,
	})
}

func listCmd(cmd *cobra.Command, args []string) {
	directory := gitmoji.FetchEmojis()

	if directory == nil {
		fmt.Println("Can't fetch the emoji directory.")
	}

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	for _, emoji := range directory.Gitmojis {
		fmt.Fprintln(w, fmt.Sprintf(" %s \t%s\t%s", emoji.Emoji, emoji.Description, emoji.Code))
	}

	w.Flush()
}
