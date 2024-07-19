package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
			err := command.callback()
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
	callback    func() error
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
	}
	return commands
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Here are the command available to use in pokedex: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
