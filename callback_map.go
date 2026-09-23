package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(conf *config) error {
	result := ShallowLocation{}
	requestURL := "https://pokeapi.co/api/v2/location-area/"
	if conf.nextLocationsURL != nil {
		requestURL = *conf.nextLocationsURL
	}
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

	conf.nextLocationsURL = result.Next
	conf.pastLocationsURL = result.Previous
	for _, local := range result.Results {
		fmt.Printf("%v\n", local.Name)
	}
	return nil
}

type ShallowLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
