package main

import "fmt"

func pokedex(cfg *config, cliArg *string) error {
	for _, v := range cfg.pokedex {
		fmt.Println("your pokemons:")
		fmt.Printf("  - %s", v.Name)
	}
	fmt.Println()
	return nil
}
