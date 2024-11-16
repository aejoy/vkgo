package longpoll

import (
	"encoding/json"
	"fmt"

	"github.com/aejoy/vkgo/consts"
	states "github.com/aejoy/vkgo/longpoll/user"
)

func (lp *LongPoll) RunUserLongPoll() {
	messagesCh := make(chan []any)
	go func() {
		lp.GetUserLongPollUpdates("https://"+lp.Session.Server, messagesCh)
	}()

	for {
		message := <-messagesCh
		go func() {
			if len(message) != 0 {
				if typ, ok := message[0].(float64); ok {
					if event := states.GetUpdate(typ); event != nil {
						event(lp.Bot, message, *lp.Scenes)
					}
				}
			}
		}()
	}
}

func (lp *LongPoll) GetUserLongPollUpdates(server string, dest chan []any) {
	for {
		body, err := lp.getByteUpdates(server)
		if err != nil {
			return
		}

		messages := states.Messages{}
		if err := json.Unmarshal(body, &messages); err != nil {
			if lp.Bot.Logger != nil {
				lp.Bot.Logger.Error(fmt.Sprintf("MessagesLongPoll JSON unmarshal error is %s", err.Error()))
			}

			return
		}

		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Info(fmt.Sprintf("MessagesLongPoll response is %s", string(body)))
		}

		switch messages.Failed {
		case consts.UpdatesStatusOK:
			lp.Session.TS = messages.TS

			for _, message := range messages.Updates {
				dest <- message
			}
		case consts.UpdatesStatusOutdated:
			lp.Session.TS = messages.TS
		case consts.UpdatesStatusExpired:
			if server, err := lp.Bot.GetUserLongPollServer(); err != nil {
				lp.Session.Key = server.Key
			}
		case consts.UpdatesStatusLost:
			if server, err := lp.Bot.GetUserLongPollServer(lp.Bot.ID); err == nil {
				lp.Session.Key = server.Key
				lp.Session.TS = server.TS
			}
		case consts.UpdatesStatusInvalid:
			lp.Session.Version = messages.MinVersion
		}
	}
}
