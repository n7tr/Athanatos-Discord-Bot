package start_end

import (
	"Athanatos/core/requests"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

func Leave(s *discordgo.Session, m *discordgo.MessageCreate) {
	data := []byte{}
	jsonData, _ := json.Marshal(data)

	requests.Sendhttp("https://discord.com/api/v9/users/@me/guilds/"+m.GuildID, "DELETE", jsonData)
}
