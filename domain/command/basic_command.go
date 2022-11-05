package command

import (
	"discord-spam-bot/domain/model"

	"github.com/bwmarrin/discordgo"
)

func (c *CommandService) basicCommand() *model.CommandNameHandlerParam {
	return &model.CommandNameHandlerParam{
		Command: &discordgo.ApplicationCommand{
			Name:        "hello",
			Description: "This is hello world for discord bot",
			Type:        discordgo.ChatApplicationCommand,
		},

		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	}
}
