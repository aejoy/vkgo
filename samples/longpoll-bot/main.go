package main

import (
	"log"
	"os"

	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/samples/scenes"
	"github.com/aejoy/vkgo/scene"
)

func main() {
	bot, err := api.New(os.Getenv("USER_TOKEN"))
	if err != nil {
		panic(err)
	}

	log.Fatalln(longpoll.Start(bot, scene.Message(scenes.NewMessageScene)))
}
