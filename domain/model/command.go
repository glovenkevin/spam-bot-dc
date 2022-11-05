package model

import "github.com/bwmarrin/discordgo"

type CommandNameHandlerParam struct {
	Command *discordgo.ApplicationCommand
	Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}
