package types

// ChangePassword

type ChangePassword struct {
	RestoreSID         string `json:"restore_sid"`
	ChangePasswordHash string `json:"change_password_hash"`
	CurrentPassword    string `json:"old_password"`
	NewPassword        string `json:"new_password"`
}

// GetActiveOffers

type ActiveOffers struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

// GetBans

type Bans struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

// GetInfo

type Info struct {
	Fields []string `json:"fields"`
}

// EditProfileInfo

type EditProfileInfo struct {
	Name                string `json:"first_name"`
	Surname             string `json:"last_name"`
	MaidenName          string `json:"maiden_name"`
	Domain              string `json:"screen_name"`
	CancelRequestID     int    `json:"cancel_request_id"`
	Sex                 int    `json:"sex"`
	Relation            int    `json:"relation"`
	RelationPartnerID   int    `json:"relation_partner_id"`
	Birthdate           string `json:"bdate"`
	BirthdateVisibility int    `json:"bdate_visibility"`
	HomeTown            string `json:"home_town"`
	CountryID           int    `json:"country_id"`
	CityID              int    `json:"city_id"`
	Status              string `json:"status"`
}

// EditAccountInfo

type EditAccountInfo struct {
	Intro         int    `json:"intro"`
	OnlyMyPosts   int    `json:"own_posts_default"`
	NoWallReplise int    `json:"no_wall_replies"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}

type SetOnline struct {
	Calls bool `json:"voip" to:"int"`
}
