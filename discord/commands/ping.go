package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init_ping() command {
	const name string = "ping"
	const description string = "Responds with Ping!"

	discord_cmd := discordgo.ApplicationCommand{Name: name, Description: description}
	cmd := command{&discord_cmd, exec_ping}

	return cmd
}

func exec_ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
