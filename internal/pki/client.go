package pki

import (
	"github.com/eikonoklastess/pokedex-repl/internal/pkcache"
	"net/http"
	"time"
)

type Client struct {
	cache      pkcache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pkcache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
