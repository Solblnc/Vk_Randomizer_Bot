package main

import (
	"Vk_Likes_Randomizer/pkg/VkBot"
	"github.com/SevereCloud/vksdk/v2/api"
)

func main() {
	token := "vk1.a.nMAHwVOmdzxi1kpy1TSA-MST3EFAHn0LAXiwjAVXr1_fQTNoM_liCQl4e8dYJyOtDZwAwQQOZPIfwuBz7Kc8W-X3bcMt6IS_6WG25QhgFMlckEm7DnbhnSdckg7K-wMQrgDIKdr0evsxx8i-5s0VxNippJ28pDGMAQOpX7JzohUSn55olKEt08oXzY_S2B8ExnX6d6uSciGnZpxLhCw9HQ" // use os.Getenv("TOKEN")
	vk := api.NewVK(token)

	bot := VkBot.NewBot(vk)

	bot.GetLikes()

}
