package responses

type ResolveDomainReply struct {
	ErrorInterface
	Response ResolveDomain `json:"response"`
}

type ResolveDomain struct {
	Type     string `json:"type,omitempty"`
	ObjectID int    `json:"object_id,omitempty"`
}
