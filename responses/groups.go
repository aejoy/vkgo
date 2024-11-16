package responses

import "github.com/aejoy/vkgo/update"

// AddGroupAddress

type AddGroupAddressReply struct {
	ErrorInterface
	Response AddGroupAddress `json:"response,omitempty"`
}

type AddGroupAddress struct {
	ID             int        `json:"id,omitempty"`
	Address        string     `json:"address,omitempty"`
	CityID         int        `json:"city_id,omitempty"`
	CountryID      int        `json:"country_id,omitempty"`
	City           City       `json:"city,omitempty"`
	Country        Country    `json:"country,omitempty"`
	Latitude       float64    `json:"latitude,omitempty"`
	Longitude      float64    `json:"longitude,omitempty"`
	OpenStatus     OpenStatus `json:"open_status,omitempty"`
	Title          string     `json:"title,omitempty"`
	WorkInfoStatus string     `json:"work_info_status,omitempty"`
}

type OpenStatus struct {
	Style         string `json:"style,omitempty"`
	PrimaryText   string `json:"primary_text,omitempty"`
	SecondaryText string `json:"secondary_text,omitempty"`
}

// AddGroupCallbackServer

type AddGroupCallbackServerReply struct {
	ErrorInterface
	Response AddGroupCallbackServer `json:"response,omitempty"`
}

type AddGroupCallbackServer struct {
	ServerID int `json:"server_id,omitempty"`
}

// AddGroupLink

type AddGroupLinkReply struct {
	Response AddGroupLink `json:"response,omitempty"`
}

type AddGroupLink struct {
	ID        int    `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	Name      string `json:"name,omitempty"`
	EditTitle int    `json:"edit_title,omitempty"`
	Desc      string `json:"desc,omitempty"`
	Photo100  string `json:"photo_100,omitempty"`
	Photo50   string `json:"photo_50,omitempty"`
}

// ApproveRequestGroup
// CreateGroup

type CreateGroup struct {
	ErrorInterface
	Response Group `json:"response,omitempty"`
}

// EditGroupManager

type EditGroupManagerReply struct {
	ErrorInterface
	Response EditGroupManager `json:"response,omitempty"`
}

type EditGroupManager struct {
	Status string `json:"status"`
}

// GetUserGroups

type UserGroupsReply struct {
	ErrorInterface
	Response UserGroups `json:"response,omitempty"`
}

type UserGroups struct {
	Groups []Group `json:"items,omitempty"`
}

// GetGroup

type GroupsReply struct {
	ErrorInterface
	Response Groups `json:"response,omitempty"`
}

type Groups struct {
	Groups []Group `json:"groups,omitempty"`
}

type Group struct {
	ErrorInterface
	ID                int       `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Domain            string    `json:"screen_name,omitempty"`
	IsClosed          int       `json:"is_closed,omitempty"`
	Deactivated       string    `json:"deactivated,omitempty"`
	IsAdmin           int       `json:"is_admin,omitempty"`
	AdminLevel        int       `json:"admin_level,omitempty"`
	IsMember          int       `json:"is_member,omitempty"`
	IsAdvertiser      int       `json:"is_advertiser,omitempty"`
	InvitedBy         int       `json:"invited_by,omitempty"`
	Type              string    `json:"type,omitempty"`
	Photo50           string    `json:"photo_50,omitempty"`
	Photo100          string    `json:"photo_100,omitempty"`
	Photo200          string    `json:"photo_200,omitempty"`
	Activity          string    `json:"activity,omitempty"`
	Addresses         Addresses `json:"addresses,omitempty"`
	AgeLimits         int       `json:"age_limits,omitempty"`
	BanInfo           BanInfo   `json:"ban_info,omitempty"`
	CanCreateTopic    int       `json:"can_create_topic,omitempty"`
	CanMessage        int       `json:"can_message,omitempty"`
	CanPost           int       `json:"can_post,omitempty"`
	CanSuggest        int       `json:"can_suggest,omitempty"`
	CanSeeAllPosts    int       `json:"can_see_all_posts,omitempty"`
	CanUploadDoc      int       `json:"can_upload_doc,omitempty"`
	CanUploadStory    int       `json:"can_upload_story,omitempty"`
	CanUploadVideo    int       `json:"can_upload_video,omitempty"`
	City              City      `json:"city,omitempty"`
	Contacts          []Contact `json:"contacts,omitempty"`
	Counters          Counters  `json:"counters,omitempty"`
	Country           Country   `json:"country,omitempty"`
	Cover             Cover     `json:"cover,omitempty"`
	CropPhoto         CropPhoto `json:"crop_photo,omitempty"`
	Description       string    `json:"description,omitempty"`
	PinnedPost        int       `json:"fixed_post,omitempty"`
	HasPhoto          int       `json:"has_photo,omitempty"`
	IsFavorite        int       `json:"is_favorite,omitempty"`
	IsHiddenFromFeed  int       `json:"is_hidden_from_feed,omitempty"`
	IsMessagesBlocked int       `json:"is_messages_blocked,omitempty"`
	Links             []Link    `json:"links,omitempty"`
	MainAlbumID       int       `json:"main_album_id,omitempty"`
	MainSection       int       `json:"main_section,omitempty"`
	Market            Market    `json:"market,omitempty"`
	MemberStatus      int       `json:"member_status,omitempty"`
	MembersCount      int       `json:"members_count,omitempty"`
	Place             Place     `json:"place,omitempty"`
	PublicDateLabel   string    `json:"public_date_label,omitempty"`
	Site              string    `json:"site,omitempty"`
	StartDate         any       `json:"start_date,omitempty"`
	FinishDate        string    `json:"finish_date,omitempty"`
	Status            string    `json:"status,omitempty"`
	Trending          int       `json:"trending,omitempty"`
	Verified          int       `json:"verified,omitempty"`
	Wall              int       `json:"wall,omitempty"`
	WikiPage          string    `json:"wiki_page,omitempty"`

	// without documentation
	// Like Like `json:"like,omitempty"`
	DeactivatedMessage string `json:"deactivated_message,omitempty"`
	DeactivatedType    string `json:"deactivated_type,omitempty"`
}

