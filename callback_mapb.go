package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(conf *config) error {
	result := ShallowLocation{}
	requestURL := conf.pastLocationsURL
	if requestURL == nil {
		return errors.New("you're on the first page")
	}
	cacheRes, ok := conf.cache.Get(*requestURL)
	if ok {
		err := json.Unmarshal(cacheRes, &result)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(*requestURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		conf.cache.Add(*requestURL, body)
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
