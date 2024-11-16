package responses

import "github.com/aejoy/vkgo/update"

type CreateChatReply struct {
	ErrorInterface
	Response CreateChat ` json:"response"`
}

type CreateChat struct {
	ChatID int   `json:"chat_id,omitempty"`
	Users  []int `json:"peer_ids,omitempty"`
}

type DeleteMessages struct {
	ErrorInterface
	Messages []DeleteMessage `json:"response,omitempty"`
}

type DeleteMessage struct {
	ChatID        int `json:"peer_id,omitempty"`
	MessageID     int `json:"message_id,omitempty"`
	ChatMessageID int `json:"conversation_message_id,omitempty"`
	Response      int `json:"response,omitempty"`
}

type DeleteChatPhotoReply struct {
	ErrorInterface
	Response DeleteChatPhoto `json:"response,omitempty"`
}

type DeleteChatPhoto struct {
	MessageID int  `json:"message_id,omitempty"`
	Chat      Chat `json:"chat,omitempty"`
}

type Chat struct {
	ID           int          `json:"id,omitempty"`
	Type         string       `json:"type,omitempty"`
	Title        string       `json:"title,omitempty"`
	AdminID      int          `json:"admin_id,omitempty"`
	Users        []int        `json:"users,omitempty"`
	PushSettings PushSettings `json:"push_settings,omitempty"`
	Photo50      string       `json:"photo_50,omitempty"`
	Photo100     string       `json:"photo_100,omitempty"`
	Photo200     string       `json:"photo_200,omitempty"`
	PhotoBase    string       `json:"photo_base,omitempty"`
	Left         int          `json:"left,omitempty"`
	Kicked       int          `json:"kicked,omitempty"`

	// without documentation
	MembersCount   int  `json:"members_count,omitempty"`
	IsDefaultPhoto bool `json:"is_default_photo,omitempty"`
}

type DeleteChatReply struct {
	ErrorInterface
	Response DeleteChat `json:"response,omitempty"`
}

type DeleteChat struct {
	LastDeletedID int `json:"last_deleted_id,omitempty"`
}

type ChatMessagesReply struct {
	ErrorInterface
	Response ChatMessages `json:"response,omitempty"`
}

type ChatMessages struct {
	Messages []update.Message `json:"items,omitempty"`
}

type MessagesReply struct {
	ErrorInterface
	Response Messages `json:"response,omitempty"`
}

type Messages struct {
	Messages []update.Message `json:"items,omitempty"`
	Users    []User           `json:"profiles,omitempty"`
	Groups   []Group          `json:"groups,omitempty"`
}

type ChatMembersReply struct {
	ErrorInterface
	Response ChatMembers `json:"response,omitempty"`
}

type ChatMembers struct {
	Settings ChatRestrictions `json:"chat_restrictions,omitempty"`
	Members  []ChatMember     `json:"items,omitempty"`
	Users    []User           `json:"profiles,omitempty"`
	Groups   []Group          `json:"groups,omitempty"`
	Contacts []Contact        `json:"contacts,omitempty"`
}

type ChatRestrictions struct {
	AdminsPromoteUsers bool `json:"admins_promote_users,omitempty"`
	OnlyAdminsEditInfo bool `json:"only_admins_edit_info,omitempty"`
	OnlyAdminsEditPin  bool `json:"only_admins_edit_pin,omitempty"`
	OnlyAdminsInvite   bool `json:"only_admins_invite,omitempty"`
	OnlyAdminsKick     bool `json:"only_admins_kick,omitempty"`
}

type ChatMember struct {
	ID        int  `json:"member_id,omitempty"`
	InvitedBy int  `json:"invited_by,omitempty"`
	IsOwner   bool `json:"is_owner,omitempty"`
	IsAdmin   bool `json:"is_admin,omitempty"`
	CanKick   bool `json:"can_kick,omitempty"`
	JoinDate  int  `json:"join_date,omitempty"`
}

type ChatsReply struct {
	ErrorInterface
	Response Chats ` json:"response,omitempty"`
}

type Chats struct {
	Chats  []Conversation `json:"items,omitempty"`
	Users  []User         `json:"profiles,omitempty"`
	Groups []Group        `json:"groups,omitempty"`
}

type Conversation struct {
	ErrorInterface
	Chat                 ConversationPeer `json:"peer,omitempty"`
	InRead               int              `json:"in_read,omitempty"`
	OutRead              int              `json:"out_read,omitempty"`
	UnreadCount          int              `json:"unread_count,omitempty"`
	Important            bool             `json:"important,omitempty"`
	Unanswered           bool             `json:"unanswered,omitempty"`
	PushSettings         PushSettings     `json:"push_settings,omitempty"`
	CanWrite             CanWriteChat     `json:"can_write,omitempty"`
	ChatSettings         ChatSettings     `json:"chat_settings,omitempty"`
	SortID               ChatSortID       `json:"sort_id,omitempty"`
	LastMessageID        int              `json:"last_message_id,omitempty"`
	LastChatMessageID    int              `json:"last_conversation_message_id,omitempty"`
	InReadChatMessageID  int              `json:"in_read_cmid,omitempty"`
	OutReadChatMessageID int              `json:"out_read_cmid,omitempty"`
	MarkedUnread         bool             `json:"is_marked_unread,omitempty"`
	CanSendMoney         bool             `json:"can_send_money,omitempty"`
	CanReceiveMoney      bool             `json:"can_receive_money,omitempty"`
}

