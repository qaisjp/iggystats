package main

import (
    "fmt"

    "github.com/bwmarrin/discordgo"
)

const modRequest = "<@%s> requested a moderator in <#%s> here: %s\nReason: %s"

func (b *bot) requestMod(s *discordgo.Session, m *discordgo.Message, parts []string) {
    url := composeMessageURL(m)
    s.ChannelSendMessage(feedChannel, fmt.Printf(modRequest, m.Author.ID, m.ChannelID, url, parts)
}
