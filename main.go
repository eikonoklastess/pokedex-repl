package main

import (
	"github.com/eikonoklastess/pokedex-repl/internal/pki"
	"time"
)

func main() {
	pokeClient := pki.NewClient(10*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pki.RespShallowPokemon),
	}

	startRepl(cfg)
}
