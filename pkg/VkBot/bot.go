package VkBot

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"

	"log"
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

func (b *Bot) Start() error {



	group, err := b.bot.GroupsGetByID(api.Params{
		"group_id": "191968702",
	})
	if err != nil {
		log.Fatal(err)
	}

	lp, err := longpoll.NewLongPoll(b.bot, group[0].ID) //191968702
	if err != nil {
		log.Printf("%d", err)
	}
	

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}
