package responses

type SaveUploadMessagesDocumentReply struct {
	ErrorInterface
	Response SaveUploadMessagesDocument `json:"response"`
}

type SaveUploadMessagesDocument struct {
	ErrorInterface
	Type     string                    `json:"type,omitempty"`
	Document SaveUploadMessageDocument `json:"doc,omitempty"`
}

type SaveUploadMessageDocument struct {
	URL       string `json:"url,omitempty"`
	Date      int    `json:"date,omitempty"`
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	OwnerID   int    `json:"owner_id,omitempty"`
	Extension string `json:"ext,omitempty"`
	Size      int    `json:"size,omitempty"`
}

type SaveMessagesDocument struct {
	Title string
	Tags  string
}