type Addresses struct {
	IsEnabled     bool        `json:"is_enabled,omitempty"`
	MainAddressID int         `json:"main_address_id,omitempty"`
	MainAddress   MainAddress `json:"main_address"`
}

type MainAddress struct {
	ID             int     `json:"id,omitempty"`
	Address        string  `json:"address,omitempty"`
	City           City    `json:"city,omitempty"`
	Country        Country `json:"country,omitempty"`
	Title          string  `json:"title,omitempty"`
	WorkInfoStatus string  `json:"work_info_status,omitempty"`
}

type Like struct {
	Liked   bool        `json:"is_liked,omitempty"`
	Friends FriendsLike `json:"friends,omitempty"`
}

type FriendsLike struct {
	Count   int `json:"count,omitempty"`
	Preview any `json:"preview,omitempty"`
}

type BanInfo struct {
	EndDate int    `json:"end_date,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type Country struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type City struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type Counter struct {
	Albums          int `json:"albums,omitempty"`
	Market          int `json:"market,omitempty"`
	Photos          int `json:"photos,omitempty"`
	Topics          int `json:"topics,omitempty"`
	Articles        int `json:"articles,omitempty"`
	Narratives      int `json:"narratives,omitempty"`
	ClassifiedYoula int `json:"classified_youla,omitempty"`
}

type Cover struct {
	Enabled       int                `json:"enabled,omitempty"`
	Images        []update.PhotoSize `json:"images,omitempty"`
	OriginalImage update.PhotoSize   `json:"original_image,omitempty"`
	PhotoID       int                `json:"photo_id"`
}

type CropPhoto struct {
	Photo update.Photo `json:"photo,omitempty"`
	Crop  Crop         `json:"crop,omitempty"`
	Rect  Crop         `json:"rect,omitempty"`
}

type Crop struct {
	X  any `json:"x,omitempty"`
	Y  any `json:"y,omitempty"`
	X2 any `json:"x2,omitempty"`
	Y2 any `json:"y2,omitempty"`
}

type Link struct {
	ID          int    `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
	Photo50     string `json:"photo_50,omitempty"`
	Photo100    string `json:"photo_100,omitempty"`
}

