package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetStatus(properties ...any) (text string, err error) {
	query, reply := make(url.Values), responses.StatusReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Text, api.Call("status.get", query.Encode(), &reply)
}

func (api *API) SetStatus(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"text"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("status.set", query.Encode(), &reply)
}
