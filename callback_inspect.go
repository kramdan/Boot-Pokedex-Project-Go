package main

import (
	"fmt"
)

func commandInspect(conf *config, s string) error {

	for _, entry := range conf.pokeDex.collection {
		if entry.Name == s {
			fmt.Printf("Name: %s\n", entry.Name)
			fmt.Printf("Height: %v\n", entry.Height)
			fmt.Printf("Weight: %v\n", entry.Weight)
			fmt.Printf("Stats:\n")
			for _, stat := range entry.Stats {
				fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Printf("Types:\n")
			for _, entry := range entry.Types {
				fmt.Printf("  - %s\n", entry.Type.Name)
			}
			return nil
		}
	}
	return fmt.Errorf("You have not caught this pokemon yet")
}
