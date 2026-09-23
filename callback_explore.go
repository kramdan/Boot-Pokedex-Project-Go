package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(conf *config, s string) error {
	if s == "" {
		return fmt.Errorf("No location given to explore {explore <location name>}")
	}
	result := Location{}
	requestURL := "https://pokeapi.co/api/v2/location-area/" + s

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
	for _, local := range result.PokemonEncounters {
		fmt.Printf(" - %v\n", local.Pokemon.Name)
	}
	return nil
}

type Location struct {
	Id                     int    `json:"id"`
	Name                   string `json:"name"`
	GameIndex              int    `json:"game_index"`
	Encounter_method_rates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					Url  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
