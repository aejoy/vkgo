package api

import (
	"fmt"
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetStorage(properties ...any) (storages responses.Storages, err error) {
	query, reply := make(url.Values), responses.StorageReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id"}),
		schema.String:      schema.NewType(schema.ParameterNames{"key"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"keys"}),
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("storage.get", query.Encode(), &reply)
}

func (api *API) GetStorageKeys(userID int) (keys []string, err error) {
	reply := responses.StorageKeysReply{}
	return reply.Response, api.Call("storage.getKeys", fmt.Sprintf("user_id=%d", userID), &reply)
}

func (api *API) SetStorage(properties ...any) (result any, err error) {
	query, reply := make(url.Values), responses.SetStorageReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"key", "value"}),
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("storage.set", query.Encode(), &reply)
}
