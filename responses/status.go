package responses

type StatusReply struct {
	ErrorInterface
	Response Status `json:"response"`
}

type Status struct {
	Text string `json:"text,omitempty"`
}
