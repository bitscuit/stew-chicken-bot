package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Players struct {
	Player []Player `json:"character_list"`
}

type Player struct {
	Cert struct {
		AvailablePoints string `json:"available_points"`
	} `json:"certs"`
}

func Ps2(args string) (string, error) {
	msg := strings.Split(args, " ")
	cmd := msg[0]
	args = strings.Join(msg[1:], " ")

	if cmd == "certs" {
		return getCerts(args)
	}
	if cmd == "alert" {
		return isAlert()
	}
	return "false", nil
}

func getCerts(args string) (string, error) {
	fmt.Println(args)
	player := strings.ToLower(args)
	baseUrl := "http://census.daybreakgames.com"
	path := "/get/ps2:v2/character?name.first_lower="
	player = url.QueryEscape(player)
	url := baseUrl + path + player
	fmt.Println(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New("Failed request")
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Failed do")
	}
	defer resp.Body.Close()

	var body Players
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		fmt.Println("Failed to decode json")
		fmt.Println(err)
	}

	if len(body.Player) < 1 {
		return "", errors.New("Could not find that player")
	}

	return args + " has " + body.Player[0].Cert.AvailablePoints + " certs", nil
}

func isAlert() (string, error) {
	return "WIP to get alerts", nil
}
