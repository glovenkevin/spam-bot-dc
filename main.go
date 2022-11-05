package main

import (
	"discord-spam-bot/config"
	"discord-spam-bot/domain/command"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := config.NewLogger()
	err := config.LoadConfig(l)
	if err != nil {
		panic(err)
	}

	dc, err := config.NewDiscordInstance(l)
	if err != nil {
		panic(err)
	}
	defer dc.Close()

	c := command.NewCommandService(dc, l)
	err, td := c.RegisterHandlers()
	if err != nil {
		panic(err)
	}
	defer td()

	l.Info("Bot started...")
	s := make(chan os.Signal, 1)
	defer close(s)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	<-s
}
