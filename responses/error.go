package responses

const InternalErrorCode = 1512

type Error struct {
	Code    int          `json:"error_code,omitempty"`
	Message string       `json:"error_msg,omitempty"`
	Params  []ErrorParam `json:"request_params,omitempty"`
}

type ErrorParam struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}
