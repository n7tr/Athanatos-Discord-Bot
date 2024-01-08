package commands

import (
	"Dynamic/core/removing"
	"Dynamic/core/start_end"

	"github.com/bwmarrin/discordgo"
)

func Leave(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".leave" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		start_end.Leave(s, m)
	}
}

func BanAll(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".ban_all" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		removing.MemberBan(s, m)
	}
}
