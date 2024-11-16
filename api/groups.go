package api

import (
	"fmt"
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) AddGroupAddress(properties ...any) (addGroupAddress responses.AddGroupAddress, err error) {
	query, reply := make(url.Values), responses.AddGroupAddressReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "country_id", "city_id", "metro_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"title", "address", "latitude", "longitude", "phone", "metro_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("groups.addAddress", query.Encode(), &reply)
}

func (api *API) AddGroupCallbackServer(properties ...any) (addGroupAddress responses.AddGroupCallbackServer, err error) {
	query, reply := make(url.Values), responses.AddGroupCallbackServerReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"url", "title", "secret_key"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("groups.addCallbackServer", query.Encode(), &reply)
}

func (api *API) AddGroupLink(properties ...any) (addGroupAddress responses.AddGroupAddress, err error) {
	query, reply := make(url.Values), responses.AddGroupAddressReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"link", "text"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("groups.addLink", query.Encode(), &reply)
}

func (api *API) ApproveRequestGroup(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "user_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.approveRequest", query.Encode(), &reply)
}

func (api *API) BanUser(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "owner_id", "end_date", "comment_visible"}),
		schema.String:  schema.NewType(schema.ParameterNames{"comment"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.ban", query.Encode(), &reply)
}

func (api *API) CreateGroup(properties ...any) (group responses.Group, err error) {
	query, reply := make(url.Values), responses.CreateGroup{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"subtype", "public_category", "public_subcategory"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type", "title", "description"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("groups.create", query.Encode(), &reply)
}

func (api *API) DeleteGroupAddress(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "address_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.deleteAddress", query.Encode(), &reply)
}

func (api *API) DeleteGroupCallbackServer(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "server_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.deleteCallbackServer", query.Encode(), &reply)
}

func (api *API) DeleteGroupLink(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "link_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.deleteLink", query.Encode(), &reply)
}

func (api *API) SetGroupOffline(groupID int) (ok bool, err error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("groups.disableOnline", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) EditGroup(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "access", "subject"}),
		schema.String:  schema.NewType(schema.ParameterNames{"title", "description", "screen_name"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.edit", query.Encode(), &reply)
}

func (api *API) EditGroupCallbackServer(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "server_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"url", "title", "secret_key"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.editCallbackServer", query.Encode(), &reply)
}

func (api *API) EditGroupLink(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "link_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"text"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.editLink", query.Encode(), &reply)
}

func (api *API) EditGroupManager(properties ...any) (status string, err error) {
	query, reply := make(url.Values), responses.EditGroupManagerReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "user_id", "is_contact"}),
		schema.String:  schema.NewType(schema.ParameterNames{"role"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Status, api.Call("groups.editManager", query.Encode(), &reply)
}

func (api *API) SetGroupOnline(groupID int) (ok bool, err error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("groups.enableOnline", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) GetUserGroups(properties ...any) (ok []responses.Group, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.UserGroupsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"user_id"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"filter", "fields"}),
	}).ConvertToQuery(query, properties...)

	return reply.Response.Groups, api.Call("groups.get", query.Encode(), &reply)
}

func (api *API) GetGroupAddresses(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"group_id", "offset", "count"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"address_ids", "latitude", "longitude"}),
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.getAddresses", query.Encode(), &reply)
}

func (api *API) GetBannedUsers(properties ...any) (count int, bans []responses.BannedUserResponse, err error) {
	query, reply := make(url.Values), responses.BannedUsersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "offset", "count"}),
		schema.String:  schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Count, reply.Response.Bans, api.Call("groups.getBanned", query.Encode(), &reply)
}

func (api *API) GetGroup(properties ...any) (group responses.Group, err error) {
	groups, err := api.GetGroups(properties...)
	if len(groups) > 0 {
		group = groups[0]
	}

	return group, err
}

func (api *API) GetGroups(properties ...any) (groups []responses.Group, err error) {
	query, reply := make(url.Values), responses.GroupsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"group_ids"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"group_ids"}),
		schema.ArrayString:  schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Groups, api.Call("groups.getById", query.Encode(), &reply)
}

