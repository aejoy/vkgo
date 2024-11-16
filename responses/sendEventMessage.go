package responses

type SendEventMessage struct {
	ErrorInterface
	Code int `json:"response,omitempty"`
}
