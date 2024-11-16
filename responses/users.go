package responses

type UsersReply struct {
	ErrorInterface
	Response []User `json:"response,omitempty"`
}

type User struct {
	ID                   int              `json:"id,omitempty"`
	Name                 string           `json:"first_name,omitempty"`
	Surname              string           `json:"last_name,omitempty"`
	Deactivated          string           `json:"deactivated,omitempty"`
	IsClosed             bool             `json:"is_closed"`
	CanAccessClosed      bool             `json:"can_access_closed"`
	About                string           `json:"about,omitempty"`
	Activities           string           `json:"activities,omitempty"`
	Birthdate            string           `json:"bdate,omitempty"`
	Blacklisted          int              `json:"blacklisted,omitempty"`
	BlacklistedByMe      int              `json:"blacklisted_by_me,omitempty"`
	Books                string           `json:"books,omitempty"`
	CanPost              int              `json:"can_post,omitempty"`
	CanSeeAllPosts       int              `json:"can_see_all_posts,omitempty"`
	CanSeeAudio          int              `json:"can_see_audio,omitempty"`
	CanSendFriendRequest int              `json:"can_send_friend_request,omitempty"`
	CanWriteMessage      int              `json:"can_write_private_message,omitempty"`
	Career               []UserCareer     `json:"career,omitempty"`
	City                 City             `json:"city,omitempty"`
	MutualFriends        int              `json:"common_count,omitempty"`
	Connections          any              `json:"connections,omitempty"`
	Contacts             UserContacts     `json:"contacts,omitempty"`
	Counters             UserCounters     `json:"counters,omitempty"`
	Country              Country          `json:"country,omitempty"`
	CropPhoto            CropPhoto        `json:"crop_photo,omitempty"`
	Domain               string           `json:"domain,omitempty"`
	Education            UserEducation    `json:"education,omitempty"`
	Exports              any              `json:"exports,omitempty"`
	NameNominative       string           `json:"first_name_nom,omitempty"`
	NameGenitive         string           `json:"first_name_gen,omitempty"`
	NameDative           string           `json:"first_name_dat,omitempty"`
	NameAccusative       string           `json:"first_name_acc,omitempty"`
	NameInstrumental     string           `json:"first_name_ins,omitempty"`
	NameAblative         string           `json:"first_name_abl,omitempty"`
	FollowersCount       int              `json:"followers_count,omitempty"`
	Games                string           `json:"games,omitempty"`
	HasMobile            int              `json:"has_mobile,omitempty"`
	HasPhoto             int              `json:"has_photo,omitempty"`
	HomeTown             string           `json:"home_town,omitempty"`
	Interests            string           `json:"interests,omitempty"`
	IsFavorite           int              `json:"is_favorite,omitempty"`
	IsFriend             int              `json:"is_friend,omitempty"`
	IsHiddenFromFeed     int              `json:"is_hidden_from_feed,omitempty"`
	IsNoIndex            int              `json:"is_no_index,omitempty"`
	SurnameNominative    string           `json:"last_name_nom,omitempty"`
	SurnameGenitive      string           `json:"last_name_gen,omitempty"`
	SurnameDative        string           `json:"last_name_dat,omitempty"`
	SurnameAccusative    string           `json:"last_name_acc,omitempty"`
	SurnameInstrumental  string           `json:"last_name_ins,omitempty"`
	SurnameAblative      string           `json:"last_name_abl,omitempty"`
	LastSeen             UserLastSeen     `json:"last_seen,omitempty"`
	Lists                []int            `json:"lists,omitempty"`
	MaidenName           string           `json:"maiden_name,omitempty"`
	Military             UserMilitary     `json:"military,omitempty"`
	Movies               string           `json:"movies,omitempty"`
	Music                string           `json:"music,omitempty"`
	Nickname             string           `json:"nickname,omitempty"`
	Occupation           UserOccupation   `json:"occupation,omitempty"`
	Online               int              `json:"online,omitempty"`
	OnlineMobile         int              `json:"online_mobile,omitempty"`
	OnlineApplication    int              `json:"online_app,omitempty"`
	OnlineInfo           OnlineInfo       `json:"online_info,omitempty"`
	Personal             UserPersonal     `json:"personal,omitempty"`
	Photo50              string           `json:"photo_50,omitempty"`
	Photo100             string           `json:"photo_100,omitempty"`
	Photo200Original     string           `json:"photo_200_orig,omitempty"`
	Photo200             string           `json:"photo_200,omitempty"`
	Photo400Original     string           `json:"photo_400_orig,omitempty"`
	PhotoID              string           `json:"photo_id,omitempty"`
	PhotoMax             string           `json:"photo_max,omitempty"`
	PhotoMaxOriginal     string           `json:"photo_max_orig,omitempty"`
	Cover                Cover            `json:"cover"`
	Quotes               string           `json:"quotes,omitempty"`
	Relatives            []UserRelative   `json:"relatives,omitempty"`
	Relation             int              `json:"relation,omitempty"`
	Schools              []UserSchool     `json:"schools,omitempty"`
	ScreenName           string           `json:"screen_name,omitempty"`
	Gender               int              `json:"sex,omitempty"`
	Site                 string           `json:"site,omitempty"`
	Status               string           `json:"status,omitempty"`
	Timezone             int              `json:"timezone,omitempty"`
	Trending             int              `json:"trending,omitempty"`
	TV                   string           `json:"tv,omitempty"`
	Universities         []UserUniversity `json:"universities,omitempty"`
	Verified             int              `json:"verified,omitempty"`
	WallDefault          string           `json:"wall_default,omitempty"`

	// without documentation
	TrackCode    string `json:"track_code,omitempty"`
	WallComments int    `json:"wall_comments,omitempty"`
	MobilePhone  string `json:"mobile_phone,omitempty"`
	HomePhone    string `json:"home_phone,omitempty"`
	FriendStatus int    `json:"friend_status,omitempty"`
}

