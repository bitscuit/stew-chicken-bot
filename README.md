# Overview
Personal Discord bot serving as a sandbox for various interesting APIs. Used discordgo's [pingpong example](https://github.com/bwmarrin/discordgo/tree/master/examples/pingpong) bot as a starting base.


This Bot will respond to "ping" with "Pong!", followed with an URL from TheMealDB API for a stew chicken recipe, and "pong" with "Ping!".

# Run

## Pre-requisites

- Working [Go](https://golang.org/dl/) environment
- Bot application created on [Discord Devleoper](https://discordapp.com/developers/applications/)
- Bot token from Discord Developer ready

```sh
go get github.com/bitscuit/stew-chicken-bot

cd $GOPATH/bin

./stew-chicken-bot -t $TOKEN
```
