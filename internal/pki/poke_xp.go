package pki

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokeInfo(pokemon *string) (RespShallowPokemon, error) {
	url := BaseURL + "/pokemon/" + *pokemon

	if val, ok := c.cache.Get(url); ok {
		pokeInfo := RespShallowPokemon{}
		err := json.Unmarshal(val, &pokeInfo)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		return pokeInfo, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	pokeInfo := RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokeInfo)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokeInfo, nil
}
