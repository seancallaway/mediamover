package media

import (
	"path/filepath"
	"slices"
	"strings"
)

// getExtensionList returns a list of file extensions for video and subtitle files.
func getExtensionList() []string {
	return []string{
		".avi",
		".m4v",
		".mkv",
		".mp4",
		".mpeg",
		".mpg",
		".mov",
		".srt",
	}
}

func IsVideoFile(filename string) bool {
	extension := strings.ToLower(filepath.Ext(filename))

	return slices.Contains(getExtensionList(), extension)
}
