package media

import (
	"errors"
	"fmt"
	"math"
	"strconv"
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

// getLanguages returns a slice of strings containing .NET-style regex patterns for languages
// commonly found in media filenames.
func getLanguagePatterns() []string {
	return []string{
		`\bArabic`,
		`\bBengali`,
		`\bBrazilian`,
		`\bBulgarian`,
		`\bCantonese`,
		`\bCatalan`,
		`\bChinese`,
		`\bCzech`,
		`\bDanish`,
		`\bDutch`,
		`\bEnglish`,
		`\bEstonian`,
		`\bFinnish`,
		`\bFlemish`,
		`\bFrench`,
		`\bGerman`,
		`\bGreek`,
		`\bHebrew`,
		`\bHindi`,
		`\bHungarian`,
		`\bIcelandic`,
		`\bItalian`,
		`\bJapanese`,
		`\bKorean`,
		`\bLatvian`,
		`\bLithuanian`,
		`\bMandarin`,
		`\bNordic`,
		`\bNorwegian`,
		`\bPersian`,
		`\bPolish`,
		`\bPortuguese`,
		`\bRomanian`,
		`\bRussian`,
		`\bSerbian`,
		`\bSlovak`,
		`\bSpanish`,
		`\bSwedish`,
		`\bTamil`,
		`\bThai`,
		`\bTurkish`,
		`\bUkrainian`,
		`\bVietnamese`,
	}
}

func cleanReleaseTitle(dirtyTitle string) string {
	fmt.Println("Dirty Title: ", dirtyTitle)
	if len(dirtyTitle) == 0 || dirtyTitle == "(" {
		return ""
	}

	trimmedTitle := strings.ReplaceAll(dirtyTitle, "_", " ")
	// Remove common sources
	re := regexp2.MustCompile(`\b(Bluray|(dvdr?|BD)rip|HDTV|HDRip|TS|R5|CAM|SCR|(WEB|DVD)?.?SCREENER|DiVX|xvid|web-?dl)\b`, regexp2.IgnoreCase)
	trimmedTitle, err := re.Replace(trimmedTitle, "", -1, -1)
	if err != nil {
		fmt.Printf("Error encountered while removing common sources: %q", err)
	}
	// Remove webdl sources
	re = regexp2.MustCompile(`\b(?<webdl>WEB[-_. ]DL|HDRIP|WEBDL|WEB-DLMux|NF|APTV|NETFLIX|NetflixU?HD|DSNY|DSNP|HMAX|AMZN|AmazonHD|iTunesHD|MaxdomeHD|WebHD|WEB$|[. ]WEB[. ](?:[xh]26[45]|DD5[. ]1)|\d+0p[. ]WEB[. ]|\b\s\/\sWEB\s\/\s\b|AMZN[. ]WEB[. ])\b`, regexp2.IgnoreCase)
	trimmedTitle, err = re.Replace(trimmedTitle, "", -1, -1)
	if err != nil {
		fmt.Printf("Error encountered while removing webdl sources: %q", err)
	}
	// Remove edition
	re = regexp2.MustCompile(`\b((Extended.|Ultimate.)?(Director.?s|Collector.?s|Theatrical|Anniversary|The.Uncut|DC|Ultimate|Final(?=(.(Cut|Edition|Version)))|Extended|Special|Despecialized|unrated|\d{2,3}(th)?.Anniversary)(.(Cut|Edition|Version))?(.(Extended|Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit))?|((Uncensored|Remastered|Unrated|Uncut|IMAX|Fan.?Edit|Edition|Restored|((2|3|4)in1)))){1,3}`, regexp2.IgnoreCase)
	trimmedTitle, err = re.Replace(trimmedTitle, "", -1, -1)
	if err != nil {
		fmt.Printf("Error encountered while removing edition: %q", err)
	}
	// Remove subtitle languages
	re = regexp2.MustCompile(`\b(TRUE.?FRENCH|videomann|SUBFRENCH|PLDUB|MULTI)`, regexp2.IgnoreCase)
	trimmedTitle, err = re.Replace(trimmedTitle, "", -1, -1)
	if err != nil {
		fmt.Printf("Error encountered while removing subtitle language: %q", err)
	}
	// Remove PROPER, etc.
	re = regexp2.MustCompile(`\b(PROPER|REAL|READ.NFO)`, regexp2.IgnoreCase)
	trimmedTitle, err = re.Replace(trimmedTitle, "", -1, -1)
	if err != nil {
		fmt.Printf("Error encountered while removing PROPER, etc.: %q", err)
	}

	// Remove languages
	for _, langPattern := range getLanguagePatterns() {
		re = regexp2.MustCompile(langPattern, regexp2.IgnoreCase)
		trimmedTitle, err = re.Replace(trimmedTitle, "", -1, -1)
		if err != nil {
			fmt.Printf("Error encountered while removing languages: %q", err)
		}
	}

	trimmedTitle = strings.TrimSpace(trimmedTitle)

	// Find gap formed by removing items
	trimmedTitle = strings.Split(trimmedTitle, "  ")[0]
	trimmedTitle = strings.Split(trimmedTitle, "..")[0]

	parts := strings.Split(trimmedTitle, ".")
	var result string = ""
	var previousAcronym bool = false
	var nextPart string = ""

	for i, part := range parts {
		if len(parts) >= i+2 {
			nextPart = parts[i+1]
		}

		if len(part) == 1 && strings.ToLower(part) != "a" {
			partAsFloat, err := strconv.ParseFloat(part, 64)
			if err == nil && math.IsNaN(partAsFloat) {
				result += part + "."
			}
		} else if strings.ToLower(part) == "a" && (previousAcronym || len(nextPart) == 1) {
			result += part + "."
			previousAcronym = true
		} else {
			if previousAcronym {
				result += " "
				previousAcronym = false
			}

			result += part + " "
		}
	}

	return strings.TrimSpace(trimmedTitle)
}

func parseTitleAndYear(filename string) (string, string, error) {
	for _, expression := range getTitleYearRegexPatterns() {
		re := regexp2.MustCompile(expression, regexp2.IgnoreCase)
		if m, _ := re.FindStringMatch(filename); m != nil {
			if title := m.GroupByName("title"); title != nil {
				cleanedTitle := strings.ReplaceAll(title.String(), ".", " ")
				//cleaned_title := cleanReleaseTitle(title.String())
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
