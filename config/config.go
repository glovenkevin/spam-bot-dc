package config

import (
	"discord-spam-bot/lib/constant"
	"discord-spam-bot/lib/pkg/loggerext"
	"fmt"
	"os"
	"strings"
)

func LoadConfigEnv(l loggerext.LoggerInterface) error {
	envFile := ".env"
	env := os.Getenv(constant.AppEnv)
	if env != "" && env != "local" {
		envFile = fmt.Sprintf(".%s.env", env)
	}

	b, err := os.ReadFile(envFile)
	if err != nil {
		l.Error(err)
		return err
	}

	ss := strings.Split(string(b), "\n")
	for _, s := range ss {
		ee := strings.Split(s, "=")
		if len(ee) == 2 {
			os.Setenv(ee[0], ee[1])
		}
	}

	return nil
}
