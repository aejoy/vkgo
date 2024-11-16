package api

import (
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) GetUser(properties ...any) (user responses.User, err error) {
	users, err := api.GetUsers(properties...)
	if len(users) > 0 {
		user = users[0]
	}

	return user, err
}

func (api *API) GetUsers(properties ...any) (user []responses.User, err error) {
	query, reply := make(url.Values), responses.UsersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"user_ids"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"user_ids"}),
		schema.String:       schema.NewType(schema.ParameterNames{"user_ids", "name_case"}),
		schema.ArrayString:  schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("users.get", query.Encode(), &reply)
}

func (api *API) GetFollowers(properties ...any) (friendsCount int, followers []responses.User, err error) {
	query, reply := url.Values{
		"fields": {"about"},
	}, responses.FollowersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id"}),
		schema.String:      schema.NewType(schema.ParameterNames{"name_case"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.FriendsCount, reply.Response.Items, api.Call("users.getFollowers", query.Encode(), &reply)
}

func (api *API) GetSubscriptions(properties ...any) (followers []responses.Subscription, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.SubscriptionsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id"}),
		schema.String:      schema.NewType(schema.ParameterNames{"name_case"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Items, api.Call("users.getFollowers", query.Encode(), &reply)
}

func (api *API) ReportUser(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type", "comment"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("users.report", query.Encode(), &reply)
}

func (api *API) SearchUser(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"group_id"}),
		schema.String:      schema.NewType(schema.ParameterNames{"q"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("users.search", query.Encode(), &reply)
}