type UserCareer struct {
	GroupID   int    `json:"group_id,omitempty"`
	Company   string `json:"company,omitempty"`
	CountryID int    `json:"country_id,omitempty"`
	CityID    int    `json:"city_id,omitempty"`
	CityName  string `json:"city_name,omitempty"`
	From      int    `json:"from,omitempty"`
	Until     int    `json:"until,omitempty"`
	Position  string `json:"position,omitempty"`
}

type UserContacts struct {
	MobilePhone string `json:"mobile_phone,omitempty"`
	HomePhone   string `json:"home_phone,omitempty"`
}

type UserCounters struct {
	Albums        int `json:"albums,omitempty"`
	Videos        int `json:"videos,omitempty"`
	Audios        int `json:"audios,omitempty"`
	Photos        int `json:"photos,omitempty"`
	Notes         int `json:"notes,omitempty"`
	Friends       int `json:"friends,omitempty"`
	Gifts         int `json:"gifts,omitempty"`
	Groups        int `json:"groups,omitempty"`
	OnlineFriends int `json:"online_friends,omitempty"`
	MutualFriends int `json:"mutual_friends,omitempty"`
	UserVideos    int `json:"user_videos,omitempty"`
	UserPhotos    int `json:"user_photos,omitempty"`
	Followers     int `json:"followers,omitempty"`
	Pages         int `json:"pages,omitempty"`
	Subscriptions int `json:"subscriptions,omitempty"`
}

type UserEducation struct {
	University     int    `json:"university,omitempty"`
	UniversityName string `json:"university_name,omitempty"`
	Faculty        int    `json:"faculty,omitempty"`
	FacultyName    string `json:"faculty_name,omitempty"`
	Graduation     int    `json:"graduation,omitempty"`
}

type UserLastSeen struct {
	Platform int `json:"platform,omitempty"`
	Time     int `json:"time,omitempty"`
}

type UserMilitary struct {
	Unit      string `json:"unit,omitempty"`
	UnitID    int    `json:"unit_id,omitempty"`
	CountryID int    `json:"country_id,omitempty"`
	From      int    `json:"from,omitempty"`
	Until     int    `json:"until,omitempty"`
}

type UserOccupation struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type UserPersonal struct {
	Political  int      `json:"political,omitempty"`
	Languages  []string `json:"langs,omitempty"`
	Religion   string   `json:"religion,omitempty"`
	InspiredBy string   `json:"inspired_by,omitempty"`
	PeopleMain int      `json:"people_main,omitempty"`
	LifeMain   int      `json:"life_main,omitempty"`
	Smoking    int      `json:"smoking,omitempty"`
	Alcohol    int      `json:"alcohol,omitempty"`
}

type UserRelative struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type UserSchool struct {
	ID            string `json:"id,omitempty"`
	CountryID     int    `json:"country,omitempty"`
	CityID        int    `json:"city,omitempty"`
	Name          string `json:"name,omitempty"`
	YearFrom      int    `json:"year_from,omitempty"`
	YearTo        int    `json:"year_to,omitempty"`
	YearGraduated int    `json:"year_graduated,omitempty"`
	Class         string `json:"class,omitempty"`
	Speciality    string `json:"speciality,omitempty"`
	Type          int    `json:"type,omitempty"`
	TypeString    string `json:"type_str,omitempty"`
}

