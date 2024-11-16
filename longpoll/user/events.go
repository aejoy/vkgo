package user

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/scene"
)

type Update func(_ *api.API, _ []any, _ scene.Scenes) string

func GetUpdate(code float64) Update {
	updates := map[float64]Update{
		4: message,
	}

	return updates[code]
}
