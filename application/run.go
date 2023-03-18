package application

import (
	"discord-spam-bot/domain/command"
	"os"
	"os/signal"
	"syscall"
)

func (app Application) Run() error {
	c := command.NewCommandService(app.discord, app.log)
	err, td := c.RegisterHandlers()
	if err != nil {
		return err
	}
	defer td()

	app.log.Info("Bot started...")
	s := make(chan os.Signal, 1)
	defer close(s)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	_ = <-s

	return nil
}
