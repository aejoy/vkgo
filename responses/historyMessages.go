package responses

import "github.com/aejoy/vkgo/update"

type HistoryMessagesReply struct {
	ErrorInterface
	Response HistoryMessages `json:"response"`
}

type HistoryMessages struct {
	Count    int              `json:"count,omitempty"`
	Messages []update.Message `json:"items,omitempty"`
	Users    []User           `json:"profiles,omitempty"`
	Chats    []Conversation   `json:"conversations,omitempty"`
	Contacts []Contact        `json:"contacts,omitempty"`
}
