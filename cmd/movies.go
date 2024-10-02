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
	rootCmd.AddCommand(moviesCmd)
}

var moviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "Load movies from the given directory",
	Long:  "Load movies from the given directory to the movie_root specified in the config file.",
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
			data, err := media.ParseFilename(file, false)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to identify %s", file)
				continue
			}

			destinationPath, err := getDestinationPath(data, false)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to determine destination directory: %q", err)
			}
			// Ensure directory exists
			if os.MkdirAll(destinationPath, os.ModePerm); err != nil {
				fmt.Fprintf(os.Stderr, "Unable to create directory %s: %q", destinationPath, err)
			}

			destinationFilename := data.Title + " (" + data.Year + ")" + filepath.Ext(file)
			finalFile := filepath.Join(destinationPath, destinationFilename)

			// TODO: Move or copy the file, depending on flag.
			if _, err := copyFile(file, finalFile); err != nil {
				fmt.Fprintf(os.Stderr, "Unable to write file %s: %q", finalFile, err)
			} else {
				if viper.GetBool("verbose") {
					fmt.Println("Wrote", finalFile)
				}
			}
		}
		return nil
	},
}
