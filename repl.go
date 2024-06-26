package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    callbackMap,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists all pokemons in give location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "Catch {pokemon_name}",
			description: "Attempt to cathc a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View caught pokemon details",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all pokemon in the pokedex",
			callback:    callbackPokedex,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
