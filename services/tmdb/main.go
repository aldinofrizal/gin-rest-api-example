package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"overview"`
	ImageUrl    string `json:"poster_path"`
}

type MovieList struct {
	Results []Movie
}

type TmdbClient struct {
	ApiKey string
	Url    string
	Client *http.Client
}

func ImplTmdbClient() *TmdbClient {
	return &TmdbClient{
		ApiKey: os.Getenv("TMDB_KEY"),
		Url:    os.Getenv("TMDB_URL"),
		Client: &http.Client{},
	}
}

func (tc *TmdbClient) GetMovies() (MovieList, error) {
	params := url.Values{}
	params.Set("api_key", tc.ApiKey)
	params.Set("language", "en-US")
	params.Set("sort_by", "popularity.desc")
	params.Set("with_original_language", "ko")
	url := fmt.Sprintf("%s/discover/tv?%s", tc.Url, params.Encode())
	movieList := MovieList{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return movieList, err
	}

	resp, err := tc.Client.Do(req)
	if err != nil {
		return movieList, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&movieList)
	if err != nil {
		return movieList, err
	}

	return movieList, nil
}
