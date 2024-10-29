package media

import (
	"errors"
	"strings"

	"github.com/dlclark/regexp2"
)

// getTitleYearRegexPatterns returns a slice of strings containing .NET-style regex patterns
// to be used by the dlclark/regexp2 library, which supports negative lookahead. These patterns
// should be ordered "best case" to "worst case".
func getTitleYearRegexPatterns() []string {
	return []string{
		// Match files with editions, like Mission.Impossible.3.Special.Edition.2011.mp4
		`^(?'title'(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*\(?\b(?'edition'(((Extended.|Ultimate.)?(Director.?s|Collector.?s|Theatrical|Anniversary|The.Uncut|Ultimate|Final(?=(.(Cut|Edition|Version)))|Extended|Rogue|Special|Despecialized|\d{2,3}(th)?.Anniversary)(.(Cut|Edition|Version))?(.(Extended|Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit))?|((Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit|Edition|Restored|((2|3|4)in1))))))\b\)?.{1,3}(?'year'(1(8|9)|20)\d{2}(?!p|i|\d+|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
		// Match files in the target movie format, like Raiders of the Lost Ark (1981).mkv
		`^(?'title'(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*\((?'year'(1(8|9)|20)\d{2}(?!p|i|(1(8|9)|20)\d{2}|\]|\W(1(8|9)|20)\d{2})))+`,
		// Match files in the standard donload format, like For.a.Few.Dollars.More.1965.avi
		`^(?<title>(?![([]).+?)?(?:(?:[-_\W](?<![)[!]))*(?<year>(1(8|9)|20)\d{2}(?!p|i|(1(8|9)|20)\d{2}|\]|\W(1(8|9)|20)\d{2})))+(\W+|_|$)(?!\\)`,
		// Getting desperate: match years wrapped in square brackets
		`^(?'title'(?![([]).+?)?(?:(?:[-_\W](?<![)!]))*(?'year'(1(8|9)|20)\d{2}(?!p|i|\d+|\W\d+)))+(\W+|_|$)(?!\\)`,
		// Last resort
		`^(?'title'.+?)?(?:(?:[-_\W](?<![)[!]))*(?'year'(1(8|9)|20)\d{2}(?!p|i|\d+|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
	}
}

func parseTitleAndYear(filename string) (string, string, error) {
	for _, expression := range getTitleYearRegexPatterns() {
		re := regexp2.MustCompile(expression, regexp2.IgnoreCase)
		if m, _ := re.FindStringMatch(filename); m != nil {
			if title := m.GroupByName("title"); title != nil {
				cleanedTitle := strings.ReplaceAll(title.String(), ".", " ")
				if len(cleanedTitle) == 0 {
					// Bad match found for title group
					continue
				}

				if year := m.GroupByName("year"); year != nil {
					return cleanedTitle, year.String(), nil
				} else {
					return cleanedTitle, "", nil
				}
			}
		}
	}

	// TODO: Try finding the title using codec or resolution position

	return "", "", errors.New("unable to parse title and year from filename")
}

func parseTitleSeasonAndEpisode(filename string) (string, string, string, error) {
	re := regexp2.MustCompile(`^(?<title>.+?)[-. ]{0,3}s?(?<season>\d?\d)[ex](?<episode>\d\d)[-. ]{0,3}(?<episode_title>.*?)[-. ]?(?:(?=pulcione|eng|ita|\w+Mux|\w+dl|\d+p|XviD|NovaRip).+)?\.([\w]{2,3})$`, regexp2.IgnoreCase)
	if m, _ := re.FindStringMatch(filename); m != nil {
		if title := m.GroupByName("title"); title != nil {
			cleanedTitle := strings.ReplaceAll(title.String(), ".", " ")
			if len(cleanedTitle) == 0 {
				return "", "", "", errors.New("unable to identify show title")
			}

			if season := m.GroupByName("season"); season != nil {
				if episode := m.GroupByName("episode"); episode != nil {
					return cleanedTitle, season.String(), episode.String(), nil
				}
			}
		}
	}

	return "", "", "", errors.New("unable to identify show")
}
