package pki

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocInfo(loc *string) (RespShallowLocInfo, error) {
	url := BaseURL + "/location-area/" + *loc

	if val, ok := c.cache.Get(url); ok {
		locInfoResp := RespShallowLocInfo{}
		err := json.Unmarshal(val, &locInfoResp)
		if err != nil {
			return RespShallowLocInfo{}, err
		}

		return locInfoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocInfo{}, err
	}

	locInfoResp := RespShallowLocInfo{}
	err = json.Unmarshal(dat, &locInfoResp)
	if err != nil {
		return RespShallowLocInfo{}, err
	}

	c.cache.Add(url, dat)
	return locInfoResp, nil
}
