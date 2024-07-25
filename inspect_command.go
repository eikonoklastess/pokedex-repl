package main

import (
	"errors"
	"fmt"
)

func inspect(cfg *config, cliArg *string) error {
	if cliArg == nil {
		return errors.New("no pokemon given to inspect")
	} else if _, ok := cfg.pokedex[*cliArg]; !ok {
		return errors.New("this isnt a pokemon present in your pokedex")
	}

	pokemonInfo := cfg.pokedex[*cliArg]

	fmt.Println()
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemonInfo.Name, pokemonInfo.Height, pokemonInfo.Weight)
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println()

	return nil
}
