package types

// AddFriend

type AddFriend struct {
	UserID int    `json:"user_id"`
	Text   string `json:"text"`
	Follow bool   `json:"follow" to:"int"`
}

// AddFriendsList

type AddFriendsList struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"user_ids"`
}

// AreFriends

type AreFriends struct {
	UserIDs  []string `json:"user_ids"`
	NeedSign int      `json:"need_sign"`
}

// EditFriends

type EditFriends struct {
	UserID  int   `json:"user_id"`
	ListIDs []int `json:"list_ids"`
}

// EditFriendsList

type EditFriendsList struct {
	Name          string `json:"name"`
	ListID        int    `json:"list_id"`
	UserIDs       []int  `json:"user_ids"`
	AddUserIDs    []int  `json:"add_user_ids"`
	DeleteUserIDs []int  `json:"delete_user_ids"`
}

// GetFriends

type Friends struct {
	UserID   int      `json:"user_id"`
	Order    string   `json:"order"`
	ListID   int      `json:"list_id"`
	Count    int      `json:"count"`
	Offset   int      `json:"offset"`
	Fields   []string `json:"fields"`
	NameCase string   `json:"name_case"`
	Ref      string   `json:"ref"`
}

// GetFriendsLists

type FriendsList struct {
	UserID       int `json:"user_id"`
	ReturnSystem int `json:"return_system"`
}

// GetMutualFriends

type MutualFriends struct {
	SourceUID       int    `json:"source_uid"`
	TargetUIDs      []int  `json:"target_uids"`
	Order           string `json:"order"`
	Count           int    `json:"count"`
	Offset          int    `json:"offset"`
	NeedCommonCount bool   `json:"need_common_count" to:"int"`
}

// GetOnlineFriends

type OnlineFriends struct {
	UserID int    `json:"user_id"`
	ListID int    `json:"list_id"`
	Order  string `json:"order"`
	Count  int    `json:"count"`
	Offset int    `json:"offset"`
}

// GetRecentFriends

// GetFriendsRequests

type FriendsRequests struct {
	Offset     int    `json:"offset"`
	Count      int    `json:"count"`
	NeedMutual bool   `json:"need_mutual" to:"int"`
	Out        bool   `json:"out" to:"int"`
	Sort       int    `json:"sort"`
	NeedViewed bool   `json:"need_viewed" to:"int"`
	Ref        string `json:"ref"`
	Fields     string `json:"fields"`
}

// GetSuggestionsFriends

type SuggestionsFriends struct {
	Filter   string `json:"filter"`
	Count    int    `json:"count"`
	Offset   int    `json:"offset"`
	Fields   string `json:"fields"`
	NameCase string `json:"name_case"`
}

// SearchFriends

type SearchFriends struct {
	UserID   int    `json:"user_id"`
	Q        string `json:"q"`
	Fields   string `json:"fields"`
	NameCase string `json:"name_case"`
	Offset   int    `json:"offset"`
	Count    int    `json:"count"`
}
