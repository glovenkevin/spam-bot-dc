package model

import "github.com/bwmarrin/discordgo"

type SendInteractionRespondParam struct {
	Ic      *discordgo.InteractionCreate
	Message string
}
