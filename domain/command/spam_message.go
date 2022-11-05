package command

import (
	"context"
	"discord-spam-bot/domain/model"
	"discord-spam-bot/lib/pkg/helper"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (c *CommandService) spamMessage() *model.CommandNameHandlerParam {
	const (
		userTarget   = "user-target"
		totalMessage = "total-message"
		interval     = "interval"
		spamMessage  = "spam-message"
	)

	return &model.CommandNameHandlerParam{
		Command: &discordgo.ApplicationCommand{
			Name:        "spam",
			Description: "Spam user that has been specified",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        userTarget,
					Description: "Mention user that being targeted for spam",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        totalMessage,
					Description: "Total spam message being sent",
					Required:    false,
					MinValue:    helper.NewFloat64(1),
					MaxValue:    100,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        interval,
					Description: "Interval between spam message being sent",
					Required:    false,
					MinValue:    helper.NewFloat64(1),
					MaxValue:    5,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        spamMessage,
					Description: "Spam message content",
					Required:    false,
					MinLength:   helper.NewInt(3),
				},
			},
		},

		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			optionMaps := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
			for _, o := range options {
				optionMaps[o.Name] = o
			}

			if optionMaps[totalMessage] == nil {
				optionMaps[totalMessage] = &discordgo.ApplicationCommandInteractionDataOption{
					Value: float64(1),
					Type:  discordgo.ApplicationCommandOptionInteger,
					Name:  totalMessage,
				}
			}
			if optionMaps[interval] == nil {
				optionMaps[interval] = &discordgo.ApplicationCommandInteractionDataOption{
					Value: float64(1),
					Type:  discordgo.ApplicationCommandOptionInteger,
					Name:  interval,
				}
			}
			if optionMaps[spamMessage] == nil || optionMaps[spamMessage].StringValue() == "" {
				optionMaps[spamMessage] = &discordgo.ApplicationCommandInteractionDataOption{
					Value: "Hei youu, youu, 你的眼睛在哪里？",
					Name:  totalMessage,
					Type:  discordgo.ApplicationCommandOptionString,
				}
			}

			user := optionMaps[userTarget].UserValue(s)
			channel, err := s.UserChannelCreate(user.ID)
			if err != nil {
				c.log.Error(err)
				c.sendInteractionResponds(&model.SendInteractionRespondParam{
					Ic:      i,
					Message: fmt.Sprintf("[MISSION FAILED] Failed to spam user <@%s> : %s", user.ID, err.Error()),
				})
				return
			}

			c.sendInteractionResponds(&model.SendInteractionRespondParam{
				Ic:      i,
				Message: fmt.Sprintf("[MISSION SUCCESS] Success spam message to user <@%s>", user.ID),
			})

			interval := optionMaps[interval].IntValue()
			count := optionMaps[totalMessage].IntValue()
			doneFlag := make(chan struct{}, 1)
			defer close(doneFlag)
			callback := func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}

				for count > 0 {
					_, err := s.ChannelMessageSend(channel.ID, optionMaps[spamMessage].StringValue())
					if err != nil {
						return err
					}

					time.Sleep(time.Duration(interval) * time.Second)
					count--
				}

				doneFlag <- struct{}{}
				return nil
			}

			timeout := 5*time.Second + (time.Duration(count) * time.Duration(interval) * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			err = callback(ctx)
			<-doneFlag
			if err != nil {
				c.log.Error(err)
			}
		},
	}
}
