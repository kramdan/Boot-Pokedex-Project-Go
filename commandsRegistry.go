package main

import pokecache "Boot-Pokedex-Project-Go/internal"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
}

type config struct {
	regis            map[string]cliCommand
	cache            *pokecache.Cache
	nextLocationsURL *string
	pastLocationsURL *string
}
