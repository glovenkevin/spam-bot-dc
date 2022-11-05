package config

import (
	"discord-spam-bot/lib/constant"
	"discord-spam-bot/lib/pkg/logger"
	coreLogger "discord-spam-bot/lib/pkg/logger/core"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func NewLogger() logger.LoggerInterface {
	env := os.Getenv(constant.AppEnv)
	if env == "local" {
		return coreLogger.NewCoreLogger(coreLogger.LevelDebug)
	}

	return coreLogger.NewCoreLogger(coreLogger.LevelInfo)
}

func NewDiscordInstance(l logger.LoggerInterface) (*discordgo.Session, error) {
	token := os.Getenv(constant.DiscordToken)
	dc, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		l.Error(err)
		return nil, err
	}

	dc.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		l.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = dc.Open()
	if err != nil {
		return nil, err
	}

	return dc, nil
}
