package api

import (
	"fmt"
	"net/url"

	"github.com/aejoy/vkgo/responses"
	"github.com/aejoy/vkgo/types"
	"github.com/aejoy/vkgo/update"
	"github.com/botscommunity/botsgo/pkg/schema"
)

func (api *API) AddChatUser(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.ResultOperationCodeReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"chat_id", "user_id", "visible_messages_count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Result == 1, api.Call("messages.addChatUser", query.Encode(), &reply)
}

func (api *API) AllowMessagesFromGroup(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"key"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.allowMessagesFromGroup", query.Encode(), &reply)
}

func (api *API) CreateChat(properties ...any) (chatID int, users []int, err error) {
	query, reply := make(url.Values), responses.CreateChatReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"group_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"user_ids"}),
		schema.String:       schema.NewType(schema.ParameterNames{"title"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.ChatID, reply.Response.Users, api.Call("messages.createChat", query.Encode(), &reply)
}

func (api *API) Spam(properties ...any) (message responses.DeleteMessage, err error) {
	messages, err := api.Spams(properties...)
	if len(messages) > 0 {
		message = messages[0]
	}

	return message, err
}

func (api *API) Spams(properties ...any) (messages []responses.DeleteMessage, err error) {
	return api.DeleteMessages(append(properties, types.DeleteMessage{Spam: true})...)
}

func (api *API) DeleteMessage(properties ...any) (message responses.DeleteMessage, err error) {
	messages, err := api.DeleteMessages(properties...)
	if len(messages) > 0 {
		message = messages[0]
	}

	return message, err
}

func (api *API) DeleteMessages(properties ...any) (messages []responses.DeleteMessage, err error) {
	query, reply := make(url.Values), responses.DeleteMessages{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"peer_id", "cmids", "delete_for_all", "spam"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"cmids"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Messages, api.Call("messages.delete", query.Encode(), &reply)
}

func (api *API) DeleteChatPhoto(properties ...any) (messageID int, chat responses.Chat, err error) {
	query, reply := make(url.Values), responses.DeleteChatPhotoReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"chat_id", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.MessageID, reply.Response.Chat, api.Call("messages.deleteChatPhoto", query.Encode(), &reply)
}

func (api *API) DeleteChat(properties ...any) (lastDeletedID int, err error) {
	query, reply := make(url.Values), responses.DeleteChatReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "group_id", "offset", "count"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.LastDeletedID, api.Call("messages.deleteConversation", query.Encode(), &reply)
}

func (api *API) DeleteReaction(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "cmid"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.deleteReaction", query.Encode(), &reply)
}

func (api *API) DenyMessagesFromGroup(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.denyMessagesFromGroup", query.Encode(), &reply)
}

func (api *API) EditMessage(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"peer_id", "conversation_message_id"}),
		schema.String:      schema.NewType(schema.ParameterNames{"message", "attachment"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"attachment"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.edit", query.Encode(), &reply)
}

func (api *API) EditChatTitle(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"chat_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"title"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.editChat", query.Encode(), &reply)
}

func (api *API) GetMessage(properties ...any) (message update.Message, users []responses.User, groups []responses.Group, err error) {
	messages, err := api.GetMessages(properties...)
	if len(messages.Messages) > 0 {
		message = messages.Messages[0]
	}

	return message, messages.Users, messages.Groups, err
}

func (api *API) GetMessages(properties ...any) (messages responses.Messages, err error) {
	query, reply := make(url.Values), responses.MessagesReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"peer_id", "cmids", "group_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"cmids"}),
		schema.ArrayString:  schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	if err = api.Call("messages.getById", query.Encode(), &reply); err != nil {
		return
	}

	for _, message := range reply.Response.Messages {
		if payload, err := update.DecodePayload(message.RawPayload); err == nil {
			message.Payload = payload
		}
	}

	return reply.Response, err
}

func (api *API) GetChatMembers(properties ...any) (members responses.ChatMembers, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.ChatMembersReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:     schema.NewType(schema.ParameterNames{"peer_id", "offset", "group_id"}),
		schema.ArrayString: schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:      nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.getConversationMembers", query.Encode(), &reply)
}

func (api *API) GetMyChats(properties ...any) (chats responses.Chats, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.ChatsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"peer_ids", "group_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"peer_ids"}),
		schema.ArrayString:  schema.NewType(schema.ParameterNames{"fields"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.getConversations", query.Encode(), &reply)
}

func (api *API) GetChat(properties ...any) (chat responses.Conversation, users []responses.User, groups []responses.Group, err error) {
	chats, err := api.GetChats(properties...)
	if len(chats.Chats) > 0 {
		chat = chats.Chats[0]
	}

	return chat, chats.Users, chats.Groups, err
}

func (api *API) GetChats(properties ...any) (chats responses.Chats, err error) {
	query, reply := url.Values{
		"extended": {"1"},
	}, responses.ChatsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"peer_ids", "group_id"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"peer_ids"}),
		schema.ArrayString:  schema.NewType(schema.ParameterNames{"fields"}),
		schema.Boolean:      schema.NewType(schema.ParameterNames{"extended"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.getConversationsById", query.Encode(), &reply)
}

func (api *API) GetIntentUsers(properties ...any) (chats responses.Chats, err error) {
	query, reply := make(url.Values), responses.ChatsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"subscribe_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"intent", "fields"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.getIntentUsers", query.Encode(), &reply)
}

func (api *API) GetChatLink(properties ...any) (link string, err error) {
	query, reply := make(url.Values), responses.ChatLinkReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "reset", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Link, api.Call("messages.getInviteLink", query.Encode(), &reply)
}

func (api *API) GetLastActivity(properties ...any) (online bool, time int, err error) {
	query, reply := make(url.Values), responses.LastActivityReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"user_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Online == 1, reply.Response.Time, api.Call("messages.getLastActivity", query.Encode(), &reply)
}

func (api *API) GetUserLongPollServer(properties ...any) (server responses.LongPollServer, err error) {
	query, reply := make(url.Values), responses.LongPollServerReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"lp_version", "need_pts", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.getLongPollServer", query.Encode(), &reply)
}

func (api *API) GetMessagesReactions(properties ...any) (reactions []responses.MessagesReaction, err error) {
	query, reply := make(url.Values), responses.MessagesReactionsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"peer_id", "cmids"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"cmids"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Items, api.Call("messages.getMessagesReactions", query.Encode(), &reply)
}

func (api *API) GetReactionaryMessages(properties ...any) (counters []responses.ReactionaryMessagesCounters, reactions []responses.ReactionaryMessagesReactions, err error) {
	query, reply := make(url.Values), responses.ReactionaryMessagesReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "cmid", "reaction_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Counters, reply.Response.Reactions, api.Call("messages.getReactedPeers", query.Encode(), &reply)
}

func (api *API) GetReactionsAssets(properties ...any) (version int, reactionIDs []int, assets []responses.ReactionsAsset, err error) {
	query, reply := make(url.Values), responses.ReactionsAssetsReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"client_version"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.Version, reply.Response.ReactionIDs, reply.Response.Assets, api.Call("messages.getReactionsAssets", query.Encode(), &reply)
}

func (api *API) IsMessagesFromGroupAllowed(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.IsMessagesFromGroupAllowedReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"group_id", "user_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.IsAllowed == 1, api.Call("messages.isMessagesFromGroupAllowed", query.Encode(), &reply)
}

func (api *API) JoinChatByInviteLink(properties ...any) (chatID int, err error) {
	query, reply := make(url.Values), responses.JoinChatByInviteLinkReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.String: schema.NewType(schema.ParameterNames{"link"}),
		schema.Struct: nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response.ChatID, api.Call("messages.joinChatByInviteLink", query.Encode(), &reply)
}

func (api *API) MarkChatAsAnswered(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "answered", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.markAsAnsweredConversation", query.Encode(), &reply)
}

func (api *API) PinMessage(properties ...any) (message responses.PinMessage, err error) {
	query, reply := make(url.Values), responses.PinMessageReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "conversation_message_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.pin", query.Encode(), &reply)
}

func (api *API) KickUser(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"chat_id", "member_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.removeChatUser", query.Encode(), &reply)
}

func (api *API) RestoreMessage(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"message_id", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.restore", query.Encode(), &reply)
}

func (api *API) SendMessage(properties ...any) (message responses.SendMessage, err error) {
	sent, err := api.SendMessages(properties...)
	if len(sent) > 0 {
		message = sent[0]
	}

	return message, err
}

func (api *API) SendMessages(properties ...any) (messages []responses.SendMessage, err error) {
	query, reply := url.Values{
		"random_id": {fmt.Sprint(types.NewRandomID())},
	}, responses.SendMessageReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:      schema.NewType(schema.ParameterNames{"peer_ids"}),
		schema.String:       schema.NewType(schema.ParameterNames{"message", "attachment"}),
		schema.ArrayInteger: schema.NewType(schema.ParameterNames{"peer_ids"}),
		schema.ArrayString:  schema.NewType(schema.ParameterNames{"attachment"}),
		schema.Struct:       nil,
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.send", query.Encode(), &reply)
}

func (api *API) SendMessageEventAnswer(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.SendEventMessage{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "user_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"event_id", "event_data"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.sendMessageEventAnswer", query.Encode(), &reply)
}

func (api *API) SendReaction(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.SendEventMessage{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "cmid", "reaction_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.sendReaction", query.Encode(), &reply)
}

func (api *API) SetActivity(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.SendEventMessage{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "group_id"}),
		schema.String:  schema.NewType(schema.ParameterNames{"type"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.setActivity", query.Encode(), &reply)
}

func (api *API) UnpinMessage(properties ...any) (ok bool, err error) {
	query, reply := make(url.Values), responses.OperationCode{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer: schema.NewType(schema.ParameterNames{"peer_id", "group_id"}),
		schema.Struct:  nil,
	}).ConvertToQuery(query, properties...)

	return reply.Code == 1, api.Call("messages.unpin", query.Encode(), &reply)
}

// without documentation

func (api *API) MuteUser(properties ...any) (mute responses.MuteUser, err error) {
	query, reply := make(url.Values), responses.MuteUserReply{}

	schema.NewSchema(schema.TypeDefs{
		schema.Integer:  schema.NewType(schema.ParameterNames{"peer_id", "member_ids"}),
		schema.Duration: schema.NewType(schema.ParameterNames{"for"}),
		schema.String:   schema.NewType(schema.ParameterNames{"action"}),
	}).ConvertToQuery(query, properties...)

	return reply.Response, api.Call("messages.changeConversationMemberRestrictions", query.Encode(), &reply)
}
