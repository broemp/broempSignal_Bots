package commands

import (
	"log"
	"strconv"

	"github.com/broemp/broempSignal_Bots/api"
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
	var message string

	user := i.ApplicationCommandData().Options[0].UserValue(s)
	userid, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		log.Println("failed to parse userid: ", err)
	}

	afk_resp, err := api.AFK_create(userid, user.Username)
	if err != nil {
		message = "Failed api request" + err.Error()
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: message,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "AFK Report",
					Color: 10181046,
					Image: &discordgo.MessageEmbedImage{
						URL: user.AvatarURL(""),
					},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Culprit",
							Value: user.Username,
						},
						{
							Name:  "Time",
							Value: afk_resp.CreatedAt.Time.String(),
						},
					},
				},
			},
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
	if len(i.ApplicationCommandData().Options) != 0 {
		// User Passed
		user := i.ApplicationCommandData().Options[0].UserValue(s)
		userid, err := strconv.ParseInt(user.ID, 10, 64)
		if err != nil {
			log.Println("failed to parse userid: ", err)
		}

		afk_list := api.AFK_get_user(userid)
		fields := []*discordgo.MessageEmbedField{
			{
				Name:  "Count",
				Value: strconv.Itoa(len(afk_list)),
			},
		}

		for _, afk := range afk_list {
			field := discordgo.MessageEmbedField{
				Name: afk.CreatedAt.Time.String(),
			}
			fields = append(fields, &field)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "AFK Report",
						Description: user.Username,
						Color:       10181046,
						Image: &discordgo.MessageEmbedImage{
							URL: user.AvatarURL(""),
						},
						Fields: fields,
					},
				},
			},
		},
		)

	} else {
		// No User Passed
		log.Println("printingTopList")
		topList := api.AFK_get_top_list()

		var topListString string
		for i, user := range topList {
			topListString += strconv.Itoa(i+1) + ". " + user.Username + " " + strconv.Itoa(user.Count) + "\n"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "AFK TOP LIST",
						Color:       15548997,
						Description: topListString,
					},
				},
			},
		})

	}
}