type Market struct {
	Enabled               int            `json:"enabled,omitempty"`
	Buttons               Buttons        `json:"buttons"`
	ShopConditions        StopConditions `json:"shop_conditions"`
	ViewedProductsEnabled bool           `json:"viewed_products_enabled"`

	IsHsEnabled bool     `json:"is_hs_enabled"`
	Type        string   `json:"type,omitempty"`
	PriceMax    string   `json:"price_max,omitempty"`
	PriceMin    string   `json:"price_min,omitempty"`
	MainAlbumID int      `json:"main_album_id,omitempty"`
	ContactID   int      `json:"contact_id,omitempty"`
	Currency    Currency `json:"currency,omitempty"`
	Symbol      string   `json:"currency_text,omitempty"`

	// without documentation
	IsShowHeaderItemsLink    int  `json:"is_show_header_items_link,omitempty"`
	IsUseSimplifiedShowcase  bool `json:"is_use_simplified_showcase,omitempty"`
	HasNotInMarketTab        bool `json:"has_not_in_market_tab,omitempty"`
	HasModerationRejectedTab bool `json:"has_moderation_rejected_tab,omitempty"`
	UnviewedOrdersCount      int  `json:"unviewed_orders_count,omitempty"`
	// CommunityManageEnabled   int  `json:"is_community_manage_enabled,omitempty"`
}

type Buttons struct {
	ButtonTypeWritePreset ButtonTypeWritePreset `json:"button_type_write_preset,omitempty"`
	ButtonTypeOpenPreset  ButtonTypeOpenPreset  `json:"button_type_open_preset,omitempty"`
	ButtonType            []ButtonType          `json:"button_type,omitempty"`
	ButtonTypeCallPreset  ButtonTypeCallPreset  `json:"button_type_call_preset,omitempty"`
	Buttons               []interface{}         `json:"buttons,omitempty"`
}

type ButtonTypeWritePreset struct {
	Recipients   []ButtonRecipient `json:"recipients,omitempty"`
	ButtonTitles []ButtonTitle     `json:"button_titles,omitempty"`
}

type ButtonTypeOpenPreset struct {
	ButtonTitles []ButtonTitle `json:"button_titles,omitempty"`
}

type ButtonType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ButtonTypeCallPreset struct {
	ValidationRegex         string     `json:"validation_regex,omitempty"`
	ValidationRegexErrorMsg string     `json:"validation_regex_error_msg,omitempty"`
	Timezones               []Timezone `json:"timezones,omitempty"`
	Weekdays                []Weekday  `json:"weekdays,omitempty"`
}

