package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetGifts(properties ...any) (int, []responses.Gift, error) {
	query, reply := make(url.Values), responses.GiftsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id", "offset", "count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Count, reply.Response.Gifts, api.Call("gifts.get", query.Encode(), &reply)
}
