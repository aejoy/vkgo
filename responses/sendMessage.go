package responses

type SendMessageReply struct {
	ErrorInterface
	Response []SendMessage `json:"response"`
}

type SendMessage struct {
	UserID        int `json:"user_id,omitempty"`
	ChatID        int `json:"peer_id,omitempty"`
	MessageID     int `json:"message_id,omitempty"`
	ChatMessageID int `json:"conversation_message_id,omitempty"`
}