type Timezone struct {
	Timezone string `json:"timezone,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Weekday struct {
	Weekday string `json:"weekday,omitempty"`
	Name    string `json:"name,omitempty"`
}

type ButtonRecipient struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ButtonTitle struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	IsDisabled bool   `json:"is_disabled,omitempty"`
}

type StopConditions struct {
	Fields          StopConditionsFields    `json:"fields,omitempty"`
	Header          string                  `json:"header,omitempty"`
	Buttons         []StopConditionsButtons `json:"buttons,omitempty"`
	IntegrationType string                  `json:"integration_type,omitempty"`
}

type StopConditionsFields struct {
	Delivery FieldsItem `json:"delivery,omitempty"`
	Payment  FieldsItem `json:"payment,omitempty"`
	Refund   FieldsItem `json:"refund,omitempty"`
}

type StopConditionsButtons struct {
	Action StopConditionsAction `json:"action,omitempty"`
	Title  string               `json:"title,omitempty"`
}

type StopConditionsAction struct {
	Type                 string               `json:"type,omitempty"`
	PerformActionWithURL PerformActionWithURL `json:"perform_action_with_url,omitempty"`
	URL                  string               `json:"url,omitempty"`
}

type PerformActionWithURL struct {
	Action string `json:"action,omitempty"`
}

type FieldsItem struct {
	Title       string `json:"title,omitempty"`
	Text        string `json:"text,omitempty"`
	OnEmptyText string `json:"on_empty_text,omitempty"`
}

type Currency struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Symbol string `json:"title,omitempty"`
}

type Place struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Latitude  any    `json:"latitude,omitempty"`
	Longitude any    `json:"longitude,omitempty"`
	Type      string `json:"type,omitempty"`
	CountryID int    `json:"country,omitempty"`
	CityID    int    `json:"city,omitempty"`
	Address   string `json:"address,omitempty"`
}

// GetBannedUsers

type BannedUsersReply struct {
	ErrorInterface
	Response BannedUsers `json:"response,omitempty"`
}

type BannedUsers struct {
	Count int                  `json:"count,omitempty"`
	Bans  []BannedUserResponse `json:"items,omitempty"`
}

type BannedUserResponse struct {
	ErrorInterface
	Type    string      `json:"type,omitempty"`
	Info    InfoBanUser `json:"ban_info,omitempty"`
	Profile User        `json:"profile,omitempty"`
}

type InfoBanUser struct {
	AdminID     int    `json:"admin_id,omitempty"`
	Comment     string `json:"comment,omitempty"`
	ShowComment bool   `json:"comment_visible,omitempty"`
	Date        int    `json:"date,omitempty"`
	EndDate     int    `json:"end_date,omitempty"`
	Reason      int    `json:"reason,omitempty"`
}

// GetGroupCallbackConfirmationCode

type GroupCallbackConfirmationCodeReply struct {
	ErrorInterface
	Response GroupCallbackConfirmationCode `json:"response,omitempty"`
}

type GroupCallbackConfirmationCode struct {
	Code string `json:"code,omitempty"`
}

// GetGroupCallbackServers

type GroupCallbackServersReply struct {
	ErrorInterface
	Response GroupCallbackServers `json:"response,omitempty"`
}

type GroupCallbackServers struct {
	Count int                   `json:"count,omitempty"`
	Items []GroupCallbackServer `json:"items,omitempty"`
}

type GroupCallbackServer struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatorID int    `json:"creator_id,omitempty"`
	URL       string `json:"url,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
	Status    string `json:"status,omitempty"`
}

// GetGroupCallbackSettings

type GroupCallbackSettingsReply struct {
	ErrorInterface
	Response GroupCallbackSettings `json:"response,omitempty"`
}

type GroupCallbackSettings struct {
	APIVersion string         `json:"api_version,omitempty"`
	Events     SettingsEvents `json:"events,omitempty"`
}

