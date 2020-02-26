# Overview
Personal Discord bot serving as a sandbox for various interesting APIs. Used discordgo's [pingpong example](https://github.com/bwmarrin/discordgo/tree/master/examples/pingpong) bot as a starting base.

The bot currently listens to messages in the Discord server and performs an action if given a command. Valid commands are prefixed with `,,`, and there are two commands:
- `,,recipe`
- `,,ps2`

## Recipe
The recipe command takes one argument, `recipe-name`, and calls TheMealDB API to search for a recipe by name. Then the URL to the recipe is returned.
```
,,recipe stew chicken
```

## PlanetSide
The ps2 command has two subcommands which interact with the PlanetSide 2 API (Census), `certs` and `alert`. The `certs` subcommand takes one argument, `player-name`, and returns that player's certs.
```
,,ps2 certs cakeisnice
```

The `alert` subcommand returns a list of active Alerts within the past hour in PlanetSide 2 Emerald server along with the duration since each one has started in minutes.
```
,,ps2 alert
```

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
