package api

import (
	"fmt"
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) AddFriend(properties ...any) (int, error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id", "follow"}),
		schema.String:  schema.NewType(schema.ParameterNames{"text"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code, api.Call("friends.add", query.Encode(), &reply)
}

func (api *API) AddFriendsList(properties ...any) (int, error) {
	query, reply := make(url.Values), responses.AddFriendsListReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_ids"}),
		schema.String:  schema.NewType(schema.ParameterNames{"name"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.ListID, api.Call("friends.addList", query.Encode(), &reply)
}

func (api *API) AreFriends(properties ...any) ([]responses.AreFriend, error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.AreFriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"user_ids", "need_sign"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"user_ids"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("friends.areFriends", query.Encode(), &reply)
}

func (api *API) DeleteFriend(userID int) (responses.DeleteFriend, error) {
	reply := responses.DeleteFriendReply{}
	return reply.Response, api.Call("friends.delete", fmt.Sprintf("user_id=%d", userID), &reply)
}

func (api *API) DeleteAllFriendsRequests() (bool, error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("friends.deleteAllRequests", "", &reply)
}

func (api *API) DeleteFriendsList(listID int) (bool, error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("friends.deleteList", fmt.Sprintf("list_id=%d", listID), &reply)
}

func (api *API) EditFriends(properties ...any) (bool, error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"list_ids", "user_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"list_ids"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("friends.edit", query.Encode(), &reply)
}

func (api *API) EditFriendsList(properties ...any) (bool, error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"list_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"add_user_ids"}),
		schema.String:       schema.NewType(schema.ParameterNames{"name"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("friends.editList", query.Encode(), &reply)
}

func (api *API) GetFriends(properties ...any) ([]responses.User, error) {
	query, reply := url.Values{
		"fields": {"a,"},
	}, responses.FriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Friends, api.Call("friends.get", query.Encode(), &reply)
}

func (api *API) GetFriendsLists(properties ...any) ([]responses.FriendsListItem, error) {
	query, reply := make(url.Values), responses.FriendsListsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"return_system", "user_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Items, api.Call("friends.getLists", query.Encode(), &reply)
}

func (api *API) GetMutualFriends(properties ...any) ([]responses.MutualFriend, error) {
	query, reply := make(url.Values), responses.MutualFriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"target_uids", "source_uid", "offset"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"target_uids"}),
		schema.String:       schema.NewType(schema.ParameterNames{"order"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("friends.getMutual", query.Encode(), &reply)
}

func (api *API) GetOnlineFriends(properties ...any) ([]int, []int, error) {
	query, reply := url.Values{
		"online_mobile": {"1"},
	}, responses.OnlineFriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id", "list_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"order"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Online, reply.Response.OnlineMobile, api.Call("friends.getOnline", query.Encode(), &reply)
}

func (api *API) GetRecentFriends(properties ...any) ([]int, error) {
	query, reply := make(url.Values), responses.RecentFriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("friends.getRecent", query.Encode(), &reply)
}

func (api *API) GetFriendsRequests(properties ...any) ([]responses.User, error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.FriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"out"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Friends, api.Call("friends.getRequests", query.Encode(), &reply)
}

func (api *API) GetSuggestionsFriends(properties ...any) ([]responses.User, error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.FriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Integer:     schema.NewType(schema.ParameterNames{"offset", "count"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Friends, api.Call("friends.getSuggestions", query.Encode(), &reply)
}

func (api *API) SearchFriends(properties ...any) ([]responses.User, error) {
	query, reply := make(url.Values), responses.FriendsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id", "offset", "count"}),
		schema.String:      schema.NewType(schema.ParameterNames{"q", "name_case"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Friends, api.Call("friends.search", query.Encode(), &reply)
}
