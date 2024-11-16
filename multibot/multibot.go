package multibot

import (
	"fmt"

	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/longpoll"
	"github.com/aejoy/vkgo/scene"
)

type MultiBot struct {
	Bots   []*api.API
	Scenes *scene.Scenes
}

func New(properties ...any) *MultiBot {
	mb := new(MultiBot)

	for _, property := range properties {
		switch property := property.(type) {
		case *api.API:
			mb.Bots = append(mb.Bots, property)
		case []*api.API:
			mb.Bots = append(mb.Bots, property...)
		case *scene.Scenes:
			mb.Scenes = property
		}
	}

	if mb.Bots == nil {
		mb.Bots = make([]*api.API, 0)
	}

	return mb
}

func (mb *MultiBot) Start() []error {
	var (
		errors   = make([]error, 0, len(mb.Bots))
		errorsCh = make(chan error)
	)

	for _, bot := range mb.Bots {
		go listenBot(bot, mb.Scenes, errorsCh)
	}

	for err := range errorsCh {
		errors = append(errors, err)
	}

	return errors
}

func listenBot(bot *api.API, scenes *scene.Scenes, errors chan error) {
	if bot.Logger != nil {
		bot.Logger.Info(fmt.Sprintf("MultiBot LongPoll  %d running", bot.ID))
	}

	if err := longpoll.Start(bot, scenes); err != nil {
		bot.Logger.Error(fmt.Sprintf("MultiBot LongPoll %d error is %s", bot.ID, err))
		errors <- err
	} else {
		errors <- nil
	}
}
