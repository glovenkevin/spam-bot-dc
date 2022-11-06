package command

import "github.com/bwmarrin/discordgo"

func (c *CommandService) errorMiddleware(originalHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// fmt.Println("Running before handler")

		errorHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			c.log.Errorf("Bad Command interaction: %s", recover())
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "505 | 💥💥💥 Bad Command!!!",
				},
			})
		}
		defer errorHandler(s, i)

		originalHandler(s, i)

		// fmt.Println("Running after handler")
	}
}
