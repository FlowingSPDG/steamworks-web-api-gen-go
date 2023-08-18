package steamworkswebapigen

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const endpoint = "https://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v1"

type GetSupportedAPIListResponse struct {
	Apilist struct {
		Interfaces []Interface `json:"interfaces"`
	} `json:"apilist"`
}

type Interface struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name       string      `json:"name"`
	Version    int         `json:"version"`
	Httpmethod string      `json:"httpmethod"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

func GetSupportedAPIList(key string) (*GetSupportedAPIListResponse, error) {
	// build URL
	u, err := url.Parse(endpoint)
	q := u.Query()
	q.Add("key", key)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := GetSupportedAPIListResponse{}
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
