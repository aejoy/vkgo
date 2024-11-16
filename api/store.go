package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetStickers(properties ...any) (stickers []responses.Sticker, err error) {
	query, reply := url.Values{
		"type":    {"stickers"},
		"filters": {"purchased"},
	}, responses.StickersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"filters"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Stickers, api.Call("store.getProducts", query.Encode(), &reply)
}
