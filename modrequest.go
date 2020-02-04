package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const modRequest = "<@%s> requested a moderator in <#%s> here: %s\nReason: %s"

func (b *bot) requestMod(m *discordgo.Message, parts []string) {
	url := composeMessageURL(m)
	if len(parts) < 1 {
		message := "No message provided."
	} else {
		message := strings.Join(parts)
	}
	b.discord.ChannelSendMessage(feedChannel, fmt.Sprintf(modRequest, m.Author.ID, m.ChannelID, url, message)
}
