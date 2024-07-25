package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eikonoklastess/pokedex-repl/internal/pki"
)

func startRepl(cfg *config) {
	fmt.Println("Welcome to Pokedex type help to display available commands")
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		var cliArg *string
		if len(words) > 1 {
			arg := words[1]
			cliArg = &arg
		}
		command, exist := getCommands()[commandName]
		if exist {
			err := command.callback(cfg, cliArg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type config struct {
	pokeapiClient    pki.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pki.RespShallowPokemon
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "fetch 20 location-area calls the next 20 upon calling it again",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "fetch the last 20 location-area its a way to go back from map",
			callback:    commandMapPrev,
		},
		"explore": {
			name:        "explore",
			description: "add to explore as an argument a location or area by their name of id to obtain a list of all possible pokemon in this area",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "enter a pokemon name to try and catch it if you succeed the pokemon is transferd to your pokedex",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon present in your pokedex shows their name, height, weight, base stats and their type",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "shows the pokemons in your pokedex",
			callback:    pokedex,
		},
	}

	return commands
}

func commandHelp(*config, *string) error {
	fmt.Println()
	fmt.Println("Here are the command available to use in pokedex: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(*config, *string) error {
	os.Exit(0)
	return nil
}
