package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/bitscuit/stew-chicken-bot/internal"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	cmd := strings.TrimPrefix(m.Content, ",,")
	cmd = strings.Trim(cmd, " ")
	if strings.HasPrefix(cmd, "recipe") {
		args := strings.TrimPrefix(cmd, "recipe")
		args = strings.Trim(args, " ")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		baseUrl := "https://www.themealdb.com/api/json/v1/"
		apiKey := "1"
		searchByName := "/search.php?s="
		mealName := url.QueryEscape(args)
		url := baseUrl + apiKey + searchByName + mealName
		fmt.Println(url)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed request")
		}
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed do")
		}
		defer resp.Body.Close()
		// fmt.Println(resp.Body)
		// body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed read resp")
		}
		var body internal.Meals
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			fmt.Println(err)
		}
		fmt.Println(body)
		if len(body.Meal) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Could not find that recipe")
		} else {
			s.ChannelMessageSend(m.ChannelID, body.Meal[0].Name+" recipe: "+url)
		}
		// s.ChannelMessageSend(m.ChannelID, url)
		// fmt.Println(body.Meal[0].Name)
		// fmt.Println(string(body))

	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
