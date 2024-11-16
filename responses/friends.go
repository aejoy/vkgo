package responses

// AddFriendsList

type AddFriendsListReply struct {
	ErrorInterface
	Response AddFriendsList `json:"response,omitempty"`
}

type AddFriendsList struct {
	ListID int `json:"list_id,omitempty"`
}

// AreFriends

type AreFriendsReply struct {
	ErrorInterface
	Response []AreFriend `json:"response,omitempty"`
}

type AreFriend struct {
	UserID          int    `json:"user_id,omitempty"`
	FriendStatus    int    `json:"friend_status,omitempty"`
	IsRequestUnread bool   `json:"is_request_unread,omitempty"`
	Sign            string `json:"sign,omitempty"`
}

// DeleteFriend

type DeleteFriendReply struct {
	ErrorInterface
	Response DeleteFriend `json:"response,omitempty"`
}

type DeleteFriend struct {
	Success           int `json:"success,omitempty"`
	FriendDeleted     int `json:"friend_deleted,omitempty"`
	InRequestDeleted  int `json:"in_request_deleted,omitempty"`
	OutRequestDeleted int `json:"out_request_deleted,omitempty"`
	SuggestionDeleted int `json:"suggestion_deleted,omitempty"`
}

// GetFriends

type FriendsReply struct {
	ErrorInterface
	Response Friends `json:"response,omitempty"`
}

type Friends struct {
	Friends []User `json:"items,omitempty"`
}

// GetFriendsLists

type FriendsListsReply struct {
	ErrorInterface
	Response FriendsLists `json:"response,omitempty"`
}

type FriendsLists struct {
	Count int               `json:"count,omitempty"`
	Items []FriendsListItem `json:"items,omitempty"`
}

type FriendsListItem struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetMutualFriends

type MutualFriendsReply struct {
	ErrorInterface
	Response []MutualFriend `json:"response,omitempty"`
}

type MutualFriend struct {
	CommonCount   int   `json:"common_count,omitempty"`
	CommonFriends []int `json:"common_friends,omitempty"`
	ID            int   `json:"id,omitempty"`
}

// GetOnlineFriends

type OnlineFriendsReply struct {
	ErrorInterface
	Response OnlineFriends `json:"response,omitempty"`
}

type OnlineFriends struct {
	Online       []int `json:"online,omitempty"`
	OnlineMobile []int `json:"online_mobile,omitempty"`
}

// GetRecentFriends

type RecentFriendsReply struct {
	ErrorInterface
	Response []int `json:"response,omitempty"`
}
