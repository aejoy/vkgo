package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetHistoryMessages(properties ...any) (historyMessages responses.HistoryMessages, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.HistoryMessagesReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"peer_id", "offset", "count"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Boolean:     schema.NewType(schema.ParameterNames{"rev"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.getHistory", query.Encode(), &reply)
}
