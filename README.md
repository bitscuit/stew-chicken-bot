# Overview
Personal Discord bot serving as a sandbox for various interesting APIs. Used discordgo's [pingpong example](https://github.com/bwmarrin/discordgo/tree/master/examples/pingpong) bot as a starting base.


This Bot will respond to "ping" with "Pong!", followed with an URL from TheMealDB API for a stew chicken recipe, and "pong" with "Ping!".

# Run

## Pre-requisites

- Working [Go](https://golang.org/dl/) environment
- Bot application created on [Discord Devleoper Portal](https://discordapp.com/developers/applications/)
- Bot token ready from [Discord Devleoper Portal](https://discordapp.com/developers/applications/)
- Invited Bot application to your Discord server

### Invite bot to server

Replace `<client_id>` with the value from [Discord Devleoper Portal](https://discordapp.com/developers/applications/)

```
https://discordapp.com/api/oauth2/authorize?client_id=<client_id>&permissions=0&scope=bot
```

### Running

```sh
go get github.com/bitscuit/stew-chicken-bot

cd $GOPATH/bin

./stew-chicken-bot -t $TOKEN
```
