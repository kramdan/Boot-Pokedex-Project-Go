package main

import "fmt"

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!")
	fmt.Printf("Usage\n\n")

	for _, item := range conf.regis {
		fmt.Printf("%v: %v\n", item.name, item.description)
	}
	return nil
}
