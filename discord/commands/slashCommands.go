package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)

func Init_commands(s *discordgo.Session) {
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
	remove_old_commands(s)
	log.Println("Adding commands...")

	init_ping(s)
	init_afk(s)
	init_afk_list(s)

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func add_commands(s *discordgo.Session, discord_cmd *discordgo.ApplicationCommand, exec func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", discord_cmd)
	if err != nil {
		log.Fatal("failed to register command: ", err)
	}

	commandHandlers[discord_cmd.Name] = exec
}

func remove_old_commands(s *discordgo.Session) {
	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Fatal("failed to remove old commands: ", err)
	}
	for _, cmd := range cmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		if err != nil {
			log.Fatal("failed to remove old commands: ", err)
		}
	}
}
