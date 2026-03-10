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

type PokeClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

const (
	baseURL              = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint = "location-area"
	pokemonEndpoint      = "pokemon"
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

func GetLocationArea(url string, client *PokeClient) (Location, error) {
	if strings.Trim(url, " ") == "" {
		url = baseURL + locationAreaEndpoint
	}
	var pokedexRes Location
	data, is_cached := client.cache.Get(url)
	if !is_cached {
		dataRaw, err := doRequest(http.MethodGet, url, nil, client)
		if err != nil {
			return Location{}, err
		}
		client.cache.Add(url, dataRaw)
		data = dataRaw
	}

	err := json.Unmarshal(data, &pokedexRes)
	if err != nil {
		return Location{}, err
	}

	return pokedexRes, nil
}

func ExploreLocation(client *PokeClient, args ...string) (Explore, error) {
	location := args[0]
	url := baseURL + locationAreaEndpoint + "/" + location
	var exploreLocationRes Explore
	data, is_cached := client.cache.Get(url)
	if !is_cached {
		dataRaw, err := doRequest(http.MethodGet, url, nil, client)
		if err != nil {
			return Explore{}, err
		}
		client.cache.Add(url, dataRaw)
		data = dataRaw
	}

	err := json.Unmarshal(data, &exploreLocationRes)
	if err != nil {
		return Explore{}, err
	}

	return exploreLocationRes, nil
}

func GetPokemon(client *PokeClient, args ...string) (Pokemon, error) {
	pokemon_name := args[0]
	url := baseURL + pokemonEndpoint + "/" + pokemon_name

	var getPokemonRes Pokemon
	data, is_cached := client.cache.Get(url)
	if !is_cached {
		dataRaw, err := doRequest(http.MethodGet, url, nil, client)
		if err != nil {
			return Pokemon{}, err
		}
		client.cache.Add(url, dataRaw)
		data = dataRaw
	}

	err := json.Unmarshal(data, &getPokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	return getPokemonRes, nil
}
