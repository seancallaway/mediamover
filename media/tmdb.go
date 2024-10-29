package media

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
)

var genreMap = map[int]string{
	12:    "Adventure",
	14:    "Fantasy",
	16:    "Animation",
	18:    "Drama",
	27:    "Horror",
	28:    "Action",
	35:    "Comedy",
	36:    "History",
	37:    "Western",
	53:    "Thriller",
	80:    "Crime",
	99:    "Documentary",
	878:   "Science Fiction",
	9648:  "Mystery",
	10402: "Music",
	10749: "Romance",
	10751: "Family",
	10752: "War",
	10770: "TV Movie",
}

type TMDBMovie struct {
	Adult            bool
	BackdropPath     string `json:"backdrop_path"`
	GenreIds         []int  `json:"genre_ids"`
	Id               int
	OriginalLanguage string `json:"original_language"`
	OriginalTitle    string `json:"original_title"`
	Overview         string
	Popularity       float32
	PosterPath       string `json:"poster_path"`
	ReleaseDate      string `json:"release_date"`
	Title            string
	Video            bool
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type TMDBResponse struct {
	Page    int
	Results []TMDBMovie
}

// getPrimaryGenre takes a movie title and a year then returns that movie's
// primary genre (should the movie be found).
func getPrimaryGenre(title string, year string) (string, error) {
	// Build Search URL
	u := url.URL{
		Scheme: "https",
		Host:   "api.themoviedb.org",
		Path:   "3/search/movie",
	}
	q := u.Query()
	q.Set("query", title)
	q.Set("primary_release_year", year)
	q.Set("api_key", viper.GetString("default.api_key"))
	u.RawQuery = q.Encode()

	// Build request from URL
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("accept", "application/json")

	// Make request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Check if API Auth failed (or other issues)
	if res.StatusCode != http.StatusOK {
		var errMsg string
		if res.StatusCode == http.StatusUnauthorized {
			errMsg = "cannot access TMDB API; ensure your API key is set correctly"
		} else {
			errMsg = "received unexpected status code from TMDB API: " + strconv.Itoa(res.StatusCode)
		}

		return "", errors.New(errMsg)
	}

	// Get response
	var movies TMDBResponse

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal(body, &movies)

	if len(movies.Results) == 0 {
		return "", errors.New("no movie found")
	}

	if len(movies.Results[0].GenreIds) == 0 {
		return "", errors.New("no genre found")
	}

	genreId := movies.Results[0].GenreIds[0]

	return genreMap[genreId], nil
}
