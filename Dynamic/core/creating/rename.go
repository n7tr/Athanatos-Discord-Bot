package creating

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func GuildRename(s *discordgo.Session, m *discordgo.MessageCreate) {

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Error getting guild: ", err)
		return
	}

	godotenv.Load()
	server_name := os.Getenv("SERVER_NAME")

	avatarData, _ := ioutil.ReadFile("avatar.jpg")
	avatarBase64 := base64.StdEncoding.EncodeToString(avatarData)

	guildID := guild.ID
	s.GuildEdit(guildID, &discordgo.GuildParams{
		Name: server_name,
		Icon: "data:image/png;base64," + avatarBase64,
	})
}
