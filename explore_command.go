package main

import (
	"errors"
	"fmt"

	"github.com/eikonoklastess/pokedex-repl/internal/pki"
)

func explore(cfg *config, loc *string) error {
	if loc == nil {
		return errors.New("no argument given to explore please give an area or location name or their id")
	}
	locInfoURL := *loc
	locInfoResp := pki.RespShallowLocInfo{}
	locInfoResp, err := cfg.pokeapiClient.ListLocInfo(&locInfoURL)
	if err != nil {
		return err
	}

	fmt.Println()
	for i, pokeInfo := range locInfoResp.PokemonEncounters {
		fmt.Printf(" %d. %s\n", i, pokeInfo.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
