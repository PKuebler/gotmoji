package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of gotmoji",
		Long:  `All software has versions. This is gotmoji's`,
		Run:   versionCmd,
	})
}

func versionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Gotmoji v0.1")
}
