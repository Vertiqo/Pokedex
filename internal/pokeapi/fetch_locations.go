package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// FetchLocations returns locations from the cache when available.
// Otherwise, it fetches them from the API and caches the response.
func (c Client) FetchLocations(url string) (LocResponse, error) {

	value, ok := c.cache.Get(url)
	if ok {
		resp := LocResponse{}
		err := json.Unmarshal(value, &resp)
		if err != nil {
			return LocResponse{}, err
		}
		return resp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocResponse{}, err
	}
	response := LocResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return LocResponse{}, err
	}

	c.cache.Add(url, data)
	return response, nil

}
