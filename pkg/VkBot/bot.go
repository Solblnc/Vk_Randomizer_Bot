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

	//parametr := api.Params{}
	//parametr["group_id"] = "https://vk.com/seleroad"
	//_, err := b.bot.GroupsGet()
	//if err != nil {
	//	log.Printf("%d", err)
	//}

	// Initializing longpool

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

	// Create a new message
	//lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
	//	log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)
	//
	//	f := params.NewMessagesSendBuilder()
	//	f.Message(obj.Message.Text + " !!!")
	//	f.RandomID(0)
	//	f.PeerID(obj.Message.PeerID)
	//
	//	_, err := b.bot.MessagesSend(f.Params)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}
