package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Recipe struct {
	Name string
	Url  string
}

type Meals struct {
	Meal []Meal `json:"meals"`
}

type Meal struct {
	Name string `json:"strMeal"`
}

func FetchRecipe(args string) (Recipe, error) {
	fmt.Println(args)
	baseUrl := "https://www.themealdb.com/api/json/v1/"
	apiKey := "1"
	searchByName := "/search.php?s="
	mealName := url.QueryEscape(args)
	url := baseUrl + apiKey + searchByName + mealName
	fmt.Println(url)
	r := Recipe{"", url}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return r, errors.New("Failed request")
		// s.ChannelMessageSend(m.ChannelID, "Failed request")
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return r, errors.New("Failed do")
		// s.ChannelMessageSend(m.ChannelID, "Failed do")
	}
	defer resp.Body.Close()

	// fmt.Println(resp.Body)
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	s.ChannelMessageSend(m.ChannelID, "Failed read resp")
	// }

	var body Meals
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		fmt.Println("Failed to decode json")
		fmt.Println(err)
	}
	// fmt.Println(body)
	if len(body.Meal) < 1 {
		return r, errors.New("Could not find that recipe")
		// s.ChannelMessageSend(m.ChannelID, "Could not find that recipe")
	}

	r.Name = body.Meal[0].Name
	return r, nil
}
