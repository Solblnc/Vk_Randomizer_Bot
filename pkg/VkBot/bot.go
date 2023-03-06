package VkBot

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

//type Params struct {
//	ChatID int
//	PeerID int
//}

type Bot struct {
	bot *api.VK
}

func NewBot(bot *api.VK) *Bot {
	return &Bot{bot: bot}
}
