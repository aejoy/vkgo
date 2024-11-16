package main

import (
	"os"

	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/multibot"
	"github.com/aejoy/vkgo/samples/scenes"
	"github.com/aejoy/vkgo/scene"
	"go.uber.org/zap"
)

func main() {
	diamond, err := api.New(os.Getenv("DIAMOND_TOKEN"))
	if err != nil {
		panic(err)
	}

	sapphire, err := api.New(os.Getenv("SAPPHIRE_TOKEN"))
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	diamond.Logger = logger
	sapphire.Logger = logger

	bots := []*api.API{diamond, sapphire}

	messageScene := scene.Message(scenes.NewMessageScene)

	multibot.New(bots, messageScene, logger).Start()
}
