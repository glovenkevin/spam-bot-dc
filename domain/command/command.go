package command

import (
	"discord-spam-bot/domain/model"
	"discord-spam-bot/domain/repo"
	amythest_http "discord-spam-bot/domain/repo/amythest/http"
	"discord-spam-bot/lib/constant"
	"discord-spam-bot/lib/pkg/loggerext"
	"net/http"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/sync/errgroup"
)

type CommandService struct {
	dc           *discordgo.Session
	log          loggerext.LoggerInterface
	amythestRepo repo.AmythestRepoInterface
}

func NewCommandService(dc *discordgo.Session, l loggerext.LoggerInterface) *CommandService {
	http := &http.Client{}

	return &CommandService{
		dc:           dc,
		log:          l,
		amythestRepo: amythest_http.NewAmythestRepoHttp(l, http),
	}
}

func (c *CommandService) getCommandItems() []*model.CommandNameHandlerParam {
	return []*model.CommandNameHandlerParam{
		c.basicCommand(),
		c.spamMessage(),
		c.generateAmythestWantedImage(),
		c.generateAmythestCircleImage(),
	}
}

func (c *CommandService) RegisterHandlers() (error, func()) {
	var (
		wg  errgroup.Group
		mtx sync.Mutex

		mapCommand  = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
		gid         = os.Getenv(constant.GuildID)
		userID      = c.dc.State.User.ID
		comandItems = c.getCommandItems()
	)

	for _, ci := range comandItems {
		ch := ci
		wg.Go(func() error {
			cmd, err := c.dc.ApplicationCommandCreate(userID, gid, ch.Command)
			if err != nil {
				return err
			}
			ch.Command.ID = cmd.ID
			ch.Command.ApplicationID = cmd.ApplicationID
			ch.Command.GuildID = cmd.GuildID

			mtx.Lock()
			mapCommand[ch.Command.Name] = ch.Handler
			mtx.Unlock()
			return nil
		})
	}
	if err := wg.Wait(); err != nil {
		c.log.Error(err)
		return err, nil
	}

	c.dc.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandName := i.ApplicationCommandData().Name
		if i.Member != nil {
			c.log.Debugf("Command Name: %s | Caller: %s | Interaction ID: %s", commandName, i.Member.User.Username, i.ID)
		}

		h, ok := mapCommand[commandName]
		if !ok {
			c.log.Errorf("Command is Not Found: %s", commandName)
			return
		}
		h(s, i)

		c.log.Debugf("Success execute Interaction ID: %s", i.ID)
	})

	td := func() {
		var wg sync.WaitGroup
		wg.Add(len(comandItems))
		for _, ci := range comandItems {
			go func(cm *discordgo.ApplicationCommand) {
				defer wg.Done()
				err := c.dc.ApplicationCommandDelete(cm.ApplicationID, cm.GuildID, cm.ID)
				if err != nil {
					c.log.Error(err)
				}
			}(ci.Command)
		}
		wg.Wait()
		c.dc.Close()
	}
	return nil, td
}
