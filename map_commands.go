// package for pokeAPI interaction functions
package main

import (
	"errors"
	"fmt"
	"github.com/eikonoklastess/pokedex-repl/internal/pki"
	"strings"
)

func commandMapNext(cfg *config, cliArg *string) error {
	var locationsResp = pki.RespShallowLocations{}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locationsResp.Previous
	cfg.nextLocationsURL = locationsResp.Next

	index := strings.Split(strings.Split(*locationsResp.Next, "=")[1], "&")[0]

	fmt.Printf("[%v/%v]\n\n", index, locationsResp.Count)
	for _, result := range locationsResp.Results {
		fmt.Printf(" - %v\n", result.Name)
	}
	fmt.Println()

	return nil
}

func commandMapPrev(cfg *config, cliArg *string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	var locationsResp = pki.RespShallowLocations{}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locationsResp.Previous
	cfg.nextLocationsURL = locationsResp.Next

	index := strings.Split(strings.Split(*locationsResp.Next, "=")[1], "&")[0]

	fmt.Printf("[%v/%v]\n\n", index, locationsResp.Count)
	for _, result := range locationsResp.Results {
		fmt.Printf(" - %v\n", result.Name)
	}
	fmt.Println()

	return nil
}
