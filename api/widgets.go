package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) UpdateWidgets(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.UpdateWidgetsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.String: schema.NewType(schema.ParameterNames{"type", "code"}),
	}).ConvertToQuery(query, properties...)

	return reply.Response == 1, api.Call("appWidgets.update", query.Encode(), &reply)
}