type SettingsEvents struct {
	AppPayload                    int `json:"app_payload,omitempty"`
	AudioNew                      int `json:"audio_new,omitempty"`
	BoardPostDelete               int `json:"board_post_delete,omitempty"`
	BoardPostEdit                 int `json:"board_post_edit,omitempty"`
	BoardPostNew                  int `json:"board_post_new,omitempty"`
	BoardPostRestore              int `json:"board_post_restore,omitempty"`
	GroupChangePhoto              int `json:"group_change_photo,omitempty"`
	GroupChangeSettings           int `json:"group_change_settings,omitempty"`
	GroupJoin                     int `json:"group_join,omitempty"`
	GroupLeave                    int `json:"group_leave,omitempty"`
	GroupOfficersEdit             int `json:"group_officers_edit,omitempty"`
	MarketCommentDelete           int `json:"market_comment_delete,omitempty"`
	MarketCommentEdit             int `json:"market_comment_edit,omitempty"`
	MarketCommentNew              int `json:"market_comment_new,omitempty"`
	MarketCommentRestore          int `json:"market_comment_restore,omitempty"`
	MessageAllow                  int `json:"message_allow,omitempty"`
	MessageDeny                   int `json:"message_deny,omitempty"`
	MessageNew                    int `json:"message_new,omitempty"`
	MessageReply                  int `json:"message_reply,omitempty"`
	MessageEdit                   int `json:"message_edit,omitempty"`
	PhotoCommentDelete            int `json:"photo_comment_delete,omitempty"`
	PhotoCommentEdit              int `json:"photo_comment_edit,omitempty"`
	PhotoCommentNew               int `json:"photo_comment_new,omitempty"`
	PhotoCommentRestore           int `json:"photo_comment_restore,omitempty"`
	PhotoNew                      int `json:"photo_new,omitempty"`
	PollVoteNew                   int `json:"poll_vote_new,omitempty"`
	UserBlock                     int `json:"user_block,omitempty"`
	UserUnblock                   int `json:"user_unblock,omitempty"`
	VideoCommentDelete            int `json:"video_comment_delete,omitempty"`
	VideoCommentEdit              int `json:"video_comment_edit,omitempty"`
	VideoCommentNew               int `json:"video_comment_new,omitempty"`
	VideoCommentRestore           int `json:"video_comment_restore,omitempty"`
	VideoNew                      int `json:"video_new,omitempty"`
	MessageReactionEvent          int `json:"message_reaction_event,omitempty"`
	WallPostNew                   int `json:"wall_post_new,omitempty"`
	WallReplyDelete               int `json:"wall_reply_delete,omitempty"`
	WallReplyEdit                 int `json:"wall_reply_edit,omitempty"`
	WallReplyNew                  int `json:"wall_reply_new,omitempty"`
	WallReplyRestore              int `json:"wall_reply_restore,omitempty"`
	WallRepost                    int `json:"wall_repost,omitempty"`
	WallSchedulePostNew           int `json:"wall_schedule_post_new,omitempty"`
	WallSchedulePostDelete        int `json:"wall_schedule_post_delete,omitempty"`
	LeadFormsNew                  int `json:"lead_forms_new,omitempty"`
	LikeAdd                       int `json:"like_add,omitempty"`
	LikeRemove                    int `json:"like_remove,omitempty"`
	MarketOrderNew                int `json:"market_order_new,omitempty"`
	MarketOrderEdit               int `json:"market_order_edit,omitempty"`
	MessageRead                   int `json:"message_read,omitempty"`
	MessageTypingState            int `json:"message_typing_state,omitempty"`
	VkpayTransaction              int `json:"vkpay_transaction,omitempty"`
	MessageEvent                  int `json:"message_event,omitempty"`
	DonutSubscriptionCreate       int `json:"donut_subscription_create,omitempty"`
	DonutSubscriptionProlonged    int `json:"donut_subscription_prolonged,omitempty"`
	DonutSubscriptionCancelled    int `json:"donut_subscription_cancelled,omitempty"`
	DonutSubscriptionExpired      int `json:"donut_subscription_expired,omitempty"`
	DonutSubscriptionPriceChanged int `json:"donut_subscription_price_changed,omitempty"`
	DonutMoneyWithdraw            int `json:"donut_money_withdraw,omitempty"`
	DonutMoneyWithdrawError       int `json:"donut_money_withdraw_error,omitempty"`
	MusicSubscriptionUpdate       int `json:"music_subscription_update,omitempty"`
	InappSubscriptionUpdate       int `json:"inapp_subscription_update,omitempty"`
	InappOrderCreate              int `json:"inapp_order_create,omitempty"`
	InappOrderRefund              int `json:"inapp_order_refund,omitempty"`
}

type GroupCatalogInfoReply struct {
	ErrorInterface
	Response GroupCatalogInfo `json:"response,omitempty"`
}

type GroupCatalogInfo struct {
	Enabled       int                        `json:"enabled,omitempty"`
	Categories    []GroupCatalogInfoCategory `json:"categories,omitempty"`
	Subcategories []Subcategory              `json:"subcategories,omitempty"`
}

type GroupCatalogInfoCategory struct {
	ID           int                           `json:"id,omitempty"`
	Name         string                        `json:"name,omitempty"`
	PageCount    int                           `json:"page_count,omitempty"`
	PagePreviews []GroupCatalogInfoPagePreview `json:"page_previews,omitempty"`
}

type GroupCatalogInfoPagePreview struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ScreenName   string `json:"screen_name,omitempty"`
	IsClosed     int    `json:"is_closed,omitempty"`
	Type         string `json:"type,omitempty"`
	IsAdmin      int    `json:"is_admin,omitempty"`
	IsMember     int    `json:"is_member,omitempty"`
	IsAdvertiser int    `json:"is_advertiser,omitempty"`
	Photo50      string `json:"photo_50,omitempty"`
	Photo100     string `json:"photo_100,omitempty"`
	Photo200     string `json:"photo_200,omitempty"`
}

