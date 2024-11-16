package types

import (
	"encoding/json"
	"time"
)

type Message struct {
	GroupID        int      `json:"group_id"`
	ChatID         int      `json:"chat_id"`
	MessageIDs     []int    `json:"message_ids"`
	ChatMessageIDs []int    `json:"cmids"`
	Fields         []string `json:"fields"`
	Extended       bool     `json:"extended"`
	Characters     int      `json:"characters"`
}

type UserLongPollServer struct {
	GroupID int  `json:"group_id"`
	Version int  `json:"lp_version"`
	NeedPTS bool `json:"need_pts"`
}

type SendMessage struct {
	GroupID         int      `json:"group_id"`
	RandomID        int      `json:"random_id"`
	Domain          string   `json:"domain"`
	ChatID          int      `json:"peer_ids"`
	ChatIDs         []int    `json:"peer_ids"` //nolint:govet
	UserID          int      `json:"user_id"`
	UserIDs         []int    `json:"user_ids"`
	Text            string   `json:"message"`
	Latitude        string   `json:"lat"`
	Longitude       string   `json:"long"`
	Attachments     []string `json:"attachment"`
	Attachment      string   `json:"attachment"` //nolint:govet
	Forward         string   `json:"forward"`
	StickerID       int      `json:"sticker_id"`
	Keyboard        string   `json:"keyboard"`
	Template        string   `json:"template"`
	Payload         string   `json:"payload"`
	ContentSource   string   `json:"content_source"`
	DontParseLinks  bool     `json:"dont_parse_links" to:"int"`
	DisableMentions bool     `json:"disable_mentions" to:"int"`
	Intent          string   `json:"intent"`
	SubscribeID     int      `json:"subscribe_id"`
	ExpireTime      int      `json:"expire_ttl"`
}

type Forward struct {
	ChatID        int  `json:"peer_id"`
	ChatMessageID int  `json:"conversation_message_ids"`
	IsReply       bool `json:"is_reply"`
}

func NewForward(chatID, chatMessageID int, isReply bool) (string, error) {
	b, err := json.Marshal(Forward{chatID, chatMessageID, isReply})
	return string(b), err
}

func NewRandomID() int {
	return int(time.Now().UnixNano())
}

type EditMessage struct {
	GroupID             int      `json:"group_id"`
	ChatID              int      `json:"peer_id"`
	Text                string   `json:"message"`
	MessageID           int      `json:"message_id"`
	ChatMessageID       int      `json:"conversation_message_id"`
	Attachment          string   `json:"attachment"`
	Attachments         []string `json:"attachment,omitempty"` //nolint:govet
	Template            string   `json:"template,omitempty"`
	Keyboard            string   `json:"keyboard,omitempty"`
	KeepForwardMessages bool     `json:"keep_forward_messages,omitempty"`
	KeepSnippets        bool     `json:"keep_snippets,omitempty"`
	DontParseLinks      bool     `json:"dont_parse_links,omitempty"`
	DisableMentions     bool     `json:"disable_mentions,omitempty"`
	Latitude            string   `json:"latitude,omitempty"`
	Longitude           string   `json:"longitude,omitempty"`
}

type DeleteMessage struct {
	GroupID         int   `json:"group_id"`
	ChatID          int   `json:"peer_id"`
	MessageID       int   `json:"messages_ids"`
	MessagesIDs     []int `json:"messages_ids"` //nolint:govet
	ChatMessageID   int   `json:"cmids"`
	ChatMessagesIDs []int `json:"cmids"` //nolint:govet
	Spam            bool  `json:"spam"`
	Everyone        bool  `json:"delete_for_all"`
	Reason          int   `json:"reason"`
}

type KickUser struct {
	ChatID   int `json:"chat_id"`
	UserID   int `json:"user_id"`
	MemberID int `json:"member_id"`
}

type PinMessage struct {
	ChatID        int `json:"peer_id"`
	MessageID     int `json:"message_id"`
	ChatMessageID int `json:"conversation_message_id"`
}

type CreateChat struct {
	GroupID int    `json:"group_id"`
	UserID  int    `json:"user_ids"`
	UserIDs []int  `json:"user_ids"` //nolint:govet
	Title   string `json:"title"`
}

type ChatMembers struct {
	GroupID int      `json:"group_id"`
	ChatID  int      `json:"peer_id"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Fields  []string `json:"fields"`
}
