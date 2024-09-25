package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showsCmd)
}

var showsCmd = &cobra.Command{
	Use:   "tvshows",
	Short: "Load TV shows from the given directory",
	Long:  "Load TV show episodes from the given directory to the tv_root specified in the config file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This doesn't do anything yet.")
	},
}
