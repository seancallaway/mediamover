package media

import (
	"testing"
)

func TestVideoExtension(t *testing.T) {
	videoFiles := []string{
		"test.mp4",
		"fake_file123.mpg",
		"This.should.work.mkv",
		"windows key.avi",
	}

	nonVideoFiles := []string{
		"mediamover.exe",
		"binary",
		"testing123.docx",
		"this.isn't.an.avi.exe",
	}

	for _, file := range videoFiles {
		if !IsVideoFile(file) {
			t.Errorf("Expected %s to be recognized as video file.", file)
		}
	}

	for _, file := range nonVideoFiles {
		if IsVideoFile(file) {
			t.Errorf("Expected %s to NOT be recognized as video file.", file)
		}
	}
}
