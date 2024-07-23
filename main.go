package main

import (
	"pokedex-repl/internal/pkcache"
	"pokedex-repl/internal/pki"
	"time"
)

func main() {
	pokeClient := pki.NewClient(5 * time.Second)
	cache := pkcache.NewCache(1 * time.Minute)
	cfg := &config{
		pokeapiClient:    pokeClient,
		cache:            &cache,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}
	startRepl(cfg)
}
