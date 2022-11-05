package command

import (
	"discord-spam-bot/domain/model"

	"github.com/bwmarrin/discordgo"
)

func (c *CommandService) sendInteractionResponds(opt *model.SendInteractionRespondParam) error {
	err := c.dc.InteractionRespond(opt.Ic.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: opt.Message,
		},
	})
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}
