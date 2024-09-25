package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(moviesCmd)
}

var moviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "Load movies from the given directory",
	Long:  "Load movies from the given directory to the movie_root specified in the config file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This doesn't do anything yet.")
	},
}
