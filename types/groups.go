package types

// AddGroupAddress

type AddGroupAddress struct {
	GroupID           int    `json:"group_id"`
	Title             string `json:"title"`
	Address           string `json:"address"`
	AdditionalAddress string `json:"additional_address"`
	CountryID         int    `json:"country_id"`
	CityID            int    `json:"city_id"`
	MetroID           int    `json:"metro_id"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Phone             string `json:"phone"`
	WorkInfoStatus    string `json:"work_info_status"`
	Timetable         string `json:"timetable"`
	IsMainAddress     bool   `json:"is_main_address" to:"int"`
}

// AddGroupCallbackServer

type AddGroupCallbackServer struct {
	GroupID   int    `json:"group_id"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	SecretKey string `json:"secret_key"`
}

// AddGroupLink

type AddGroupLink struct {
	GroupID int    `json:"group_id"`
	Link    string `json:"link"`
	Text    string `json:"text"`
}

// ApproveRequestGroup

type ApproveRequestGroup struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

// BanUser

type BanUser struct {
	GroupID        int    `json:"group_id"`
	OwnerID        int    `json:"owner_id"`
	EndDate        int    `json:"end_date"`
	Reason         int    `json:"reason"`
	Comment        string `json:"comment"`
	CommentVisible bool   `json:"comment_visible" to:"int"`
}

// CreateGroup

type CreateGroup struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Type              string `json:"type"`
	PublicCategory    int    `json:"public_category"`
	PublicSubcategory int    `json:"public_subcategory"`
	Subtype           int    `json:"subtype"`
}

// DeleteGroupAddress

type DeleteGroupAddress struct {
	GroupID   int `json:"group_id"`
	AddressID int `json:"address_id"`
}

// DeleteGroupCallbackServer

type DeleteGroupCallbackServer struct {
	GroupID  int `json:"group_id"`
	ServerID int `json:"server_id"`
}

// DeleteGroupLink

type DeleteGroupLink struct {
	GroupID int `json:"group_id"`
	LinkID  int `json:"link_id"`
}

// EditGroup

type EditGroup struct {
	GroupID                  int    `json:"group_id"`
	Title                    string `json:"title"`
	Description              string `json:"description"`
	ScreenName               string `json:"screen_name"`
	Access                   int    `json:"access"`
	Website                  string `json:"website"`
	Subject                  int    `json:"subject"`
	EMail                    string `json:"email"`
	Phone                    string `json:"phone"`
	RSS                      string `json:"rss"`
	EventStartDate           int    `json:"event_start_date"`
	EventFinishDate          int    `json:"event_finish_date"`
	EventGroupID             int    `json:"event_group_id"`
	PublicCategory           int    `json:"public_category"`
	PublicSubcategory        int    `json:"public_subcategory"`
	PublicDate               string `json:"public_date"`
	Wall                     int    `json:"wall"`
	Topics                   int    `json:"topics"`
	Photos                   int    `json:"photos"`
	Video                    int    `json:"video"`
	Audio                    int    `json:"audio"`
	Links                    bool   `json:"links" to:"int"`
	Events                   bool   `json:"events" to:"int"`
	Places                   bool   `json:"places" to:"int"`
	Contacts                 bool   `json:"contacts" to:"int"`
	Docs                     int    `json:"docs"`
	Wiki                     int    `json:"wiki"`
	Messages                 bool   `json:"messages" to:"int"`
	Articles                 bool   `json:"articles" to:"int"`
	Addresses                bool   `json:"addresses" to:"int"`
	AgeLimits                int    `json:"age_limits"`
	Market                   bool   `json:"market" to:"int"`
	MarketButtons            string `json:"market_buttons"`
	MarketComments           bool   `json:"market_comments" to:"int"`
	MarketCountry            string `json:"market_country"`
	MarketCity               string `json:"market_city"`
	MarketCurrency           int    `json:"market_currency"`
	MarketContact            int    `json:"market_contact"`
	MarketWiki               int    `json:"market_wiki"`
	ObsceneFilter            bool   `json:"obscene_filter" to:"int"`
	ObsceneStopwords         bool   `json:"obscene_stopwords" to:"int"`
	ToxicFilter              bool   `json:"toxic_filter" to:"int"`
	DisableRepliesFromGroups bool   `json:"disable_replies_from_groups" to:"int"`
	ObsceneWord              string `json:"obscene_word"`
	MainSection              int    `json:"main_section"`
	SecondarySection         int    `json:"secondary_section"`
	Country                  int    `json:"country"`
	City                     int    `json:"city"`
}

