package responses

type StorageKeysReply struct {
	ErrorInterface
	Response []string `json:"response,omitempty"`
}

type StorageReply struct {
	ErrorInterface
	Response Storages `json:"response,omitempty"`
}

type Storages []struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type SetStorageReply struct {
	ErrorInterface
	Response any `json:"response,omitempty"`
}
