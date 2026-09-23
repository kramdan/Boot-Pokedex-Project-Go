package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

func commandCatch(conf *config, s string) error {
	if s == "" {
		return fmt.Errorf("No Pokemon given to catch {catch <pokemon name>}")
	}
	var result Pokemon
	requestURL := "https://pokeapi.co/api/v2/pokemon/" + s

	cacheRes, ok := conf.cache.Get(requestURL)
	if ok {
		err := json.Unmarshal(cacheRes, &result)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(requestURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		conf.cache.Add(requestURL, body)
		err = json.Unmarshal(body, &result)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", result.Name)
	chance := rand.IntN(result.BaseExperience + (result.BaseExperience / 2))
	if result.BaseExperience < chance {
		fmt.Printf("%s escaped!\n", result.Name)
	} else {
		fmt.Printf("%s was caught!\n", result.Name)
		conf.pokeDex.Add(requestURL, result)
	}
	return nil
}
