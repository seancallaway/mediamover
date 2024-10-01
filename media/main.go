package media

import (
	"path/filepath"
)

type Media struct {
	Title    string
	Year     string
	Season   string
	Episode  string
	AltTitle string
}

func ParseFilename(filename string, tvShow bool) (Media, error) {
	name := filepath.Base(filename)
	if tvShow {
		title, season, episode, err := parseTitleSeasonAndEpisode(name)
		if err != nil {
			return Media{}, err
		}
		return Media{Title: title, Season: season, Episode: episode}, nil
	} else {
		title, year, err := parseTitleAndYear(name)
		if err != nil {
			return Media{}, err
		}
		return Media{Title: title, Year: year}, nil
	}
}
