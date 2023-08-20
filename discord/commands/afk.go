package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init_afk(s *discordgo.Session) {
	const name string = "afk"
	const description string = "Report somebody as AFK"

	cmd_options := []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The person you want to report",
			Required:    true,
			Type:        discordgo.ApplicationCommandOptionUser,
		},
	}

	cmd := discordgo.ApplicationCommand{Name: name, Description: description, Options: cmd_options}

	add_commands(s, &cmd, exec_afk)
}

func exec_afk(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Added",
		},
	})
}

func init_afk_list(s *discordgo.Session) {
	const name string = "list"
	const description string = "Get AFK Stats"

	cmd_options := []*discordgo.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The person you want to get infos on",
			Required:    false,
			Type:        discordgo.ApplicationCommandOptionUser,
		},
	}

	cmd := discordgo.ApplicationCommand{Name: name, Description: description, Options: cmd_options}

	add_commands(s, &cmd, exec_afk_list)
}

func exec_afk_list(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var message string
	if len(i.ApplicationCommandData().Options) >= 0 {
		// Handle User
		message = "List user info"
	} else {
		message = "Top List"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}