type UserUniversity struct {
	ID              int    `json:"id,omitempty"`
	CountryID       int    `json:"country,omitempty"`
	CityID          int    `json:"city,omitempty"`
	Name            string `json:"name,omitempty"`
	FacultyID       int    `json:"faculty,omitempty"`
	FacultyName     string `json:"faculty_name,omitempty"`
	ChairID         int    `json:"chair,omitempty"`
	ChairName       string `json:"chair_name,omitempty"`
	Graduation      int    `json:"graduation,omitempty"`
	EducationForm   string `json:"education_form,omitempty"`
	EducationStatus string `json:"education_status,omitempty"`
}

type OnlineInfo struct {
	Visible       bool `json:"visible,omitempty"`
	LastSeen      int  `json:"last_seen,omitempty"`
	IsOnline      bool `json:"is_online,omitempty"`
	ApplicationID int  `json:"app_id,omitempty"`
	IsMobile      bool `json:"is_mobile,omitempty"`
}

type FollowersReply struct {
	ErrorInterface
	Response Followers `json:"response,omitempty"`
}

type Followers struct {
	FriendsCount int    `json:"friends_count,omitempty"`
	Items        []User `json:"items,omitempty"`
}

type SubscriptionsReply struct {
	ErrorInterface
	Response Subscriptions `json:"response,omitempty"`
}

type Subscriptions struct {
	Items []Subscription `json:"items,omitempty"`
}

