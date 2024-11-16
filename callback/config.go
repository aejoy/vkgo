package callback

import "github.com/aejoy/vkgo/api"

func (cb *Callback) Add(bot *api.API) bool {
	cb.Bots[bot.ID] = bot
	return cb.Bots[bot.ID] != nil
}

func (cb *Callback) Remove(bot *api.API) bool {
	delete(cb.Bots, bot.ID)
	return cb.Bots[bot.ID] == nil
}
