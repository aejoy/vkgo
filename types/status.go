package types

type GetStatus struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

type SetStatus struct {
	GroupID int    `json:"group_id"`
	Text    string `json:"text"`
}
