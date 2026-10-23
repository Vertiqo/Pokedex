package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Fetcher is helper that helped in refactoring the methods below
// by taking over the shared requirements to keep the code DRY
func (c Client) Fetcher(url string) ([]byte, error) {
	value, ok := c.cache.Get(url)
	if ok {
		// fmt.Println("cache used") debug test to see if cache used or not	fmt.Println("cache used")
		return value, nil
	}
	// fmt.Println("making http request") debug test to see if cache used or not
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, data)
	return data, nil
}

// FetchLocations returns locations from the cache when available.
// Otherwise, it fetches them from the API and caches the response.
func (c Client) FetchLocations(url string) (LocResponse, error) {
	data, err := c.Fetcher(url)
	if err != nil {
		return LocResponse{}, err
	}
	resp := LocResponse{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return LocResponse{}, err
	}
	return resp, nil

}

func (c Client) FetchEncounters(url string) (LocEncounters, error) {
	data, err := c.Fetcher(url)
	if err != nil {
		return LocEncounters{}, err
	}
	resp := LocEncounters{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return LocEncounters{}, err
	}
	return resp, nil
}
