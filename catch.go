package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func catch(cfg *config, pokemon *string) error {
	if pokemon == nil {
		return errors.New("no argument given to explore please give an area or location name or their id")
	} else if _, ok := cfg.pokedex[*pokemon]; ok {
		return errors.New("this pokemon is already caught check out your pokedex")
	}

	pokemonInfo, err := cfg.pokeapiClient.PokeInfo(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Trowing pokeball at %s\n", pokemonInfo.Name)
	fmt.Print("Wait ")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(". ")
	}
	//chance := int(math.Round(float64(pokemonInfo.BaseExperience) / 100.0))
	if rand.Intn(2) == 1 {
		fmt.Printf("%s was caught\n", pokemonInfo.Name)
		cfg.pokedex[pokemonInfo.Name] = pokemonInfo
	} else {
		fmt.Printf("%s escaped\n", pokemonInfo.Name)
	}

	return nil
}
