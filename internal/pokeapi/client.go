package pokeapi

import (
	"time"

	"github.com/vertiqo/pokedex/internal/pokecache"
)

// Client holds a shared cache for PokeAPI responses.

type Client struct {
	cache *pokecache.Cache
}

// NewClient creates a Client with a cache that removes expired
// entries at the given interval.

func NewClient(duration time.Duration) Client {
	client := Client{
		cache: pokecache.NewCache(duration),
	}
	return client

}
