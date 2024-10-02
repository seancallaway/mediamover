package media

import (
	"testing"
)

func TestMovieRegex(t *testing.T) {
	var filenames = []string{
		"Mission.Impossible.3.Special.Edition.2011.mp4",
		"Raiders of the Lost Ark (1981).mkv",
		"For.a.Few.Dollars.More.1965.avi",
		"Apartment.7A.2024.1080p.WEBRip.x264.AAC5.1-[YTS.MX].mp4",
		"Fly.Me.To.The.Moon.2024.1080p.WEBRip.x265.10bit.AAC5.1-[YTS.MX].mp4",
	}

	var expectedResults = []struct {
		Title string
		Year  string
	}{
		{"Mission Impossible 3", "2011"},
		{"Raiders of the Lost Ark", "1981"},
		{"For a Few Dollars More", "1965"},
		{"Apartment 7A", "2024"},
		{"Fly Me To The Moon", "2024"},
	}

	for i, filename := range filenames {
		title, year, err := parseTitleAndYear(filename)
		if err != nil {
			t.Errorf("Failed to parse title and year from filename: %q", err)
		}

		if expectedResults[i].Title != title || expectedResults[i].Year != year {
			t.Errorf("Parse result was incorrect; got %s %s, want %s %s", title, year, expectedResults[i].Title, expectedResults[i].Year)
		}
	}
}

func TestShowRegex(t *testing.T) {
	var filenames = []string{
		"SEAL.Team.S07E05.1080p.HEVC.x265-MeGusta.mkv",
		"Futurama.S09E10.XviD-AFG[EZTVx.to].avi",
		"Only.Murders.in.the.Building.S04E05.Adaptation.1080p.HEVC.x265-MeGusta [TD].mkv",
	}

	var expectedResults = []struct {
		Title   string
		Season  string
		Episode string
	}{
		{"SEAL Team", "07", "05"},
		{"Futurama", "09", "10"},
		{"Only Murders in the Building", "04", "05"},
	}

	for i, filename := range filenames {
		title, season, episode, err := parseTitleSeasonAndEpisode(filename)
		if err != nil {
			t.Errorf("Failed to parse title, season, and episode from filename: %q", err)
		}

		if expectedResults[i].Title != title || expectedResults[i].Season != season || expectedResults[i].Episode != episode {
			t.Errorf("Parse result was incorrect; got %s %s %s, want %s %s %s", title, season, episode, expectedResults[i].Title, expectedResults[i].Season, expectedResults[i].Episode)
		}
	}
}