func (api *API) GetGroupCallbackConfirmationCode(groupID int) (confirmationCode string, err error) {
	reply := responses.GroupCallbackConfirmationCodeReply{}
	return reply.Response.Code, api.Call("groups.getCallbackConfirmationCode", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) GetGroupCallbackServers(properties ...any) (callbackServer []responses.GroupCallbackServer, err error) {
	query, reply := make(url.Values), responses.GroupCallbackServersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"group_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"server_ids"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Items, api.Call("groups.getCallbackServers", query.Encode(), &reply)
}

func (api *API) GetGroupCallbackSettings(properties ...any) (version string, events responses.SettingsEvents, err error) {
	query, reply := make(url.Values), responses.GroupCallbackSettingsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "server_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.APIVersion, reply.Response.Events, api.Call("groups.getCallbackSettings", query.Encode(), &reply)
}

func (api *API) GetGroupCatalogInfo(properties ...any) (enabled int, categories []responses.GroupCatalogInfoCategory, subcategories []responses.Subcategory, err error) {
	query, reply := url.Values{
		"extended":      {"1"},
		"subcategories": {"1"},
	}, responses.GroupCatalogInfoReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"subcategories"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Enabled, reply.Response.Categories, reply.Response.Subcategories, api.Call("groups.getCatalogInfo", query.Encode(), &reply)
}

func (api *API) GetGroupInvitedUsers(properties ...any) (enabled int, categories []responses.GroupCatalogInfoCategory, subcategories []responses.Subcategory, err error) {
	query, reply := make(url.Values), responses.GroupCatalogInfoReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"group_id", "offset", "count"}),
		schema.String:      schema.NewType(schema.ParameterNames{"name_case"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Enabled, reply.Response.Categories, reply.Response.Subcategories, api.Call("groups.getInvitedUsers", query.Encode(), &reply)
}

func (api *API) GetUserGroupInvites(properties ...any) (all []responses.Group, users []responses.User, groups []responses.Group, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.GroupInvitesReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"offset", "count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Groups, reply.Response.Users, reply.Response.Groups, api.Call("groups.getInvites", query.Encode(), &reply)
}

func (api *API) GetGroupLongPollServer(groupID int) (server responses.LongPollServer, err error) {
	reply := responses.LongPollServerReply{}
	return reply.Response, api.Call("groups.getLongPollServer", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) GetGroupLongPollSettings(groupID int) (ok bool, version string, events responses.SettingsEvents, err error) {
	reply := responses.LongPollSettingsReply{}
	return reply.Response.IsEnabled, reply.Response.APIVersion, reply.Response.Events, api.Call("groups.getLongPollSettings", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) GetGroupMembers(properties ...any) (ok []responses.GroupMember, nextFrom string, err error) {
	query, reply := url.Values{"fields": {"a,"}}, responses.GroupMembersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"group_id", "offset", "count"}),
		schema.String:      schema.NewType(schema.ParameterNames{"sort"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields", "filter"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Items, reply.Response.NextFrom, api.Call("groups.getMembers", query.Encode(), &reply)
}

func (api *API) GetGroupOnlineStatus(groupID int) (status string, err error) {
	reply := responses.GroupOnlineStatusReply{}
	return reply.Response.Status, api.Call("groups.getOnlineStatus", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) GetGroupRequests(properties ...any) (groups []responses.User, err error) {
	query, reply := url.Values{"fields": {"a,"}}, responses.GroupRequestsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Items, api.Call("groups.getRequests", query.Encode(), &reply)
}

func (api *API) GetGroupSettings(groupID int) (groups responses.GroupSettings, err error) {
	reply := responses.GroupSettingsReply{}
	return reply.Response, api.Call("groups.getSettings", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) GetGroupTokenPermissions() (mask int, permissions []responses.GroupTokenPermission, err error) {
	reply := responses.GroupTokenPermissionsReply{}
	return reply.Response.Mask, reply.Response.Permissions, api.Call("groups.getTokenPermissions", "", &reply)
}

func (api *API) InviteGroup(properties ...any) (code bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "user_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.invite", query.Encode(), &reply)
}

func (api *API) IsGroupMember(properties ...any) (member, canInvite int, err error) {
	query, reply := url.Values{"extended": {"1"}}, responses.IsGroupMemberReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"group_id", "user_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"user_ids"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Member, reply.Response.CanInvite, api.Call("groups.isMember", query.Encode(), &reply)
}

func (api *API) JoinGroup(properties ...any) (code bool, err error) {
	query, reply := url.Values{"extended": {"1"}}, responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"not_sure"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.join", query.Encode(), &reply)
}

func (api *API) LeaveGroup(groupID int) (code bool, err error) {
	reply := responses.OperationCode{}
	return reply.Code == 1, api.Call("groups.leave", fmt.Sprintf("group_id=%d", groupID), &reply)
}

func (api *API) RemoveGroupUser(properties ...any) (code bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "user_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.removeUser", query.Encode(), &reply)
}

func (api *API) ReorderGroupLink(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "link_id", "after"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.reorderLink", query.Encode(), &reply)
}

func (api *API) SearchGroup(properties ...any) (groups []responses.Group, err error) {
	query, reply := make(url.Values), responses.UserGroupsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"country_id", "city_id", "sort"}),
		schema.String:  schema.NewType(schema.ParameterNames{"q", "type"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Groups, api.Call("groups.search", query.Encode(), &reply)
}

func (api *API) SetGroupLongPollSettings(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "enabled"}),
		schema.String:  schema.NewType(schema.ParameterNames{"api_version"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.setLongPollSettings", query.Encode(), &reply)
}

func (api *API) UnbanUser(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "owner_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("groups.unban", query.Encode(), &reply)
}
