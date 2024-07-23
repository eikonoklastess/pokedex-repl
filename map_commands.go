// package for pokeAPI interaction functions
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	pkURL = "https://pokeapi.co/api/v2/location-area?"
)

func commandMapNext(cfg *config) error {
	locationsCached, ok := cfg.cache.Get(cfg.nextLocationsURL)
	if ok {
		locationsCachedStr := strings.Split(string(locationsCached), "\n")
		fmt.Println("USED CACHE")
		for _, loc := range locationsCachedStr {
			fmt.Println(loc)
		}
		fmt.Println()

		cfg.prevLocationsURL = cfg.nextLocationsURL
		urlParts := strings.Split(*cfg.nextLocationsURL, "=")
		index, _ := strconv.Atoi(urlParts[1])
		urlParts[1] = strconv.Itoa(index + 20)
		*cfg.nextLocationsURL = strings.Join(urlParts, "=")

		return nil
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	locationsBuffer := []byte{}
	fmt.Println()
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
		nameSplit := []byte(loc.Name)
		for _, nameByte := range nameSplit {
			locationsBuffer = append(locationsBuffer, nameByte)
		}
		locationsBuffer = append(locationsBuffer, '\n')
	}
	fmt.Println()

	if cfg.nextLocationsURL == nil {
		cfg.cache.Add("https://pokeapi.co/api/v2/location-area?offset=0&limit=20", locationsBuffer)
	} else {
		cfg.cache.Add(*cfg.nextLocationsURL, locationsBuffer)
	}

	location := func(p *string) string {
		if p == nil {
			return "no location"
		} else {
			return *p
		}
	}

	fmt.Printf("current location: %v\n\n", location(cfg.nextLocationsURL))
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	fmt.Printf("next location: %v\n", location(cfg.nextLocationsURL))
	fmt.Printf("prev location: %v\n", location(cfg.prevLocationsURL))

	return nil
}

func commandMapPrev(cfg *config) error {

	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsCached, ok := cfg.cache.Get(cfg.prevLocationsURL)
	locationsCachedStr := strings.Split(string(locationsCached), "\n")
	if ok {
		fmt.Println("USED CACHE")
		for _, loc := range locationsCachedStr {
			fmt.Println(loc)
		}
		fmt.Println()

		locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		cfg.nextLocationsURL = locationsResp.Next
		cfg.prevLocationsURL = locationsResp.Previous

		return nil
	}

	fmt.Println("you should not be here!")

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	locationsBuffer := []byte{}
	fmt.Println()
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
		nameSplit := []byte(loc.Name)
		for _, nameByte := range nameSplit {
			locationsBuffer = append(locationsBuffer, nameByte)
		}
		locationsBuffer = append(locationsBuffer, '\n')
	}
	fmt.Println()

	return nil
}
