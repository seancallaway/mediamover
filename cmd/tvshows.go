package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/seancallaway/mediamover/media"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		for _, file := range fileList {
			data, err := media.ParseFilename(file, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to identify %s\n", file)
				continue
			}

			destinationPath, err := getDestinationPath(data, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to determine destination directory: %q\n", err)
			}
			// Ensure directory exists
			if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
				fmt.Fprintf(os.Stderr, "Unable to create directory %s: %q\n", destinationPath, err)
			}

			destinationFilename := data.Title + " S" + data.Season + "E" + data.Episode + filepath.Ext(file)
			finalFile := filepath.Join(destinationPath, destinationFilename)

			// TODO: Move or copy the file, depending on flag.
			if _, err := copyFile(file, finalFile); err != nil {
				fmt.Fprintf(os.Stderr, "Unable to write file %s: %q\n", finalFile, err)
			} else {
				if viper.GetBool("verbose") {
					fmt.Println("Wrote", finalFile)
				}
			}
		}
		return nil
	},
}
