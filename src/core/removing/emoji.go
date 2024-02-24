package removing

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func EmojiDelete(s *discordgo.Session, m *discordgo.MessageCreate) {
	emojis, err := s.GuildEmojis(m.GuildID)

	if err != nil {
		log.Println(err)
		return
	}

	for _, emoji := range emojis {
		err := s.GuildEmojiDelete(m.GuildID, emoji.ID)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
