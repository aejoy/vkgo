package user

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/scene"
	"github.com/aejoy/vkgo/types"
)

func message(bot *api.API, messages []any, scenes scene.Scenes) string {
	if len(messages) >= 1 && scenes.MessageFunc != nil {
		if messageID, ok := messages[1].(float64); ok {
			message, _, _, err := bot.GetMessage(types.Message{
				MessageIDs: []int{int(messageID)},
			})
			if err != nil {
				bot.Logger.Error(err.Error())
				return ""
			}

			scenes.MessageFunc(bot, message)
		}

		return "message_new"
	}

	return ""
}
