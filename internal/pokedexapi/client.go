package pokedexapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Chinchzilla/pokedex-go/internal/pokecache"
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

type PokeClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

const (
	baseURL              = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint = "location-area"
)

func GetPokedexAPIClient(interval time.Duration) *PokeClient {
	return &PokeClient{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(interval),
	}
}

func GetLocationArea(url string, client *PokeClient) (PokedexResponse, error) {
	if strings.Trim(url, " ") == "" {
		url = baseURL + locationAreaEndpoint
	}
	var pokedexRes PokedexResponse
	data, is_cached := client.cache.Get(url)
	if !is_cached {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return PokedexResponse{}, err
		}

		res, err := client.httpClient.Do(req)
		if err != nil {
			return PokedexResponse{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return PokedexResponse{}, errors.New("failed to fetch location area")
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokedexResponse{}, err
		}
		client.cache.Add(url, data)
	}

	err := json.Unmarshal(data, &pokedexRes)
	if err != nil {
		return PokedexResponse{}, err
	}

	return pokedexRes, nil
}
