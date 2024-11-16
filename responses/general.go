package responses

type Response interface {
	GetError() *Error
}

type ErrorInterface struct {
	Error *Error `json:"error"`
}

func (r *ErrorInterface) GetError() *Error {
	return r.Error
}

type OperationCode struct {
	ErrorInterface
	Code int `json:"response,omitempty"`
}

type ResultOperationCodeReply struct {
	ErrorInterface
	Response ResultOperationCode `json:"response,omitempty"`
}

type ResultOperationCode struct {
	Result int `json:"result,omitempty"`
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
