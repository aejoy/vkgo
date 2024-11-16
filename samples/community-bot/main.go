package main

import (
	"log"
	"os"

	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/samples/scenes"
	"github.com/aejoy/vkgo/scene"

	"github.com/aejoy/vkgo/api"
)

func main() {
	bot, err := api.New(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	log.Fatalln(longpoll.Start(bot, scene.Message(scenes.NewMessageScene)))
}
