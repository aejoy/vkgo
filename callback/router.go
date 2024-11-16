package callback

import (
	"encoding/json"

	"github.com/aejoy/vkgo/scene"
	"github.com/aejoy/vkgo/update"
	fiber "github.com/gofiber/fiber/v2"
)

type Router struct {
	Config *Callback
}

func (r *Router) Fiber(ctx *fiber.Ctx) error {
	updates := update.Update{}
	if err := json.Unmarshal(ctx.Body(), &updates); err != nil {
		return err
	}

	if bot := r.Config.Bots[updates.GroupID]; bot != nil || r.Config.everyone {
		if updates.Type == "confirmation" {
			return ctx.SendString(ctx.Params("confirmation"))
		}

		scene.Use(bot, updates, *r.Config.Scenes)

		return ctx.SendString("ok")
	}

	return ctx.SendString("unregistered bot or access is denied to outsiders")
}
