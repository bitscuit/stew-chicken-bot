package main

import (
	"flag"
	"fmt"
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
	defer dg.Close()

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
	if (m.Author.ID == s.State.User.ID) || (m.Author.Bot) {
		return
	}

	// commands start with ",,"
	if !strings.HasPrefix(m.Content, ",,") {
		return
	}
	msg := strings.TrimPrefix(m.Content, ",,")
	msg = strings.Trim(msg, " ")
	action := strings.Split(msg, " ")
	cmd := action[0]
	args := strings.Join(action[1:], " ")

	if cmd == "recipe" {
		recipe, err := internal.FetchRecipe(args)
		fmt.Println(recipe)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, recipe.Name+" recipe: "+recipe.Url)
		}
	} else if cmd == "ps2" {
		result, err := internal.Ps2(args)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, result)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "unrecognized command")
	}
}