type Subscription struct {
	ErrorInterface
	ID                   int              `json:"id,omitempty"`
	TrackCode            string           `json:"track_code,omitempty"`
	Name                 string           `json:"first_name,omitempty"`
	Surname              string           `json:"last_name,omitempty"`
	Deactivated          string           `json:"deactivated,omitempty"`
	IsClosed             bool             `json:"is_closed"`
	CanAccessClosed      bool             `json:"can_access_closed"`
	About                string           `json:"about,omitempty"`
	Activities           string           `json:"activities,omitempty"`
	Birthdate            string           `json:"bdate,omitempty"`
	Blacklisted          int              `json:"blacklisted,omitempty"`
	BlacklistedByMe      int              `json:"blacklisted_by_me,omitempty"`
	Books                string           `json:"books,omitempty"`
	CanPost              int              `json:"can_post,omitempty"`
	CanSeeAllPosts       int              `json:"can_see_all_posts,omitempty"`
	CanSeeAudio          int              `json:"can_see_audio,omitempty"`
	CanSendFriendRequest int              `json:"can_send_friend_request,omitempty"`
	CanWriteMessage      int              `json:"can_write_private_message,omitempty"`
	Career               []UserCareer     `json:"career,omitempty"`
	City                 City             `json:"city,omitempty"`
	MutualFriends        int              `json:"common_count,omitempty"`
	Connections          any              `json:"connections,omitempty"`
	Contacts             UserContacts     `json:"contacts,omitempty"`
	Counters             UserCounters     `json:"counters,omitempty"`
	Country              Country          `json:"country,omitempty"`
	CropPhoto            CropPhoto        `json:"crop_photo,omitempty"`
	Domain               string           `json:"domain,omitempty"`
	Education            UserEducation    `json:"education,omitempty"`
	Exports              any              `json:"exports,omitempty"`
	NameNominative       string           `json:"first_name_nom,omitempty"`
	NameGenitive         string           `json:"first_name_gen,omitempty"`
	NameDative           string           `json:"first_name_dat,omitempty"`
	NameAccusative       string           `json:"first_name_acc,omitempty"`
	NameInstrumental     string           `json:"first_name_ins,omitempty"`
	NameAblative         string           `json:"first_name_abl,omitempty"`
	FollowersCount       int              `json:"followers_count,omitempty"`
	Games                string           `json:"games,omitempty"`
	HasMobile            int              `json:"has_mobile,omitempty"`
	HasPhoto             int              `json:"has_photo,omitempty"`
	HomeTown             string           `json:"home_town,omitempty"`
	Interests            string           `json:"interests,omitempty"`
	IsFavorite           int              `json:"is_favorite,omitempty"`
	IsFriend             int              `json:"is_friend,omitempty"`
	IsHiddenFromFeed     int              `json:"is_hidden_from_feed,omitempty"`
	IsNoIndex            int              `json:"is_no_index,omitempty"`
	SurnameNominative    string           `json:"last_name_nom,omitempty"`
	SurnameGenitive      string           `json:"last_name_gen,omitempty"`
	SurnameDative        string           `json:"last_name_dat,omitempty"`
	SurnameAccusative    string           `json:"last_name_acc,omitempty"`
	SurnameInstrumental  string           `json:"last_name_ins,omitempty"`
	SurnameAblative      string           `json:"last_name_abl,omitempty"`
	LastSeen             UserLastSeen     `json:"last_seen,omitempty"`
	Lists                string           `json:"lists,omitempty"`
	MaidenName           string           `json:"maiden_name,omitempty"`
	Military             UserMilitary     `json:"military,omitempty"`
	Movies               string           `json:"movies,omitempty"`
	Music                string           `json:"music,omitempty"`
	Nickname             string           `json:"nickname,omitempty"`
	Occupation           UserOccupation   `json:"occupation,omitempty"`
	Online               int              `json:"online,omitempty"`
	Personal             UserPersonal     `json:"personal,omitempty"`
	Photo50              string           `json:"photo_50,omitempty"`
	Photo100             string           `json:"photo_100,omitempty"`
	Photo200Original     string           `json:"photo_200_orig,omitempty"`
	Photo200             string           `json:"photo_200,omitempty"`
	Photo400Original     string           `json:"photo_400_orig,omitempty"`
	PhotoID              string           `json:"photo_id,omitempty"`
	PhotoMax             string           `json:"photo_max,omitempty"`
	PhotoMaxOriginal     string           `json:"photo_max_orig,omitempty"`
	Quotes               string           `json:"quotes,omitempty"`
	Relatives            []UserRelative   `json:"relatives,omitempty"`
	Relation             int              `json:"relation,omitempty"`
	Schools              []UserSchool     `json:"schools,omitempty"`
	ScreenName           string           `json:"screen_name,omitempty"`
	Gender               int              `json:"sex,omitempty"`
	Site                 string           `json:"site,omitempty"`
	Status               string           `json:"status,omitempty"`
	Timezone             int              `json:"timezone,omitempty"`
	Trending             int              `json:"trending,omitempty"`
	TV                   string           `json:"tv,omitempty"`
	Universities         []UserUniversity `json:"universities,omitempty"`
	Verified             int              `json:"verified,omitempty"`
	WallDefault          string           `json:"wall_default,omitempty"`
	IsAdmin              int              `json:"is_admin,omitempty"`
	AdminLevel           int              `json:"admin_level,omitempty"`
	IsMember             int              `json:"is_member,omitempty"`
	IsAdvertiser         int              `json:"is_advertiser,omitempty"`
	InvitedBy            int              `json:"invited_by,omitempty"`
	Type                 string           `json:"type,omitempty"`
	Activity             string           `json:"activity,omitempty"`
	Addresses            Addresses        `json:"addresses,omitempty"`
	AgeLimits            int              `json:"age_limits,omitempty"`
	BanInfo              BanInfo          `json:"ban_info,omitempty"`
	CanCreateTopic       int              `json:"can_create_topic,omitempty"`
	CanMessage           int              `json:"can_message,omitempty"`
	CanSuggest           int              `json:"can_suggest,omitempty"`
	CanUploadDoc         int              `json:"can_upload_doc,omitempty"`
	CanUploadStory       int              `json:"can_upload_story,omitempty"`
	CanUploadVideo       int              `json:"can_upload_video,omitempty"`
	Cover                Cover            `json:"cover,omitempty"`
	Description          string           `json:"description,omitempty"`
	PinnedPost           int              `json:"fixed_post,omitempty"`
	IsMessagesBlocked    int              `json:"is_messages_blocked,omitempty"`
	Links                []Link           `json:"links,omitempty"`
	MainAlbumID          int              `json:"main_album_id,omitempty"`
	MainSection          int              `json:"main_section,omitempty"`
	Market               Market           `json:"market,omitempty"`
	MemberStatus         int              `json:"member_status,omitempty"`
	MembersCount         int              `json:"members_count,omitempty"`
	Place                Place            `json:"place,omitempty"`
	PublicDateLabel      string           `json:"public_date_label,omitempty"`
	StartDate            any              `json:"start_date,omitempty"`
	FinishDate           string           `json:"finish_date,omitempty"`
	Wall                 int              `json:"wall,omitempty"`
	WikiPage             string           `json:"wiki_page,omitempty"`
}

type SearchUserReply struct {
	ErrorInterface
	Response SearchUser `json:"response,omitempty"`
}

type SearchUser struct {
	Items []User `json:"items,omitempty"`
}
