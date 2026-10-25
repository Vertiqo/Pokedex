package pokeapi

import (
	"encoding/json"
)

func (c Client) FetchPokemon(url string) (Pokemon, error) {
	data, err := c.fetcher(url)
	if err != nil {
		return Pokemon{}, err
	}
	resp := Pokemon{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return Pokemon{}, err
	}
	return resp, nil

}
