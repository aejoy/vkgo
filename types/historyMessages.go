package types

type HistoryMessages struct {
	GroupID        int      `json:"group_id"`
	ChatID         int      `json:"peer_id"`
	UserID         int      `json:"user_id"`
	StartMessageID int      `json:"start_message_id"`
	Offset         int      `json:"offset"`
	Count          int      `json:"count"`
	Rev            bool     `json:"rev"`
	Fields         []string `json:"fields"`
}
