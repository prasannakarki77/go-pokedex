package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)

	if ok {
		// Cache hit
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code:%v", resp.StatusCode)
	}
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreaResp, nil
}