// EditGroupCallbackServer

type EditGroupCallbackServer struct {
	GroupID   int    `json:"group_id"`
	ServerID  int    `json:"server_id"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	SecretKey string `json:"secret_key"`
}

// EditGroupLink

type EditGroupLink struct {
	GroupID int    `json:"group_id"`
	LinkID  int    `json:"link_id"`
	Text    string `json:"text"`
}

// EditGroupManager

type EditGroupManager struct {
	GroupID         int    `json:"group_id"`
	UserID          int    `json:"user_id"`
	Role            string `json:"role"`
	IsContact       bool   `json:"is_contact" to:"int"`
	ContactPosition string `json:"contact_position"`
	ContactPhone    string `json:"contact_phone"`
	ContactEMail    string `json:"contact_email"`
}

// GetUserGroups

type GetUserGroups struct {
	UserID   int      `json:"user_id"`
	Extended bool     `json:"extended" to:"int"`
	Filter   string   `json:"filter"`
	Fields   []string `json:"fields"`
	Offset   int      `json:"offset"`
	Count    int      `json:"count"`
}

// GetGroupAddresses

type GetGroupAddresses struct {
	GroupID    int      `json:"group_id"`
	AddressIDs []any    `json:"address_ids"`
	Latitude   string   `json:"latitude"`
	Longitude  string   `json:"longitude"`
	Offset     int      `json:"offset"`
	Count      int      `json:"count"`
	Fields     []string `json:"fields"`
}

// GetBannedUsers

type GetBannedUsers struct {
	GroupID int      `json:"group_id"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Fields  []string `json:"fields"`
	OwnerID int      `json:"owner_id"`
}

// GetGroup, GetGroups

type GetGroups struct {
	GroupIDs string   `json:"group_ids"`
	GroupID  string   `json:"group_id"`
	Fields   []string `json:"fields"`
}

// GetGroupCallbackServers

type GetGroupCallbackServers struct {
	GroupID   int   `json:"group_id"`
	ServerIDs []any `json:"server_ids"`
}

// GetGroupCallbackSettings

type GetGroupCallbackSettings struct {
	GroupID  int `json:"group_id"`
	ServerID int `json:"server_id"`
}

// GetGroupCatalogInfo

type GetGroupCatalogInfo struct {
	Extended      bool `json:"extended" to:"int"`
	Subcategories bool `json:"subcategories" to:"int"`
}

// GetGroupInvitedUsers

type GetGroupInvitedUsers struct {
	Offset   int  `json:"offset"`
	Count    int  `json:"count"`
	Extended bool `json:"extended" to:"int"`
}

// GetGroupMembers

type GetGroupMembers struct {
	GroupID int      `json:"group_id"`
	Sort    string   `json:"sort"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Fields  []string `json:"fields"`
}

// GetGroupRequests

type GetGroupRequests struct {
	GroupID int      `json:"group_id"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Fields  []string `json:"fields"`
}

// InviteGroup

type InviteGroup struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

// IsGroupMember

type IsGroupMember struct {
	GroupID int   `json:"group_id"`
	UserID  int   `json:"user_id"`
	UserIDs []int `json:"user_i_ds"`
}

// JoinGroup

type JoinGroup struct {
	GroupID int    `json:"group_id"`
	NotSure string `json:"not_sure"`
}

// RemoveGroupUser

type RemoveGroupUser struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

// ReorderGroupLink

type ReorderGroupLink struct {
	GroupID int `json:"group_id"`
	LinkID  int `json:"link_id"`
	After   int `json:"after"`
}

// SearchGroup

type SearchGroup struct {
	Q         string `json:"q"`
	Type      string `json:"type"`
	CountryID int    `json:"country_id"`
	CityID    int    `json:"city_id"`
	Future    bool   `json:"future" to:"int"`
	Market    bool   `json:"market" to:"int"`
	Sort      int    `json:"sort"`
	Offset    int    `json:"offset"`
	Count     int    `json:"count"`
}

// SetGroupLongPollSettings

