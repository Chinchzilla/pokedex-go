package main

import (
	"time"

	"github.com/Chinchzilla/pokedex-go/internal/pokedexapi"
)

const cacheInterval = 5 * time.Second

func main() {
	cfg := &Config{
		pokedexAPIClient: pokedexapi.NewClient(cacheInterval),
	}

	startRepl(cfg)
}
