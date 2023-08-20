package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Init_commands(s *discordgo.Session) {
	log.Println("Adding commands...")
	commandList := add_commands()

	for _, cmd := range commandList {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd.data)
		if err != nil {
			log.Fatal("failed to register command", err)
		}
		s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			cmd.exec(s, i)
		},
		)
	}
}

func add_commands() []command {
	commandList := make([]command, 0)

	commandList = append(commandList, init_ping())

	return commandList
}

type command struct {
	data *discordgo.ApplicationCommand
	exec func(s *discordgo.Session, i *discordgo.InteractionCreate)
}
