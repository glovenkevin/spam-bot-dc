package command

import (
	"bytes"
	"context"
	"discord-spam-bot/domain/model"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func (c *CommandService) generateAmythestWantedImage() *model.CommandNameHandlerParam {
	const (
		imageUrl = "image-url"
	)
	return &model.CommandNameHandlerParam{
		Command: &discordgo.ApplicationCommand{
			Name:        "generate-wanted",
			Description: "Generate Image Wanted using Amythest API",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        imageUrl,
					Description: "Put your image URL to be placed inside the template here (PNG, JPG, JPEG)",
					Required:    true,
				},
			},
		},

		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMaps := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
			for _, o := range options {
				optionMaps[o.Name] = o
			}

			url := optionMaps[imageUrl].StringValue()
			c.sendInteractionResponds(&model.SendInteractionRespondParam{
				Ic:      i,
				Message: "OK Bro",
			})

			userId := i.Member.User.ID
			imgByte, err := c.amythestRepo.GenerateWanted(context.Background(), url)
			if err != nil {
				c.log.Error(err)
				s.ChannelMessage(i.ChannelID, fmt.Sprintf("Fail to generate image for <@%s>", userId))
				return
			}

			var sf string
			mt := http.DetectContentType(imgByte)
			switch mt {
			case "image/jpeg":
				sf = ".jpeg"
			case "image/png":
				sf = ".png"
			}

			_, err = s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
				Content: fmt.Sprintf("Here is your image sir <@%s>", userId),
				File: &discordgo.File{
					Name:        fmt.Sprintf("%s_wanted%s", userId, sf),
					ContentType: mt,
					Reader:      bytes.NewReader(imgByte),
				},
			})
			if err != nil {
				c.log.Error(err)
				return
			}
		},
	}
}