type Subcategory struct {
	ID      int      `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Genders []Gender `json:"genders,omitempty"`
}

type Gender struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetGroupInvitedUsers

type GroupInvitedUsers struct {
	Response GroupInvitedUser `json:"response,omitempty"`
}

type GroupInvitedUser struct {
	Items []User `json:"items,omitempty"`
}

// GetGroupInvites

type GroupInvitesReply struct {
	ErrorInterface
	Response GroupInvites `json:"response,omitempty"`
}

type GroupInvites struct {
	Count  int     `json:"count,omitempty"`
	Items  []Group `json:"items,omitempty"`
	Users  []User  `json:"profiles,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}

// GetGroupLongPollServer

type LongPollServerReply struct {
	ErrorInterface
	Response LongPollServer `json:"response,omitempty"`
}

type LongPollServer struct {
	Server  string `json:"server,omitempty"`
	Key     string `json:"key,omitempty"`
	TS      any    `json:"ts,omitempty"`
	PTS     any    `json:"pts,omitempty"`
	Version int    `json:"version,omitempty"`
}

// GetGroupLongPollSettings

type LongPollSettingsReply struct {
	ErrorInterface
	Response LongPollSettings `json:"response,omitempty"`
}

type LongPollSettings struct {
	Events     SettingsEvents `json:"events,omitempty"`
	IsEnabled  bool           `json:"is_enabled,omitempty"`
	APIVersion string         `json:"api_version,omitempty"`
}

// GetGroupMembers

type GroupMembersReply struct {
	ErrorInterface
	Response GroupMembers `json:"response,omitempty"`
}

type GroupMembers struct {
	Count    int           `json:"count,omitempty"`
	Items    []GroupMember `json:"items,omitempty"`
	NextFrom string        `json:"next_from,omitempty"`
}

type GroupMember struct {
	ID              int    `json:"id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	CanAccessClosed bool   `json:"can_access_closed,omitempty"`
	IsClosed        bool   `json:"is_closed,omitempty"`
	Deactivated     string `json:"deactivated,omitempty"`
}

// GetGroupOnlineStatus

type GroupOnlineStatusReply struct {
	ErrorInterface
	Response GroupOnlineStatus `json:"response,omitempty"`
}

type GroupOnlineStatus struct {
	Status string `json:"status,omitempty"`
}

// GetGroupRequests

type GroupRequestsReply struct {
	ErrorInterface
	Response GroupRequests `json:"response,omitempty"`
}

type GroupRequests struct {
	Items []User `json:"items,omitempty"`
}

// GetGroupSettings

type GroupSettingsReply struct {
	ErrorInterface
	Response GroupSettings `json:"response,omitempty"`
}

type GroupSettings struct {
	Audio                    int              `json:"audio,omitempty"`
	Articles                 int              `json:"articles,omitempty"`
	CityID                   int              `json:"city_id,omitempty"`
	CityName                 string           `json:"city_name,omitempty"`
	Description              string           `json:"description,omitempty"`
	Docs                     int              `json:"docs,omitempty"`
	Messages                 int              `json:"messages,omitempty"`
	ObsceneFilter            int              `json:"obscene_filter,omitempty"`
	ObsceneStopwords         int              `json:"obscene_stopwords,omitempty"`
	ToxicFilter              int              `json:"toxic_filter,omitempty"`
	DisableRepliesFromGroups int              `json:"disable_replies_from_groups,omitempty"`
	Photos                   int              `json:"photos,omitempty"`
	Title                    string           `json:"title,omitempty"`
	Topics                   int              `json:"topics,omitempty"`
	Video                    int              `json:"video,omitempty"`
	Wall                     int              `json:"wall,omitempty"`
	Wiki                     int              `json:"wiki,omitempty"`
	Access                   int              `json:"access,omitempty"`
	ActionButton             ActionButton     `json:"action_button,omitempty"`
	Address                  string           `json:"address,omitempty"`
	RecognizePhoto           int              `json:"recognize_photo,omitempty"`
	Contacts                 int              `json:"contacts,omitempty"`
	Links                    int              `json:"links,omitempty"`
	LiveCovers               LiveCovers       `json:"live_covers,omitempty"`
	SectionsList             []Section        `json:"sections_list,omitempty"`
	MainSection              int              `json:"main_section,omitempty"`
	SecondarySection         int              `json:"secondary_section,omitempty"`
	AgeLimits                int              `json:"age_limits,omitempty"`
	CountryID                int              `json:"country_id,omitempty"`
	CountryName              string           `json:"country_name,omitempty"`
	Events                   int              `json:"events,omitempty"`
	Market                   Market           `json:"market,omitempty"`
	MarketServices           MarketServices   `json:"market_services,omitempty"`
	Narratives               int              `json:"narratives,omitempty"`
	BotsCapabilities         int              `json:"bots_capabilities,omitempty"`
	BotsStartButton          int              `json:"bots_start_button,omitempty"`
	BotsAddToChat            int              `json:"bots_add_to_chat,omitempty"`
	Clips                    int              `json:"clips,omitempty"`
	Textlives                int              `json:"textlives,omitempty"`
	ObsceneWords             []any            `json:"obscene_words,omitempty"`
	PublicCategory           int              `json:"public_category,omitempty"`
	PublicCategoryList       []PublicCategory `json:"public_category_list,omitempty"`
	PublicSubcategory        int              `json:"public_subcategory,omitempty"`
	SuggestedPrivacy         int              `json:"suggested_privacy,omitempty"`
	Website                  string           `json:"website,omitempty"`
	Phone                    string           `json:"phone,omitempty"`
	EnableReplies            int              `json:"enable_replies,omitempty"`
	IsSuggestedPostEnabled   bool             `json:"is_suggested_post_enabled,omitempty"`
}

type ActionButton struct {
	IsEnabled bool `json:"is_enabled,omitempty"`
}

type LiveCovers struct {
	IsEnabled bool `json:"is_enabled,omitempty"`
}

type Section struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type MarketServices struct {
	Enabled         bool        `json:"enabled,omitempty"`
	CanMessage      bool        `json:"can_message,omitempty"`
	CommentsEnabled bool        `json:"comments_enabled,omitempty"`
	ContactID       int         `json:"contact_id,omitempty"`
	Currency        Currency    `json:"currency,omitempty"`
	ViewType        ViewType    `json:"view_type,omitempty"`
	BlockName       BlockName   `json:"block_name,omitempty"`
	ButtonLabel     ButtonLabel `json:"button_label,omitempty"`
}

type BlockName struct {
	SelectedItemID int    `json:"selected_item_id,omitempty"`
	Items          []Item `json:"items,omitempty"`
}

type ButtonLabel struct {
	SelectedItemID int               `json:"selected_item_id,omitempty"`
	Items          []ButtonLabelItem `json:"items,omitempty"`
}

type ButtonLabelItem struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	IsDisabled bool   `json:"is_disabled,omitempty"`
}

type ViewType struct {
	SelectedItemID int    `json:"selected_item_id,omitempty"`
	Items          []Item `json:"items,omitempty"`
}

type PublicCategory struct {
	ID            int           `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	Subcategories []Subcategory `json:"subcategories,omitempty"`
}

// GetGroupTokenPermissions

type GroupTokenPermissionsReply struct {
	ErrorInterface
	Response GroupTokenPermissions `json:"response,omitempty"`
}

type GroupTokenPermissions struct {
	Mask        int                    `json:"mask,omitempty"`
	Permissions []GroupTokenPermission `json:"permissions,omitempty"`
}

type GroupTokenPermission struct {
	Name    string `json:"name,omitempty"`
	Setting int    `json:"setting,omitempty"`
}

// InviteGroup
// IsGroupMember

type IsGroupMemberReply struct {
	ErrorInterface
	Response IsGroupMember `json:"response,omitempty"`
}

type IsGroupMember struct {
	Member    int `json:"member,omitempty"`
	CanInvite int `json:"can_invite,omitempty"`
}
