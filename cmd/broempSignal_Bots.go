package main

import (
	"log"

	"github.com/broemp/broempSignal_Bots/api"
	"github.com/broemp/broempSignal_Bots/discord"
	"github.com/broemp/broempSignal_Bots/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Couldn't load config: ", err)
	}

	api.Init_API(config)
	discord.InitDiscord(config.DISCORD_TOKEN)
}
