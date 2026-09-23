package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, s string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
