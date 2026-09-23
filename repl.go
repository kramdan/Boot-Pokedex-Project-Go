package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var output []string
	textLower := strings.ToLower(text)
	textSplit := strings.Fields(textLower)
	for _, s := range textSplit {
		strings.ReplaceAll(s, " ", "")
		if s != "" {
			output = append(output, s)
		}

	}
	return output
}
