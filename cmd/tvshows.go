package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showsCmd)
}

var showsCmd = &cobra.Command{
	Use:   "tvshows [path to load]",
	Short: "Load TV shows from the given directory",
	Long:  "Load TV show episodes from the given directory to the tv_root specified in the config file.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get list of files to load.
		fileList, err := getFileList(args[0])
		if err != nil {
			return err
		}
		if len(fileList) == 0 {
			fmt.Println("No files to load in", args[0])
			return nil
		}

		// Load files
		fmt.Println("Will load the following files:\n")
		for _, file := range fileList {
			fmt.Println("-", file)
		}
		return nil
	},
}
