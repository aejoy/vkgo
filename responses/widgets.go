package responses

type UpdateWidgetsReply struct {
	ErrorInterface
	Response int `json:"response,omitempty"`
}
