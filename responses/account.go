package responses

// ChangePassword

type ChangePasswordReply struct {
	ErrorInterface
	Response ChangePassword `json:"response,omitempty"`
}

type ChangePassword struct {
	Token string `json:"token,omitempty"`
}

// GetActiveOffers

type ActiveOffersReply struct {
	ErrorInterface
	Response ActiveOffers `json:"response,omitempty"`
}

type ActiveOffers struct {
	Offers []int `json:"items,omitempty"`
}

// GetAppPermissions

type AppPermissionsReply struct {
	ErrorInterface
	Response int `json:"response,omitempty"`
}

// GetBans

type BansReply struct {
	ErrorInterface
	Response Bans `json:"response"`
}

type Bans struct {
	Bans  []int  `json:"items,omitempty"`
	Users []User `json:"profiles,omitempty"`
}

// GetCounters

type CountersReply struct {
	ErrorInterface
	Response Counters `json:"response,omitempty"`
}

type Counters struct {
	Photos int `json:"photos,omitempty"`
	Albums int `json:"albums,omitempty"`
	Audios int `json:"audios,omitempty"`
	Videos int `json:"videos,omitempty"`
	Topics int `json:"topics,omitempty"`
	Docs   int `json:"docs,omitempty"`

	AppRequests          int              `json:"app_requests,omitempty"`
	Faves                int              `json:"faves,omitempty"`
	DiscoverBadge        int              `json:"menu_discover_badge,omitempty"`
	ClipsBadge           int              `json:"menu_clips_badge,omitempty"`
	SuperAppFriendsBadge int              `json:"menu_superapp_friends_badge,omitempty"`
	NewClipsBadge        int              `json:"menu_new_clips_badge,omitempty"`
	Messages             int              `json:"messages,omitempty"`
	SDK                  int              `json:"sdk,omitempty"`
	Calls                int              `json:"calls,omitempty"`
	MessagesFolders      []MessagesFolder `json:"messages_folders,omitempty"`

	// without documentation
	Addresses       int `json:"addresses,omitempty"`
	VideoPlaylists  int `json:"video_playlists,omitempty"`
	Narratives      int `json:"narratives,omitempty"`
	Clips           int `json:"clips,omitempty"`
	ClipsFollowers  int `json:"clips_followers,omitempty"`
	ClassifiedYoula int `json:"classified_youla,omitempty"`
}

type MessagesFolder struct {
	ID      int `json:"folder_id,omitempty"`
	Count   int `json:"total_count,omitempty"`
	UnMuted int `json:"unmuted_count,omitempty"`
}

// GetInfo

type InfoReply struct {
	ErrorInterface
	Response Info `json:"response,omitempty"`
}

type Info struct {
	Country       string `json:"country,omitempty"`
	Secure        int    `json:"https_required,omitempty"`
	TwoFactor     int    `json:"2fa_required,omitempty"`
	OwnerPosts    int    `json:"own_posts_default,omitempty"`
	NoWallReplies int    `json:"no_wall_replies,omitempty"`
	Intro         int    `json:"intro,omitempty"`
	Language      int    `json:"lang,omitempty"`

	// without documentation
	CommunityComments bool              `json:"community_comments,omitempty"`
	LinkRedirects     map[string]string `json:"link_redirects,omitempty"`
	VKPayEndpoint     string            `json:"vk_pay_endpoint_v2,omitempty"`
	Translation       []string          `json:"messages_translation_language_pairs,omitempty"`
	ObsceneFilter     bool              `json:"obscene_text_filter,omitempty"`
}

// GetProfileInfo

type ProfileInfoReply struct {
	ErrorInterface
	Response ProfileInfo `json:"response,omitempty"`
}

type ProfileInfo struct {
	Name                string      `json:"first_name,omitempty"`
	Surame              string      `json:"last_name,omitempty"`
	MaidenName          string      `json:"maiden_name,omitempty"`
	Domain              string      `json:"screen_name,omitempty"`
	Sex                 int         `json:"sex,omitempty"`
	Relation            int         `json:"relation,omitempty"` // 📡👽
	RelationPartner     User        `json:"relation_partner,omitempty"`
	RelationPending     int         `json:"relation_pending,omitempty"`
	RelationRequests    []User      `json:"relation_requests,omitempty"`
	Birthdate           string      `json:"bdate,omitempty"`
	BirthdateVisibility int         `json:"bdate_visibility,omitempty"`
	HomeTown            string      `json:"home_town,omitempty"`
	Country             Country     `json:"country,omitempty"`
	City                City        `json:"city,omitempty"`
	NameRequest         NameRequest `json:"name_request,omitempty"`
	Status              string      `json:"status"`
	Phone               string      `json:"phone"`

	// without documentation
	ID                         int                        `json:"id,omitempty"`
	Photo200                   string                     `json:"photo_200"`
	IsServiceAccount           bool                       `json:"is_service_account"`
	IsEsiaVerified             bool                       `json:"is_esia_verified"`
	IsEsiaLinked               bool                       `json:"is_esia_linked"`
	IsTinkoffLinked            bool                       `json:"is_tinkoff_linked"`
	IsTinkoffVerified          bool                       `json:"is_tinkoff_verified"`
	IsVerified                 bool                       `json:"is_verified"`
	IsSberVerified             bool                       `json:"is_sber_verified"`
	OauthLinked                []string                   `json:"oauth_linked"`
	OauthVerification          []any                      `json:"oauth_verification"`
	AccountVerificationProfile AccountVerificationProfile `json:"account_verification_profile"`
	VerificationStatus         string                     `json:"verification_status"`
	PromoVerifications         []any                      `json:"promo_verifications"`
}

type AccountVerificationProfile struct {
	Name       string `json:"first_name,omitempty"`
	Surname    string `json:"last_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	Gender     int    `json:"sex,omitempty"`
	Birthdate  string `json:"birthdate,omitempty"`
}

type PushSettings struct {
	Sound           int `json:"sound,omitempty"`
	NoSound         any `json:"no_sound,omitempty"`
	UntilDisabled   int `json:"disabled_until,omitempty"`
	ForeverDisabled any `json:"disabled_forever,omitempty"`
	// needs improvement
}

// EditProfileInfo

type EditProfileInfo struct {
	ErrorInterface
	Changed     int         `json:"changed,omitempty"`
	NameRequest NameRequest `json:"name_request,omitempty"`
}

type NameRequest struct {
	ID      int    `json:"id,omitempty"`
	Status  string `json:"status,omitempty"`
	Name    string `json:"first_name,omitempty"`
	Surname string `json:"last_name,omitempty"`
}
