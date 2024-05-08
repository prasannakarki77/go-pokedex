package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapCommand() error {
	res, _ := http.Get("https://pokeapi.co/api/v2/location-area")

	resBody, _ := io.ReadAll(res.Body)
	locationResult := LocationResult{}

	errorr := json.Unmarshal(resBody, &locationResult)

	if errorr != nil {
		fmt.Println(errorr)
	} else {
		for _, location := range locationResult.Results {
			fmt.Println(location.Name)
		}
	}
	fmt.Println(locationResult.Next)

	return nil

}
