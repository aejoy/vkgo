package types

type User struct {
	UserID   int      `json:"user_ids"`
	UserIDs  []int    `json:"user_ids"` //nolint:govet
	Fields   []string `json:"fields"`
	NameCase string   `json:"name_case"`
}

type Followers struct {
	UserIDs  []int    `json:"user_ids"`
	Offset   int      `json:"offset"`
	Count    int      `json:"count"`
	Fields   []string `json:"fields"`
	NameCase string   `json:"name_case"`
}

type Subscriptions struct {
	UserID int      `json:"user_id"`
	Offset int      `json:"offset"`
	Count  int      `json:"count"`
	Fields []string `json:"fields"`
}

type ReportUser struct {
	UserID  int    `json:"user_id"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
}

type SearchUser struct {
	Query             string `json:"q"`
	Sort              int    `json:"sort"`
	Offset            int    `json:"offset"`
	Count             int    `json:"count"`
	Fields            string `json:"fields"`
	CityID            int    `json:"city"`
	CountryID         int    `json:"country"`
	Hometown          int    `json:"hometown"`
	UniversityCountry int    `json:"university_country"`
	University        int    `json:"university"`
	UniversityYear    int    `json:"university_year"`
	UniversityFaculty int    `json:"university_faculty"`
	UniversityChair   int    `json:"university_chair"`
	Gender            int    `json:"sex"`
	Status            int    `json:"status"`
	AgeFrom           int    `json:"age_from"`
	AgeTo             int    `json:"age_to"`
	BirthDay          int    `json:"birth_day"`
	BirthMonth        int    `json:"birth_month"`
	BirthYear         int    `json:"birth_year"`
	Online            int    `json:"online"`
	HasPhoto          int    `json:"has_photo"`
	SchoolCountry     int    `json:"school_country"`
	SchoolCity        int    `json:"school_city"`
	SchoolClass       int    `json:"school_class"`
	School            int    `json:"school"`
	SchoolYear        int    `json:"school_year"`
	Religion          string `json:"religion"`
	Company           string `json:"company"`
	Position          string `json:"position"`
	GroupID           int    `json:"group_id"`
	FromList          string `json:"from_list"`
	ScreenRef         string `json:"screen_ref"`
}
