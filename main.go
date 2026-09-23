package main

import (
	pokecache "Boot-Pokedex-Project-Go/internal"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	runningConf := config{
		regis: registry,
		cache: pokecache.NewCache(10 * time.Second),
	}
	reader := bufio.NewScanner(os.Stdin)
	reader.Err()
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		text := reader.Text()
		cleanText := cleanInput(text)
		firstWord := cleanText[0]
		secondWord := ""
		if len(cleanText) > 1 {
			secondWord = cleanText[1]
		}
		command, exists := runningConf.regis[firstWord]
		if exists {
			err := command.callback(&runningConf, secondWord)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("")
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
