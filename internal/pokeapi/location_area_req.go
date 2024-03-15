package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaRes, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Print("cache hit!\n")
		// cahce hit
		locationAreaRes := LocationAreaRes{}
		err := json.Unmarshal(dat, &locationAreaRes)
		if err != nil {
			return LocationAreaRes{}, err
		}
		return locationAreaRes, nil
	}
	fmt.Print("cache miss!\n")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaRes{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaRes{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	locationAreaRes := LocationAreaRes{}
	err = json.Unmarshal(dat, &locationAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreaRes, nil
}

func (c *Client) GetLocationAreas(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Print("cache hit!\n")
		// cahce hit
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Print("cache miss!\n")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationArea, nil
}
