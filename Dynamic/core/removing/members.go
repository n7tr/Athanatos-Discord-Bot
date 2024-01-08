package removing

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func MemberBan(s *discordgo.Session, m *discordgo.MessageCreate) {
	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		log.Println("Can't take the list of members:", err)
		return
	}

	for _, member := range members {
		id := member.User.ID

		err := s.GuildBanCreate(m.GuildID, id, 0)
		if err != nil {
			log.Println("Can't ban members:", err)
			return
		}
	}
}
