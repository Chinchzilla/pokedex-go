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

type LocationAreaResponse struct {
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []PokedexResults `json:"results"`
}

type PokemonEncounter struct {
	Pokemon PokedexResults `json:"pokemon"`
}

type ExploreResponse struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokeClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

const (
	baseURL              = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint = "location-area"
)

func NewClient(interval time.Duration) *PokeClient {
	return &PokeClient{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(interval),
	}
}

func doRequest(method, url string, body io.Reader, client *PokeClient) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch location area")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func GetLocationArea(url string, client *PokeClient) (LocationAreaResponse, error) {
	if strings.Trim(url, " ") == "" {
		url = baseURL + locationAreaEndpoint
	}
	var pokedexRes LocationAreaResponse
	data, is_cached := client.cache.Get(url)
	if !is_cached {
		dataRaw, err := doRequest(http.MethodGet, url, nil, client)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		client.cache.Add(url, dataRaw)
		data = dataRaw
	}

	err := json.Unmarshal(data, &pokedexRes)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return pokedexRes, nil
}

func ExploreLocation(client *PokeClient, args ...string) (ExploreResponse, error) {
	location := args[0]
	url := baseURL + locationAreaEndpoint + "/" + location
	var exploreLocationRes ExploreResponse
	data, is_cached := client.cache.Get(url)
	if !is_cached {
		dataRaw, err := doRequest(http.MethodGet, url, nil, client)
		if err != nil {
			return ExploreResponse{}, err
		}
		client.cache.Add(url, dataRaw)
		data = dataRaw
	}

	err := json.Unmarshal(data, &exploreLocationRes)
	if err != nil {
		return ExploreResponse{}, err
	}

	return exploreLocationRes, nil
}
