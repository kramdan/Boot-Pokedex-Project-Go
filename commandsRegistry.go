package main

import pokecache "Boot-Pokedex-Project-Go/internal"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

var registry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exits the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays 20 map locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays last 20 map locations",
		callback:    commandMapB,
	},
	"explore": {
		name:        "explore",
		description: "Explores a specificed location and returns all pokemon that can be found",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Attempts to catch a specificed Pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Shows detailed information about a given Pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Shows the names of all pokemon you have caught",
		callback:    commandPokeDex,
	},
}

type config struct {
	regis            map[string]cliCommand
	cache            *pokecache.Cache
	pokeDex          *PokeDex
	nextLocationsURL *string
	pastLocationsURL *string
}