type ConversationPeer struct {
	ID      int    `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	LocalID int    `json:"local_id,omitempty"`
}

type ChatSortID struct {
	MajorID int `json:"major_id,omitempty"`
	MinorID int `json:"minor_id,omitempty"`
}

type CanWriteChat struct {
	Allowed bool `json:"allowed,omitempty"`
	Reason  int  `json:"reason,omitempty"`
}

type ChatSettings struct {
	MembersCount   int                `json:"members_count,omitempty"`
	Title          string             `json:"title,omitempty"`
	PinnedMessage  PinMessage         `json:"pinned_message,omitempty"`
	State          string             `json:"state,omitempty"`
	Photo          update.PhotoAction `json:"photo,omitempty"`
	ActivesID      []int              `json:"active_ids,omitempty"`
	IsGroupChannel bool               `json:"is_group_channel,omitempty"`
}

type IntentUsersReply struct {
	ErrorInterface
	Response IntentUsers `json:"response"`
}

type IntentUsers struct {
	Items []int  `json:"items,omitempty"`
	Users []User `json:"profiles,omitempty"`
}

type ChatLinkReply struct {
	ErrorInterface
	Response ChatLink ` json:"response,omitempty"`
}

type ChatLink struct {
	Link string `json:"link,omitempty"`
}

type LastActivityReply struct {
	ErrorInterface
	Response LastActivity `json:"response,omitempty"`
}

type LastActivity struct {
	Online int `json:"online,omitempty"`
	Time   int `json:"time,omitempty"`
}

type MessagesReactionsReply struct {
	ErrorInterface
	Response MessagesReactions `json:"response,omitempty"`
}

type MessagesReactions struct {
	Items []MessagesReaction `json:"items,omitempty"`
}

type MessagesReaction struct {
	ChatMessageID int                      `json:"cmid,omitempty"`
	Counters      []MessageReactionCounter `json:"counters,omitempty"`
}

type MessageReactionCounter struct {
	ReactionID int   `json:"reaction_id,omitempty"`
	Count      int   `json:"count,omitempty"`
	UserIDs    []int `json:"user_ids,omitempty"`
}

type ReactionaryMessagesReply struct {
	ErrorInterface
	Response ReactionaryMessages `json:"response,omitempty"`
}

type ReactionaryMessages struct {
	Reactions []ReactionaryMessagesReactions `json:"reactions,omitempty"`
	Counters  []ReactionaryMessagesCounters  `json:"counters"`
}

type ReactionaryMessagesReactions struct {
	UserID     int `json:"user_id,omitempty"`
	ReactionID int `json:"reaction_id,omitempty"`
}

type ReactionaryMessagesCounters struct {
	ReactionID int   `json:"reaction_id,omitempty"`
	Count      int   `json:"count,omitempty"`
	UserIDs    []int `json:"user_ids,omitempty"`
}

type ReactionsAssetsReply struct {
	ErrorInterface
	Response ReactionsAssets `json:"response,omitempty"`
}

type ReactionsAssets struct {
	Version     int              `json:"version,omitempty"`
	Assets      []ReactionsAsset `json:"assets,omitempty"`
	ReactionIDs []int            `json:"reaction_ids,omitempty"`
}

type ReactionsAsset struct {
	ReactionID int                `json:"reaction_id,omitempty"`
	Links      ReactionsAssetLink `json:"links,omitempty"`
}

type ReactionsAssetLink struct {
	BigAnimation   string `json:"big_animation,omitempty"`
	SmallAnimation string `json:"small_animation,omitempty"`
	Static         string `json:"static,omitempty"`
}

type IsMessagesFromGroupAllowedReply struct {
	ErrorInterface
	Response IsMessagesFromGroupAllowed `json:"response,omitempty"`
}

type IsMessagesFromGroupAllowed struct {
	IsAllowed int `json:"is_allowed,omitempty"`
}

type JoinChatByInviteLinkReply struct {
	ErrorInterface
	Response JoinChatByInviteLink `json:"response,omitempty"`
}

type JoinChatByInviteLink struct {
	ChatID int `json:"chat_id,omitempty"`
}

type PinMessageReply struct {
	ErrorInterface
	Response PinMessage `json:"response"`
}

type PinMessage struct {
	ID          int                 `json:"id,omitempty"`
	Date        int                 `json:"date,omitempty"`
	UserID      int                 `json:"from_id,omitempty"`
	Text        string              `json:"text,omitempty"`
	Attachments []update.Attachment `json:"attachments,omitempty"`
	Geolocation update.Geolocation  `json:"geo,omitempty"`
	Forwards    []string            `json:"fwd_messages,omitempty"`
}

type MuteUserReply struct {
	ErrorInterface
	Response MuteUser `json:"response"`
}

type MuteUser struct {
	FailedMemberIDs []int `json:"failed_member_ids"`
}
