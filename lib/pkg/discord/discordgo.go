package discord

import (
	"discord-spam-bot/lib/constant"
	"discord-spam-bot/lib/pkg/loggerext"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func NewDiscordgoInstance(l loggerext.LoggerInterface) (*discordgo.Session, error) {
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
