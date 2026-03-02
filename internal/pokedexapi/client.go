package pokedexapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type PokedexResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokedexResponse struct {
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []PokedexResults `json:"results"`
}

const (
	baseURL              = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint = "location-area"
)

func GetPokedexAPIClient() *http.Client {
	return &http.Client{}
}

func GetLocationArea(url string, client *http.Client) (PokedexResponse, error) {
	if strings.Trim(url, " ") == "" {
		url = baseURL + locationAreaEndpoint
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PokedexResponse{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return PokedexResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PokedexResponse{}, errors.New("failed to fetch location area")
	}

	decoder := json.NewDecoder(res.Body)

	var pokedexRes PokedexResponse
	err = decoder.Decode(&pokedexRes)
	if err != nil {
		return PokedexResponse{}, err
	}

	return pokedexRes, nil
}
