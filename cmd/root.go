package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use:   "gotmoji",
	Short: "Gotmoji is a gitmoji port in go",
	Long: `A interactive command line tool for using emojis on commits
      without dependencies.
      Gotmoji: https://github.com/pkuebler/gotmoji/

      Gitmoji Overview: https://gitmoji.carloscuesta.me/
      NodeJS - Orginal: https://github.com/carloscuesta/gitmoji-cli`,
	Run: commitCmd,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
