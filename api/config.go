package api

func (api *API) SetMessagePack() string {
	api.ContentType = "application/msgpack"
	return api.ContentType
}

func (api *API) SetJSON() string {
	api.ContentType = "application/json"
	return api.ContentType
}
