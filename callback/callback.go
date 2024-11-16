package callback

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/scene"
)

// Callback 🌍 is a configuration structure of bots, scenes, routers, and "everyone".
type Callback struct {
	Bots     map[int]*api.API
	Scenes   *scene.Scenes
	Router   Router
	everyone bool
}

// New 🌠 returns an instance of *Callback
// Accepts 🤖 *api.API or 🎪 *scene.Scenes arguments.
func New(properties ...any) *Callback {
	cb := &Callback{
		Bots:   make(map[int]*api.API),
		Scenes: scene.NewScene(),
	}

	cb.Router = Router{
		Config: cb,
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *api.API:
			cb.Bots[property.ID] = property
		case *scene.Scenes:
			cb.Scenes = property
		}
	}

	return cb
}
