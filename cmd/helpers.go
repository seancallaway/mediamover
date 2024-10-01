package cmd

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/seancallaway/mediamover/media"
	"github.com/spf13/viper"
)

func getFileList(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return nil, errors.New(path + " does not exist")
	}

	var files []string

	err = filepath.WalkDir(path, func(filename string, d fs.DirEntry, err error) error {
		if !d.IsDir() && media.IsVideoFile(filename) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func copyFile(src string, dest string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, errors.New(src + " is not a regular file")
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	numBytes, err := io.Copy(destination, source)
	return numBytes, err
}

// getDestinationPath returns the path to which the video file should be copied.
// In the future, this will support user configuration, but for now is static.
func getDestinationPath(data media.Media, isTV bool) (string, error) {
	if isTV {
		fileRoot := viper.GetString("default.tv_root")
		if strings.HasPrefix(fileRoot, "~/") {
			home, _ := os.UserHomeDir()
			fileRoot = filepath.Join(home, fileRoot[2:])
		}

		showPath := filepath.Join(fileRoot, data.Title)
		return filepath.Join(showPath, "Season "+data.Season), nil
	} else {
		fileRoot := viper.GetString("default.movie_root")
		if strings.HasPrefix(fileRoot, "~/") {
			home, _ := os.UserHomeDir()
			fileRoot = filepath.Join(home, fileRoot[2:])
		}

		return filepath.Join(fileRoot, data.Genre), nil
	}
}
