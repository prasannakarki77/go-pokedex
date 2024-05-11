package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonsInArea(location string) (LocationAreaPokemonEncounterResp, error) {
	endpoint := "/location-area"

	fullURL := baseURL + endpoint + "/" + location

	data, ok := c.cache.Get(fullURL)

	if ok {
		// Cache hit
		locationAreaPokemonEncounterResp := LocationAreaPokemonEncounterResp{}
		err := json.Unmarshal(data, &locationAreaPokemonEncounterResp)
		if err != nil {
			return LocationAreaPokemonEncounterResp{}, err
		}
		return locationAreaPokemonEncounterResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaPokemonEncounterResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaPokemonEncounterResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaPokemonEncounterResp{}, fmt.Errorf("bad status code:%v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaPokemonEncounterResp{}, err
	}

	locationAreaPokemonEncounterResp := LocationAreaPokemonEncounterResp{}

	err = json.Unmarshal(dat, &locationAreaPokemonEncounterResp)

	if err != nil {
		return LocationAreaPokemonEncounterResp{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationAreaPokemonEncounterResp, nil

}
