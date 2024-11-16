package types

type Execute struct {
	Code    any    `json:"code"`
	Version string `json:"func_v"`
}
