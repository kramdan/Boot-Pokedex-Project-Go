package main

import "fmt"

func commandPokeDex(conf *config, s string) error {
	for _, entry := range conf.pokeDex.collection {
		fmt.Printf(" - %s\n", entry.Name)
	}
	return nil
}
