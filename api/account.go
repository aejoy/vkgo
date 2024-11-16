package api

import (
	"fmt"
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) Ban(ownerID int) (bool, error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("account.ban", fmt.Sprintf("owner_id=%d", ownerID), &reply)
}

func (api *API) ChangePassword(properties ...any) (string, error) {
	query, reply := make(url.Values), responses.ChangePasswordReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.String: schema.NewType(schema.ParameterNames{"old_password", "new_password"}),
		schema.Struct: nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Token, api.Call("account.changePassword", query.Encode(), &reply)
}

func (api *API) GetActiveOffers(properties ...any) ([]int, error) {
	query, reply := make(url.Values), responses.ActiveOffersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"offset", "count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Offers, api.Call("account.getActiveOffers", query.Encode(), &reply)
}

func (api *API) GetAppPermissions(userID int) (int, error) {
	reply := responses.AppPermissionsReply{}
	return reply.Response, api.Call("account.getAppPermissions", fmt.Sprintf("user_id=%d", userID), &reply)
}

func (api *API) GetBans(properties ...any) ([]int, []responses.User, error) {
	query, reply := make(url.Values), responses.BansReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"offset", "count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Bans, reply.Response.Users, api.Call("account.getBanned", query.Encode(), &reply)
}

func (api *API) GetCounters() (responses.Counters, error) {
	reply := responses.CountersReply{}
	return reply.Response, api.Call("account.getCounters", "", &reply)
}

func (api *API) GetInfo(properties ...any) (responses.Info, error) {
	query, reply := make(url.Values), responses.InfoReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("account.getInfo", query.Encode(), &reply)
}

func (api *API) GetProfileInfo() (info responses.ProfileInfo, err error) {
	reply := responses.ProfileInfoReply{}
	return reply.Response, api.Call("account.getProfileInfo", "", &reply)
}

func (api *API) EditProfileInfo(properties ...any) (int, responses.NameRequest, error) {
	query, reply := make(url.Values), responses.EditProfileInfo{}

	schema.NewSchema(schema.TypeDefs{
		schema.Struct: nil,
	}).ConvertToQuery(query, properties...)

	return reply.Changed, reply.NameRequest, api.Call("account.saveProfileInfo", query.Encode(), &reply)
}

func (api *API) EditAccountInfo(properties ...any) (bool, error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Struct: nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("account.setInfo", query.Encode(), &reply)
}

func (api *API) SetOffline() (bool, error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("account.setOffline", "", &reply)
}

func (api *API) SetOnline(properties ...any) (bool, error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Struct: nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("account.setOnline", query.Encode(), &reply)
}

func (api *API) Unban(ownerID int) (bool, error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("account.unban", fmt.Sprintf("owner_id=%d", ownerID), &reply)
}
