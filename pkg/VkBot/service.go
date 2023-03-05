package VkBot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	"log"
)

func (b *Bot) GetLikes() []object.UsersUser {
	p := params.NewLikesGetListBuilder()

	p.Type("post")
	p.Count(1000)
	p.OwnerID(-191968702)
	p.ItemID(100)

	list, err := b.bot.LikesGetListExtended(p.Params)

	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(list.Items, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var res []User
	reader := bytes.NewReader(data)
	dec := json.NewDecoder(reader)
	err = dec.Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	return list.Items
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
