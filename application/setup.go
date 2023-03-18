package application

import (
	"discord-spam-bot/config"
	"discord-spam-bot/lib/constant"
	"discord-spam-bot/lib/pkg/discord"
	coreLogger "discord-spam-bot/lib/pkg/loggerext/core"
	"os"
)

func Setup() (Application, error) {
	var (
		app      Application
		baseInit = []func(*Application) error{
			initLogger(),
			initDiscord(),
		}
	)

	err := config.LoadConfigEnv(app.log)
	if err != nil {
		return app, err
	}

	err = runInit(baseInit...)(&app)
	if err != nil {
		return app, err
	}

	return app, nil
}

func initLogger() func(*Application) error {
	return func(app *Application) error {
		app.log = coreLogger.NewCoreLogger(coreLogger.LevelInfo)

		env := os.Getenv(constant.AppEnv)
		if env == "local" {
			app.log = coreLogger.NewCoreLogger(coreLogger.LevelDebug)
		}
		return nil
	}
}

func initDiscord() func(*Application) error {
	return func(app *Application) error {
		dc, err := discord.NewDiscordgoInstance(app.log)
		if err != nil {
			return err
		}

		app.discord = dc
		return nil
	}
}

func runInit(pp ...func(app *Application) error) func(app *Application) error {
	return func(app *Application) error {
		for _, fn := range pp {
			err := fn(app)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
