package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
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

type WorldEvents struct {
	Events []Event `json:"world_event_list"`
}

type Event struct {
	MetagameEventID    string   `json:"metagame_event_id"`
	Timestamp          string   `json:"timestamp"`
	MetagameEventState string   `json:"metagame_event_state_name"`
	MetagameEventType  struct { // only alerts will have this
		Name struct {
			English string `json:"en"`
		} `json:"name"`
	} `json:"metagame_event_id_join_metagame_event"`
}

func isAlert() (string, error) {
	searchTime := time.Now().Unix() - 3600
	baseUrl := "http://census.daybreakgames.com"
	path := "/get/ps2:v2/world_event"
	search := "?type=METAGAME&world_id=17&after=" + strconv.FormatInt(searchTime, 10) + "&c:limit=50&c:join=metagame_event^terms:description.en=*lock"
	url := baseUrl + path + search
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

	var body WorldEvents
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", errors.New("Failed to decode json")
	}

	fmt.Println(body)

	count := 0
	result := make(map[string]Event)
	// loop through Events to remove non-alerts and ended alerts
	for i := len(body.Events) - 1; i > -1; i-- {
		if body.Events[i].MetagameEventType.Name.English == "" { // might be a better way to do this
			continue
		}
		if body.Events[i].MetagameEventState == "ended" {
			v, ok := result[body.Events[i].MetagameEventID]
			if ok { // remove matching start alert
				delete(result, v.MetagameEventID)
			}
			delete(result, body.Events[i].MetagameEventID)
		} else {
			result[body.Events[i].MetagameEventID] = body.Events[i]
		}
		count++
	}
	fmt.Println(result)
	fmt.Println(count)
	alerts := ""
	for _, v := range result {
		ts, err := strconv.ParseInt(v.Timestamp, 10, 64)
		if err != nil {
			return "", errors.New("Something went horribly wrong")
		}
		duration := int(time.Since(time.Unix(ts, -1)).Minutes())
		alerts += v.MetagameEventType.Name.English + " started " + strconv.Itoa(duration) + " minutes go\n"
	}
	if alerts == "" {
		alerts = "No alerts"
	}
	return alerts, nil
}
