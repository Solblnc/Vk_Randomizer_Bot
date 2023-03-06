package main

import (
	"Vk_Likes_Randomizer/pkg/VkBot"
	"Vk_Likes_Randomizer/pkg/config"
	"github.com/SevereCloud/vksdk/v2/api"
)

func main() {

	token := config.FromEnv("token")
	vk := api.NewVK(token)

	bot := VkBot.NewBot(vk)

	bot.GetLikes()

}
