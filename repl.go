package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex-repl/internal/pkcache"
	"pokedex-repl/internal/pki"
	"strings"
)

type config struct {
	pokeapiClient    pki.Client
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pkcache.Cache
}

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
		command, exist := getCommands()[commandName]
		if exist {
			err := command.callback(cfg)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
	return commands
}

func commandHelp(*config) error {
	fmt.Println()
	fmt.Println("Here are the command available to use in pokedex: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(*config) error {
	os.Exit(0)
	return nil
}
