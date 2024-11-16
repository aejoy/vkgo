package responses

type Execute struct {
	ErrorInterface
	Result any `json:"response,omitempty"`
}
