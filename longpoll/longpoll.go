package longpoll

import (
	"errors"

	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/consts"
	"github.com/aejoy/vkgo/responses"
	"github.com/aejoy/vkgo/scene"
)

var (
	ErrNotFound      = errors.New("bot not found")
	ErrSessionClosed = errors.New("session closed")
)

// LongPoll - basic configuration structure
// Bot - pointer to the bot parameters
// Session - pointer to the session, received from the Create function
// Scenes - pointer to the list of scenes to process the types of events
// Timeout - time of waiting for the next request.
type LongPoll struct {
	Bot     *api.API
	Session *responses.LongPollServer
	Scenes  *scene.Scenes
	Timeout int
	Mode    int
	BotType int // 1 - community, 2 - user
}

func New(properties ...any) (*LongPoll, error) {
	lp := &LongPoll{
		Session: new(responses.LongPollServer),
		Scenes:  scene.NewScene(),
		Timeout: consts.DefaultLongPollTimeout,
		Mode:    consts.DefaultLongPollMode,
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *api.API:
			lp.Bot = property
		case *scene.Scenes:
			lp.Scenes = property
		}
	}

	if lp.Bot.ID == 0 {
		return nil, ErrNotFound
	}

	switch lp.Bot.IsGroup {
	case true:
		server, err := lp.Bot.GetGroupLongPollServer(lp.Bot.ID)
		if err != nil {
			return lp, err
		}

		lp.Session = &server
		lp.BotType = consts.CommunityBotType
	default:
		server, err := lp.Bot.GetUserLongPollServer()
		if err != nil {
			return lp, err
		}

		lp.Session = &server
		lp.BotType = consts.UserBotType
	}

	return lp, nil
}