type SetGroupLongPollSettings struct {
	GroupID                       int    `json:"group_id"`
	Enabled                       bool   `json:"enabled" to:"int"`
	APIVersion                    string `json:"api_version"`
	MessageNew                    bool   `json:"message_new" to:"int"`
	MessageReply                  bool   `json:"message_reply" to:"int"`
	MessageAllow                  bool   `json:"message_allow" to:"int"`
	MessageDeny                   bool   `json:"message_deny" to:"int"`
	MessageEdit                   bool   `json:"message_edit" to:"int"`
	MessageTypingState            bool   `json:"message_typing_state" to:"int"`
	PhotoNew                      bool   `json:"photo_new" to:"int"`
	AudioNew                      bool   `json:"audio_new" to:"int"`
	VideoNew                      bool   `json:"video_new" to:"int"`
	WallReplyNew                  bool   `json:"wall_reply_new" to:"int"`
	WallReplyEdit                 bool   `json:"wall_reply_edit" to:"int"`
	WallReplyDelete               bool   `json:"wall_reply_delete" to:"int"`
	WallReplyRestore              bool   `json:"wall_reply_restore" to:"int"`
	WallPostNew                   bool   `json:"wall_post_new" to:"int"`
	WallRepost                    bool   `json:"wall_repost" to:"int"`
	BoardPostNew                  bool   `json:"board_post_new" to:"int"`
	BoardPostEdit                 bool   `json:"board_post_edit" to:"int"`
	BoardPostRestore              bool   `json:"board_post_restore" to:"int"`
	BoardPostDelete               bool   `json:"board_post_delete" to:"int"`
	PhotoCommentNew               bool   `json:"photo_comment_new" to:"int"`
	PhotoCommentEdit              bool   `json:"photo_comment_edit" to:"int"`
	PhotoCommentDelete            bool   `json:"photo_comment_delete" to:"int"`
	PhotoCommentRestore           bool   `json:"photo_comment_restore" to:"int"`
	VideoCommentNew               bool   `json:"video_comment_new" to:"int"`
	VideoCommentEdit              bool   `json:"video_comment_edit" to:"int"`
	VideoCommentDelete            bool   `json:"video_comment_delete" to:"int"`
	VideoCommentRestore           bool   `json:"video_comment_restore" to:"int"`
	MarketCommentNew              bool   `json:"market_comment_new" to:"int"`
	MarketCommentEdit             bool   `json:"market_comment_edit" to:"int"`
	MarketCommentDelete           bool   `json:"market_comment_delete" to:"int"`
	MarketCommentRestore          bool   `json:"market_comment_restore" to:"int"`
	PollVoteNew                   bool   `json:"poll_vote_new" to:"int"`
	GroupJoin                     bool   `json:"group_join" to:"int"`
	GroupLeave                    bool   `json:"group_leave" to:"int"`
	GroupChangeSettings           bool   `json:"group_change_settings" to:"int"`
	GroupChangePhoto              bool   `json:"group_change_photo" to:"int"`
	GroupOfficersEdit             bool   `json:"group_officers_edit" to:"int"`
	UserBlock                     bool   `json:"user_block" to:"int"`
	UserUnblock                   bool   `json:"user_unblock" to:"int"`
	LikeAdd                       bool   `json:"like_add" to:"int"`
	LikeRemove                    bool   `json:"like_remove" to:"int"`
	MessageEvent                  bool   `json:"message_event" to:"int"`
	DonutSubscriptionCreate       bool   `json:"donut_subscription_create" to:"int"`
	DonutSubscriptionProlonged    bool   `json:"donut_subscription_prolonged" to:"int"`
	DonutSubscriptionCancelled    bool   `json:"donut_subscription_cancelled" to:"int"`
	DonutSubscriptionPriceChanged bool   `json:"donut_subscription_price_changed" to:"int"`
	DonutSubscriptionExpired      bool   `json:"donut_subscription_expired" to:"int"`
	DonutMoneyWithdraw            bool   `json:"donut_money_withdraw" to:"int"`
	DonutMoneyWithdrawError       bool   `json:"donut_money_withdraw_error" to:"int"`
}

// UnbanUser

type UnbanUser struct {
	GroupID int `json:"group_id"`
	OwnerID int `json:"owner_id"`
}
