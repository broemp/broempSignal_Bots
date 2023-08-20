package discord

import (
	"log"
	"os"
	"os/signal"

	"github.com/broemp/broempSignal_Bots/discord/commands"
	"github.com/bwmarrin/discordgo"
)

func InitDiscord(discordToken string) {
	log.Println("Starting Discord Bot")
	s, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal("Couldn't create Discord session: ", err)
	}

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	commands.Init_commands(s)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
