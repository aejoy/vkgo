package longpoll

import (
	"encoding/json"
	"fmt"

	"github.com/aejoy/vkgo/consts"
	"github.com/aejoy/vkgo/scene"
	"github.com/aejoy/vkgo/update"
)

func (lp *LongPoll) RunGroupLongPoll() {
	updatesCh := make(chan update.Update)
	go func() {
		lp.GetGroupLongPollUpdates(lp.Session.Server, updatesCh)
	}()

	for {
		upd := <-updatesCh
		go func() {
			scene.Use(lp.Bot, upd, *lp.Scenes)
		}()
	}
}

func (lp *LongPoll) GetGroupLongPollUpdates(server string, dest chan update.Update) {
	for {
		body, err := lp.getByteUpdates(server)
		if err != nil {
			return
		}

		updates := update.Updates{}
		if err := json.Unmarshal(body, &updates); err != nil {
			if lp.Bot.Logger != nil {
				lp.Bot.Logger.Error(fmt.Sprintf("GroupLongPoll JSON unmarshal error is %s", err.Error()))
			}

			return
		}

		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Info(fmt.Sprintf("GroupLongPoll response is %s", string(body)))
		}

		switch updates.Failed {
		case consts.UpdatesStatusOK:
			lp.Session.TS = updates.TS

			for _, upd := range updates.Updates {
				dest <- upd
			}
		case consts.UpdatesStatusOutdated:
			lp.Session.TS = updates.TS
		case consts.UpdatesStatusExpired:
			if server, err := lp.Bot.GetGroupLongPollServer(lp.Bot.ID); err == nil {
				lp.Session.Key = server.Key
			}
		case consts.UpdatesStatusLost:
			if server, err := lp.Bot.GetGroupLongPollServer(lp.Bot.ID); err == nil {
				lp.Session.Key = server.Key
				lp.Session.TS = server.TS
			}
		case consts.UpdatesStatusInvalid:
			lp.Session.Version = updates.MinVersion
		}
	}
}
