package types

type BanUserComment struct {
	Text string `json:"text"`
	Show bool   `json:"show"`
}

type BanUsers struct {
	GroupID int      `json:"group_id"`
	UserID  int      `json:"user_id"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Fields  []string `json:"fields"`
}
