package start_end

import (
	"github.com/bwmarrin/discordgo"
)

func Leave(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.GuildLeave(m.GuildID)
}
