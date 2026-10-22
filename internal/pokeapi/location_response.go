package pokeapi

// LocResponse represents the JSON response from the PokeAPI's
// location-area list endpoint. The `json:"..."` tags tell
// encoding/json how to map each JSON key to these Go fields
// (Go fields must be capitalized to be exported, but JSON keys
// are lowercase, so the tags bridge that naming gap).
//
// Results is a slice of anonymous structs rather than a named type
// since each entry only has two fields (Name and URL) and isn't
// reused anywhere else in this package.

type LocResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
