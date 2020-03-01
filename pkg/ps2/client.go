package ps2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}

type Players struct {
	Player []struct {
		Cert struct {
			Available string `json:"available_points"`
		} `json:"certs"`
	} `json:"character_list"`
}

func (c *Client) GetCerts(player string) (string, error) {
	path := &url.URL{Path: "/get/ps2:v2/character"}
	path.RawQuery = "name.first_lower=" + player
	u := c.BaseURL.ResolveReference(path)
	resp, err := c.HttpClient.Get(u.String())
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Invalid URL")
	}

	var players Players
	if err := json.NewDecoder(resp.Body).Decode(&players); err != nil {
		fmt.Println(err)
		return "", errors.New("API returned invalid JSON")
	}
	defer resp.Body.Close()

	if len(players.Player) < 1 {
		return "", errors.New("Could not find that player")
	}

	fmt.Println(u.String())
	s := fmt.Sprintf("%s has %s certs", player, players.Player[0].Cert.Available)
	return s, nil
}
