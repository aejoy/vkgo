package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) Execute(properties ...any) (any, error) {
	query, reply := make(url.Values), responses.Execute{}

	schema.NewSchema(schema.TypeDefs{
		schema.Float:  schema.NewType(schema.ParameterNames{"func_v"}),
		schema.String: schema.NewType(schema.ParameterNames{"code"}),
		schema.Struct: nil,
	}).ConvertToQuery(query, properties...)

	return reply.Result, api.Call("execute", query.Encode(), &reply)
}
