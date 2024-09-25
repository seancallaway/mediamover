package cmd

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func isVideoFile(extention string) bool {
	videoFileExtensions := []string{
		"avi",
		"mkv",
		"mp4",
		"mpg",
	}

	return slices.Contains(videoFileExtensions, strings.ToLower(extention))
}

func getFileList(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(path + "does not exist")
		}
	}

	var files []string

	err = filepath.WalkDir(path, func(filename string, d fs.DirEntry, err error) error {
		if !d.IsDir() && isVideoFile(filepath.Ext(filename)) {
			files = append(files, filename)
		}
		return nil
	})
	if err != nil {
		return files, err
	}

	return files, nil
}
