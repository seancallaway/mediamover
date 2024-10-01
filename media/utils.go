package media

import (
	"path/filepath"
	"slices"
	"strings"
)

func getExtensionList() []string {
	return []string{
		".avi",
		".m4v",
		".mkv",
		".mp4",
		".mpeg",
		".mpg",
		".mov",
	}
}

func IsVideoFile(filename string) bool {
	extension := strings.ToLower(filepath.Ext(filename))

	return slices.Contains(getExtensionList(), extension)
}
