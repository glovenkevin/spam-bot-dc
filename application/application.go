package application

import (
	"discord-spam-bot/lib/pkg/loggerext"

	"github.com/bwmarrin/discordgo"
)

type Application struct {
	log     loggerext.LoggerInterface
	discord *discordgo.Session
}
