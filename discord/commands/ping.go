package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init_ping(s *discordgo.Session) {
	const name string = "ping"
	const description string = "Responds with Ping!"

	cmd := discordgo.ApplicationCommand{Name: name, Description: description}

	add_commands(s, &cmd, exec_ping)
}

func exec_ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
