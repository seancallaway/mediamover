package cmd

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/seancallaway/mediamover/media"
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
